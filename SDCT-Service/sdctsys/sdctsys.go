package sdctsys

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/inconshreveable/log15"
	"github.com/sdct/proof"
	"github.com/sdct/utils"
)

// ConfidentialTx is a tx for sdct transfer system(using aggregate bulletproof).
type ConfidentialTx struct {
	nonce, token *big.Int

	balance  *proof.CTEncPoint
	pk1, pk2 *utils.ECPoint
	transfer *proof.MRTwistedELGamalCTPub

	refreshBalance *proof.CTEncPoint

	updatedBalance *proof.CTEncPoint

	// proof
	sigmaPTEqualityProof *proof.PTEqualityProof
	bulletProof          *proof.AggRangeProof
	sigmaCTValidProof    *proof.CTValidProof
	sigmaDlogeqProof     *proof.DLESigmaProof
}

type soliditySDCTInput struct {
	Points [40]*big.Int

	L, R [proof.LRsize * 2]*big.Int

	Scalars [10]*big.Int
}

// ToSolidityInput formats tx to solidity to verify contract
func (tx *ConfidentialTx) ToSolidityInput() *soliditySDCTInput {
	input := soliditySDCTInput{}
	input.Points[0] = tx.pk1.X
	input.Points[1] = tx.pk1.Y
	input.Points[2] = tx.pk2.X
	input.Points[3] = tx.pk2.Y
	input.Points[4] = tx.transfer.X1.X
	input.Points[5] = tx.transfer.X1.Y
	input.Points[6] = tx.transfer.X2.X
	input.Points[7] = tx.transfer.X2.Y
	input.Points[8] = tx.transfer.X3.X
	input.Points[9] = tx.transfer.X3.Y
	input.Points[10] = tx.transfer.Y.X
	input.Points[11] = tx.transfer.Y.Y
	input.Points[12] = tx.sigmaPTEqualityProof.A1.X
	input.Points[13] = tx.sigmaPTEqualityProof.A1.Y
	input.Points[14] = tx.sigmaPTEqualityProof.A2.X
	input.Points[15] = tx.sigmaPTEqualityProof.A2.Y
	input.Points[16] = tx.sigmaPTEqualityProof.A3.X
	input.Points[17] = tx.sigmaPTEqualityProof.A3.Y
	input.Points[18] = tx.sigmaPTEqualityProof.B.X
	input.Points[19] = tx.sigmaPTEqualityProof.B.Y
	input.Points[20] = tx.refreshBalance.X.X
	input.Points[21] = tx.refreshBalance.X.Y
	input.Points[22] = tx.refreshBalance.Y.X
	input.Points[23] = tx.refreshBalance.Y.Y
	input.Points[24] = tx.sigmaCTValidProof.A.X
	input.Points[25] = tx.sigmaCTValidProof.A.Y
	input.Points[26] = tx.sigmaCTValidProof.B.X
	input.Points[27] = tx.sigmaCTValidProof.B.Y
	input.Points[28] = tx.sigmaDlogeqProof.A1.X
	input.Points[29] = tx.sigmaDlogeqProof.A1.Y
	input.Points[30] = tx.sigmaDlogeqProof.A2.X
	input.Points[31] = tx.sigmaDlogeqProof.A2.Y
	// range proof.
	input.Points[32] = tx.bulletProof.A.X
	input.Points[33] = tx.bulletProof.A.Y
	input.Points[34] = tx.bulletProof.S.X
	input.Points[35] = tx.bulletProof.S.Y
	input.Points[36] = tx.bulletProof.T1.X
	input.Points[37] = tx.bulletProof.T1.Y
	input.Points[38] = tx.bulletProof.T2.X
	input.Points[39] = tx.bulletProof.T2.Y

	// L, R
	for i := 0; i < tx.bulletProof.Len(); i++ {
		input.L[i*2] = tx.bulletProof.Li(i).X
		input.L[i*2+1] = tx.bulletProof.Li(i).Y

		input.R[i*2] = tx.bulletProof.Ri(i).X
		input.R[i*2+1] = tx.bulletProof.Ri(i).Y
	}

	// scalar
	input.Scalars[0] = tx.sigmaPTEqualityProof.Z1
	input.Scalars[1] = tx.sigmaPTEqualityProof.Z2
	input.Scalars[2] = tx.sigmaCTValidProof.Z1
	input.Scalars[3] = tx.sigmaCTValidProof.Z2
	input.Scalars[4] = tx.sigmaDlogeqProof.Z
	// range proof.
	input.Scalars[5] = tx.bulletProof.T()
	input.Scalars[6] = tx.bulletProof.TX()
	input.Scalars[7] = tx.bulletProof.U()
	// inner proof.
	input.Scalars[8] = tx.bulletProof.AIP()
	input.Scalars[9] = tx.bulletProof.BIP()

	return &input
}

// Custom returns custom field added to generate challenge point.
func (tx *ConfidentialTx) Custom() []*big.Int {
	customs := make([]*big.Int, 0)
	customs = append(customs, tx.nonce)
	customs = append(customs, tx.token)
	customs = append(customs, tx.pk1.X)
	customs = append(customs, tx.pk1.Y)
	customs = append(customs, tx.pk2.X)
	customs = append(customs, tx.pk2.Y)
	customs = append(customs, tx.transfer.X1.X)
	customs = append(customs, tx.transfer.X1.Y)
	customs = append(customs, tx.transfer.X2.X)
	customs = append(customs, tx.transfer.X2.Y)
	customs = append(customs, tx.transfer.X3.X)
	customs = append(customs, tx.transfer.X3.Y)
	customs = append(customs, tx.transfer.Y.X)
	customs = append(customs, tx.transfer.Y.Y)

	return customs
}

// CreateConfidentialTx creates confidential transaction to transfer assets from alice to bob.
func CreateConfidentialTx(params proof.AggRangeParams, alice *Account, bob *ecdsa.PublicKey, v, token *big.Int) (*ConfidentialTx, error) {
	ctx := ConfidentialTx{}
	alicePublicKey := &alice.sk.PublicKey

	ctx.nonce = new(big.Int).SetUint64(alice.nonce)
	ctx.token = new(big.Int).Set(token)
	ctx.pk1 = new(utils.ECPoint).SetFromPublicKey(alicePublicKey)
	ctx.pk2 = new(utils.ECPoint).SetFromPublicKey(bob)

	transferEnc, err := proof.EncryptTransfer(params, alicePublicKey, bob, v.Bytes())
	if err != nil {
		return nil, err
	}
	ctx.transfer = transferEnc.Pub()
	// prove alice/bob/auth encrypt the same msg.
	ctx.sigmaPTEqualityProof, err = proof.GeneratePTEqualityProof(params, alicePublicKey, bob, transferEnc)
	if err != nil {
		return nil, err
	}
	ctx.balance = alice.balance
	updateBalanceCT := new(proof.CTEncPoint).Sub(ctx.balance, ctx.transfer.First())
	ctx.updatedBalance = updateBalanceCT.Copy()
	refreshBalanceCT, err := proof.Refresh(params, alice.sk, updateBalanceCT)
	// for speed up.
	refreshBalanceCT.EncMsg = new(big.Int).Sub(alice.m, v).Bytes()
	if err != nil {
		return nil, err
	}
	ctx.refreshBalance = refreshBalanceCT.CopyPublicPoint()
	customs := ctx.Custom()
	ctx.sigmaDlogeqProof, err = proof.GenerateDLESigmaProof(params, updateBalanceCT, ctx.refreshBalance,
		alice.sk, customs...)
	if err != nil {
		return nil, err
	}

	ctx.sigmaCTValidProof, err = proof.GenerateCTValidProof(params, alicePublicKey, refreshBalanceCT)
	if err != nil {
		return nil, err
	}

	vlist := make([]*big.Int, 0)
	vlist = append(vlist, new(big.Int).SetBytes(transferEnc.EncMsg))
	vlist = append(vlist, new(big.Int).SetBytes(refreshBalanceCT.EncMsg))
	random := make([]*big.Int, 0)
	random = append(random, transferEnc.R)
	random = append(random, refreshBalanceCT.R)
	ctx.bulletProof, err = proof.GenerateAggRangeProof(proof.Reserve(params), vlist, random)
	if err != nil {
		return nil, err
	}

	return &ctx, nil
}

// VerifyConfidentialTx checks tx .
func VerifyConfidentialTx(params proof.AggRangeParams, ctx *ConfidentialTx) bool {

	if !proof.VerifyPTEqualityProof(params, ctx.pk1.ToPublicKey(), ctx.pk2.ToPublicKey(), ctx.transfer, ctx.sigmaPTEqualityProof) {
		log.Warn("verify pte equality proof failed")
		return false
	}

	updatedBalance := new(proof.CTEncPoint).Sub(ctx.balance, ctx.transfer.First())

	customs := ctx.Custom()
	if !proof.VerifyDLESigmaProof(params, updatedBalance, ctx.refreshBalance, ctx.pk1.ToPublicKey(), ctx.sigmaDlogeqProof, customs...) {
		log.Warn("verify dle sigma proof failed")
		return false
	}

	if !proof.VerifyCTValidProof(params, ctx.pk1.ToPublicKey(), ctx.refreshBalance, ctx.sigmaCTValidProof) {
		log.Warn("verify ct valid proof failed")
		return false
	}

	vpoints := make([]*utils.ECPoint, 0)
	vpoints = append(vpoints, ctx.transfer.Y)
	vpoints = append(vpoints, ctx.refreshBalance.Y)
	if !proof.VerifyAggRangeProof(proof.Reserve(params), vpoints, ctx.bulletProof) {
		log.Warn("verify aggregate proof failed")
		return false
	}

	return true
}

// BurnETHTx includes info to burn an account and withdraw eth.
type BurnETHTx struct {
	Account  *utils.ECPoint       `json:"account"`
	Amount   *big.Int             `json:"amount"`
	Proof    *proof.DLESigmaProof `json:"proof"`
	Receiver common.Address
}

type burnEHTTxInput struct {
	Receiver common.Address
	Amount   *big.Int
	PubKey   [2]*big.Int
	Proof    [4]*big.Int
	Z        *big.Int
}

func (btx *BurnETHTx) ToSolidityInput() *burnEHTTxInput {
	return &burnEHTTxInput{
		Receiver: btx.Receiver,
		Amount:   new(big.Int).Set(btx.Amount),
		PubKey:   [2]*big.Int{btx.Account.X, btx.Account.Y},
		Proof:    [4]*big.Int{btx.Proof.A1.X, btx.Proof.A1.Y, btx.Proof.A2.X, btx.Proof.A2.Y},
		Z:        btx.Proof.Z,
	}
}

// CreateBurnETHTx creates tx to burn eth on chain.
func CreateBurnETHTx(params proof.AggRangeParams, alice *Account, receiver, token common.Address) (*BurnETHTx, error) {
	tx := BurnETHTx{}

	tx.Account = new(utils.ECPoint).SetFromPublicKey(&alice.sk.PublicKey)
	tx.Amount = new(big.Int).Set(alice.m)

	// generate proof to prove alice has the sk and the amount is indeed same with value encrypted.
	// alice's encrypted balance should be right set.
	proof, err := proof.GenerateEqualProof(params, tx.Amount, alice.balance, alice.sk, new(big.Int).SetUint64(alice.nonce),
		receiver.Hash().Big(), token.Hash().Big())
	if err != nil {
		return nil, err
	}
	tx.Proof = proof

	return &tx, nil
}

// VerifyBurnETHTx .
func VerifyBurnETHTx(params proof.AggRangeParams, nonce *big.Int, receiver common.Address, balance *proof.CTEncPoint, btx *BurnETHTx) bool {
	return proof.VerifyEqualProof(params, balance, btx.Amount, btx.Account, btx.Proof, nonce, receiver.Hash().Big())
}
