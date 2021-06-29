package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/inconshreveable/log15"
	"github.com/sdct/client"
	"github.com/sdct/contracts/sdctsystem"
	"github.com/sdct/deployer"
	"github.com/sdct/proof"
	"github.com/sdct/sdctsys"
	"github.com/sdct/utils"
)

var (
	precision = new(big.Int).SetUint64(100)
)

func main() {
	params := proof.DAggRangeProofParams32()
	proof.BuildAndLoadMapIfNotExist(params.H(), 32, 7, 4)

	// get client
	// rpcclient := client.GetLocalRPC()
	ethclient := client.GetLocal()

	// get ethereum account
	authAlice := client.GetAccountWithKey("99d7445c7d56fd6ef8fae65cc8a5453da75d5866010a9c842c3ea4a503bd8ecc")
	client.SetNonce(authAlice, ethclient)
	authDeploy := authAlice

	authBob := client.GetAccountWithKey("06e77e5cacfdbb648cf576ec5fd70c7b8a5249f4d59329c0d51b216823adfb07")
	client.SetNonce(authBob, ethclient)
	authCarol := client.GetAccountWithKey("62115f81c985a62c69e44dcc3d097aec02bf97fd95fb322bd32bc65e2f0a8131")
	client.SetNonce(authCarol, ethclient)

	addrs, sdct := deployer.DeploySDCTSystemAllContract(authDeploy, ethclient, &params.Priv().PublicKey)
	deployer.InitVector(authDeploy, ethclient, addrs[2], 32)

	fmt.Scanln()

	token := common.Address{}
	// alice
	aliceAmount := new(big.Int).SetUint64(128)
	name := "Alice"
	alice := sdctsys.CreateTestAccount(params, name, aliceAmount)
	// log out alice info
	alice.Info()

	// bob.
	bobAmount := new(big.Int).SetUint64(128)
	name = "Bob"
	bob := sdctsys.CreateTestAccount(params, name, bobAmount)
	bob.Info()

	carolAmount := new(big.Int).SetUint64(128)
	name = "Carol"
	carol := sdctsys.CreateTestAccount(params, name, carolAmount)
	carol.Info()

	fmt.Scanln()
	initTestAccount(alice, params, token, aliceAmount, "Alice", authAlice, ethclient, sdct)
	initTestAccount(bob, params, token, bobAmount, "Bob", authBob, ethclient, sdct)
	initTestAccount(carol, params, token, carolAmount, "Carol", authCarol, ethclient, sdct)

	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Println("Deposit succeeds")
	fmt.Println("Alice's current balance", alice.M())
	fmt.Println("Bob's current balance", bob.M())
	fmt.Println("Carol's current balance", carol.M())
	fmt.Scanln()

	transferAmount := new(big.Int).SetUint64(128)
	authCarol.GasLimit = 8000000
	ctx := aggTransfer(params, carol, bob, token, transferAmount, authCarol, ethclient, sdct)

	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Println("CTx transfer succeeds: Carol transfer 128 coins to Bob")
	fmt.Println("Carol's current balance", carol.M())
	fmt.Println("Bob's current balance", bob.M())
	fmt.Scanln()

	authBob.GasLimit = 8000000
	ctx = aggTransfer(params, bob, alice, token, transferAmount, authBob, ethclient, sdct)
	fmt.Println("")
	fmt.Println("CTx transfer succeeds: Bob transfer 128 coins to Alice")
	fmt.Println("Bob's current balance", bob.M())
	fmt.Println("Alice's current balance", alice.M())
	fmt.Scanln()

	fmt.Println("-----------------------------------------------------------------------------")
	receiver := authAlice.From
	burn(params, alice, receiver, token, authAlice, ethclient, sdct)
	fmt.Println("Alice deposites SDCT token back to ETH")
	fmt.Println("-----------------------------------------------------------------------------")
	receiver = authBob.From
	burn(params, bob, receiver, token, authBob, ethclient, sdct)
	fmt.Println("Bob deposites SDCT token back to ETH")
	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Println("Carol deposites SDCT token back to ETH")

	fmt.Scanln()

	// inspect
	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Println("Supervision Begins")
	fmt.Println("Inspect ctx")
	pkx := elliptic.MarshalCompressed(params.Curve(), alice.Priv().X, alice.Priv().Y)
	pky := elliptic.MarshalCompressed(params.Curve(), bob.Priv().X, bob.Priv().Y)
	v := openCTx(ctx, params)
	fmt.Printf("%x transfer %d SDCT tokens to %x\n", pkx, v, pky)
}

func initTestAccount(alice *sdctsys.Account, params proof.AggRangeParams, token common.Address, amount *big.Int, name string, auth *bind.TransactOpts, ethclient *ethclient.Client, sdct *sdctsystem.Sdctsystem) *sdctsys.Account {
	alicePK := [2]*big.Int{alice.Priv().PublicKey.X, alice.Priv().PublicKey.Y}

	auth.Value = new(big.Int).Mul(utils.Ether, amount)
	aliceTx, err := sdct.DepositAccountETH(auth, alicePK)
	if err != nil {
		panic(err)
	}

	receipt := client.WaitForTx(ethclient, aliceTx.Hash())
	auth.Nonce.Add(auth.Nonce, utils.One)

	// check for alice's encrypted balance.
	aliceEncryptB, _ := sdct.GetUserBalance(utils.CallOpt(), alice.Priv().PublicKey.X, alice.Priv().PublicKey.Y, token)
	alice.UpdateBalance(aliceEncryptB.Nonce, aliceEncryptB.Ct)

	log.Info("Deposit account succeeds", "name", name, "amount", amount, "gas", receipt.GasUsed, "tx", aliceTx.Hash().Hex())
	auth.Value = nil

	return alice
}

func aggTransfer(params proof.AggRangeParams, from, to *sdctsys.Account, token common.Address, amount *big.Int, auth *bind.TransactOpts, ethclient *ethclient.Client, sdct *sdctsystem.Sdctsystem) *sdctsys.ConfidentialTx {
	ctx, _ := sdctsys.CreateConfidentialTx(params, from, &to.Priv().PublicKey, amount, token.Hash().Big())
	input := ctx.ToSolidityInput()

	var tx *types.Transaction
	tx, _ = sdct.TransferETH(auth, input.Points, input.Scalars, input.L, input.R)

	receipt := client.WaitForTx(ethclient, tx.Hash())
	auth.Nonce.Add(auth.Nonce, utils.One)

	// check sender's balance.
	encryptB, _ := sdct.GetUserBalance(utils.CallOpt(), from.Priv().PublicKey.X, from.Priv().PublicKey.Y, token)
	from.UpdateBalance(encryptB.Nonce, encryptB.Ct)

	// check receiver's balance.
	encryptB, _ = sdct.GetUserBalance(utils.CallOpt(), to.Priv().PublicKey.X, to.Priv().PublicKey.Y, token)
	to.UpdateBalance(encryptB.Nonce, encryptB.Ct)

	log.Info("CTx transfer", "token", token.Hash().Hex(), "amount", amount, "gas", receipt.GasUsed, "tx", tx.Hash().Hex())

	return ctx
}

func burn(params proof.AggRangeParams, from *sdctsys.Account, receiver, token common.Address, auth *bind.TransactOpts, ethclient *ethclient.Client, sdct *sdctsystem.Sdctsystem) {
	tx, _ := sdctsys.CreateBurnETHTx(params, from, receiver, token)

	input := tx.ToSolidityInput()

	var btx *types.Transaction
	btx, _ = sdct.BurnETH(auth, receiver, input.Amount, input.PubKey, input.Proof, input.Z)

	receipt := client.WaitForTx(ethclient, btx.Hash())
	auth.Nonce.Add(auth.Nonce, utils.One)

	// check balance.
	encryptB, _ := sdct.GetUserBalance(utils.CallOpt(), from.Priv().PublicKey.X, from.Priv().PublicKey.Y, token)
	from.UpdateBalance(encryptB.Nonce, encryptB.Ct)

	log.Info("Burn tx succeeds", "token", token.Hash().Hex(), "name", from.Name(), "amount", input.Amount, "gas", receipt.GasUsed, "tx", btx.Hash().Hex())
}

func toHex(key *ecdsa.PrivateKey) string {
	return hex.EncodeToString(crypto.FromECDSA(key))
}

func openCTx(ctx *sdctsys.ConfidentialTx, params proof.CTParams) *big.Int {
	input := ctx.ToSolidityInput()
	ct := proof.CTEncPoint{}
	ct.X = utils.NewECPoint(input.Points[8], input.Points[9], params.Curve())
	ct.Y = utils.NewECPoint(input.Points[10], input.Points[11], params.Curve())

	return new(big.Int).SetBytes(proof.Decrypt(params, params.Priv(), &ct))
}
