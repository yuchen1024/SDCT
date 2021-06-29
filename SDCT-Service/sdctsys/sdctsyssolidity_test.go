package sdctsys

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
	"testing"

	"github.com/sdct/contracts/sdctsystem"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/inconshreveable/log15"
	"github.com/sdct/client"
	"github.com/sdct/contracts/tokenconverter"
	"github.com/sdct/deployer"
	"github.com/sdct/proof"
	"github.com/sdct/utils"
	"github.com/stretchr/testify/require"
)

var (
	precision = new(big.Int).SetUint64(100)
)

func TestSDCTSystemContractETHLocal(t *testing.T) {
	rpcclient := client.GetLocalRPC()
	ethclient := client.GetLocal()
	auth := rpcclient.GetAccountWithETH()

	testSDCTSystemContract(t, false, auth, ethclient)
}

func TestSDCTSystemContractTokenLocal(t *testing.T) {
	rpcclient := client.GetLocalRPC()
	ethclient := client.GetLocal()
	auth := rpcclient.GetAccountWithETH()

	testSDCTSystemContract(t, true, auth, ethclient)
}

func testSDCTSystemContract(t *testing.T, tokenTest bool, auth *bind.TransactOpts, ethclient *ethclient.Client) {
	params := proof.DAggRangeProofParams32()
	proof.BuildAndLoadMapIfNotExist(params.H(), 32, 7, 4)
	addrs, sdct := deployer.DeploySDCTSystemAllContract(auth, ethclient, &params.Priv().PublicKey)

	deployer.InitVector(auth, ethclient, addrs[2], 32)

	token := common.Address{}
	if tokenTest {
		token = setForToken(t, addrs, auth, ethclient)
	}

	// alice
	aliceAmount := new(big.Int).SetUint64(512)
	name := "Alice"
	alice := CreateTestAccount(params, name, aliceAmount)
	// log out alice info
	alice.Info()

	// bob.
	bobAmount := new(big.Int).SetUint64(256)
	name = "Bob"
	bob := CreateTestAccount(params, name, bobAmount)
	bob.Info()

	initTestAccount(t, alice, params, token, aliceAmount, "Alice", auth, ethclient, sdct)
	initTestAccount(t, bob, params, token, bobAmount, name, auth, ethclient, sdct)

	// log out alice/bob
	fmt.Println("")
	fmt.Println("Deposit succeeds")
	fmt.Println("Alice's current balance", alice.m)
	fmt.Println("Bob's current balance", bob.m)

	transferAmount := new(big.Int).SetUint64(128)
	auth.GasLimit = 8000000
	ctx := aggTransfer(t, params, alice, bob, token, transferAmount, auth, ethclient, sdct)

	fmt.Println("")
	fmt.Println("CTx transfer succeeds")
	fmt.Println("Alice's current balance", alice.m)
	fmt.Println("Bob's current balance", bob.m)

	receiver := auth.From
	burn(t, params, alice, receiver, token, auth, ethclient, sdct)
	burn(t, params, bob, receiver, token, auth, ethclient, sdct)

	// inspect
	fmt.Println("")
	fmt.Println("Supervision Begins")
	fmt.Println("Inspect ctx")
	pkx := elliptic.MarshalCompressed(params.Curve(), alice.Priv().X, alice.Priv().Y)
	pky := elliptic.MarshalCompressed(params.Curve(), bob.Priv().X, bob.Priv().Y)
	v := openCTx(ctx, params)
	fmt.Printf("%x transfer %d sdct tokens to %x\n", pkx, v, pky)
}

func setForToken(t *testing.T, addrs []common.Address, auth *bind.TransactOpts, ethclient *ethclient.Client) common.Address {
	token, tokenCon := deployer.DeployToken(auth, ethclient)

	sdctAddr := addrs[len(addrs)-1]
	approveAmount := new(big.Int).SetUint64(100000)

	_, err := tokenCon.Approve(auth, sdctAddr, approveAmount)
	require.Nil(t, err, "approve for token failed", err)
	auth.Nonce.Add(auth.Nonce, utils.One)

	tokenConverter, err := tokenconverter.NewTokenconverter(addrs[1], ethclient)
	require.Nil(t, err, "get token converter failed", err)

	_, err = tokenConverter.AddToken(auth, token, utils.One, "")
	require.Nil(t, err, "add token failed", err)
	auth.Nonce.Add(auth.Nonce, utils.One)

	return token
}

func initTestAccount(t *testing.T, alice *Account, params proof.AggRangeParams, token common.Address, amount *big.Int, name string, auth *bind.TransactOpts, ethclient *ethclient.Client, sdct *sdctsystem.Sdctsystem) *Account {
	alicePK := [2]*big.Int{alice.sk.PublicKey.X, alice.sk.PublicKey.Y}

	var aliceTx *types.Transaction
	var err error
	if isToken(token) {
		auth.Value = new(big.Int).Mul(utils.Ether, amount)
		auth.Value.Div(auth.Value, precision)
		aliceTx, err = sdct.DepositAccountETH(auth, alicePK)

	} else {
		aliceTx, err = sdct.DepositAccount(auth, alicePK, token, amount)
	}

	require.Nil(t, err, "deposit contract failed")

	receipt := client.WaitForTx(ethclient, aliceTx.Hash())
	auth.Nonce.Add(auth.Nonce, utils.One)
	// check for alice's encrypted balance.
	aliceEncryptB, _ := sdct.GetUserBalance(utils.CallOpt(), alice.sk.PublicKey.X, alice.sk.PublicKey.Y, token)
	alice.UpdateBalance(aliceEncryptB.Nonce, aliceEncryptB.Ct)
	require.Equal(t, amount.Bytes(), alice.m.Bytes(), "account balance on chain not same with local", "expect", amount, "actual", alice.m)
	log.Info("Deposit account succeeds", "token", token.Hash().Hex(), "name", name, "amount", amount, "gas", receipt.GasUsed, "tx", aliceTx.Hash().Hex())
	auth.Value = nil

	return alice
}

func aggTransfer(t *testing.T, params proof.AggRangeParams, from, to *Account, token common.Address, amount *big.Int, auth *bind.TransactOpts, ethclient *ethclient.Client, sdct *sdctsystem.Sdctsystem) *ConfidentialTx {
	ctx, err := CreateConfidentialTx(params, from, &to.sk.PublicKey, amount, token.Hash().Big())
	require.Nil(t, err, "generate confidential tx failed", err)
	input := ctx.ToSolidityInput()

	var tx *types.Transaction
	if isToken(token) {
		tx, err = sdct.TransferETH(auth, input.Points, input.Scalars, input.L, input.R)
	} else {
		tx, err = sdct.Transfer(auth, input.Points, input.Scalars, token.Hash().Big(), input.L, input.R)
	}

	require.Nil(t, err, "create agg transfer tx failed", err)
	receipt := client.WaitForTx(ethclient, tx.Hash())
	auth.Nonce.Add(auth.Nonce, utils.One)

	// check sender's balance.
	before := new(big.Int).Set(from.m)
	shouldBe := before.Sub(before, amount)
	encryptB, _ := sdct.GetUserBalance(utils.CallOpt(), from.sk.PublicKey.X, from.sk.PublicKey.Y, token)
	from.UpdateBalance(encryptB.Nonce, encryptB.Ct)
	require.Equal(t, shouldBe.Bytes(), from.m.Bytes(), "sender's balance on chain invalid", "expect", shouldBe, "actual", from.m)

	// check receiver's balance.
	before = new(big.Int).Set(to.m)
	shouldBe = before.Add(before, amount)
	encryptB, _ = sdct.GetUserBalance(utils.CallOpt(), to.sk.PublicKey.X, to.sk.PublicKey.Y, token)
	to.UpdateBalance(encryptB.Nonce, encryptB.Ct)
	require.Equal(t, shouldBe.Bytes(), to.m.Bytes(), "receiver's balance on chain invalid", "expect", shouldBe, "actual", to.m)

	log.Info("CTx transfer", "token", token.Hash().Hex(), "amount", amount, "gas", receipt.GasUsed, "tx", tx.Hash().Hex())

	return ctx
}

func burn(t *testing.T, params proof.AggRangeParams, from *Account, receiver, token common.Address, auth *bind.TransactOpts, ethclient *ethclient.Client, sdct *sdctsystem.Sdctsystem) {
	tx, err := CreateBurnETHTx(params, from, receiver, token)
	require.Nil(t, err, "generate burn eth tx failed")

	input := tx.ToSolidityInput()

	var btx *types.Transaction
	if isToken(token) {
		btx, err = sdct.BurnETH(auth, receiver, input.Amount, input.PubKey, input.Proof, input.Z)
	} else {
		btx, err = sdct.Burn(auth, receiver, token.Hash().Big(), input.Amount, input.PubKey, input.Proof, input.Z)
	}

	require.Nil(t, err, "create burn tx failed")
	receipt := client.WaitForTx(ethclient, btx.Hash())
	auth.Nonce.Add(auth.Nonce, utils.One)

	// check balance.
	encryptB, _ := sdct.GetUserBalance(utils.CallOpt(), from.sk.PublicKey.X, from.sk.PublicKey.Y, token)
	from.UpdateBalance(encryptB.Nonce, encryptB.Ct)
	require.Equal(t, uint64(0), from.m.Uint64(), "receiver's balance on chain invalid", "expect", 0, "actual", from.m)

	log.Info("Burn tx succeeds", "token", token.Hash().Hex(), "name", from.name, "amount", input.Amount, "gas", receipt.GasUsed, "tx", btx.Hash().Hex())
}

func isToken(token common.Address) bool {
	return token.Hash().Big().Cmp(utils.Zero) == 0
}

func openCTx(ctx *ConfidentialTx, params proof.CTParams) *big.Int {
	input := ctx.ToSolidityInput()
	ct := proof.CTEncPoint{}
	ct.X = utils.NewECPoint(input.Points[8], input.Points[9], params.Curve())
	ct.Y = utils.NewECPoint(input.Points[10], input.Points[11], params.Curve())

	return new(big.Int).SetBytes(proof.Decrypt(params, params.Priv(), &ct))
}
