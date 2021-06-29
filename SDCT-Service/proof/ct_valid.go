package proof

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	log "github.com/inconshreveable/log15"
	"github.com/sdct/utils"
)

// CTValidProof is a proof to prove the ct tx is valid.
type CTValidProof struct {
	A, B   *utils.ECPoint
	Z1, Z2 *big.Int
}

// GenerateCTValidProof generates a valid CTValidProof.
func GenerateCTValidProof(params BaseParams, pk *ecdsa.PublicKey, ct *TwistedELGamalCT) (*CTValidProof, error) {
	curve := pk.Curve
	n := curve.Params().N
	a, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}
	b, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}

	proof := CTValidProof{}
	// A = pk*a.
	proof.A = new(utils.ECPoint).SetFromPublicKey(pk)
	proof.A.ScalarMult(proof.A, a)

	// B = g*a + h*b.
	proof.B = new(utils.ECPoint).ScalarMult(params.H(), b)
	tmp := new(utils.ECPoint).ScalarMult(params.G(), a)
	proof.B.Add(proof.B, tmp)

	e, err := utils.ComputeChallengeByECPoints(n, proof.A, proof.B)
	if err != nil {
		return nil, err
	}

	// z1 = a + e*r.
	proof.Z1 = new(big.Int).Set(e)
	proof.Z1.Mul(proof.Z1, ct.R)
	proof.Z1.Add(proof.Z1, a)
	proof.Z1.Mod(proof.Z1, n)
	// z2 = b + e*v.
	proof.Z2 = new(big.Int).Set(e)
	proof.Z2.Mul(proof.Z2, new(big.Int).SetBytes(ct.EncMsg))
	proof.Z2.Add(proof.Z2, b)
	proof.Z2.Mod(proof.Z2, n)

	return &proof, nil
}

// VerifyCTValidProof checks the ctvalid proof is valid or not.
func VerifyCTValidProof(params BaseParams, pk *ecdsa.PublicKey, ct *CTEncPoint, proof *CTValidProof) bool {
	n := pk.Curve.Params().N
	e, err := utils.ComputeChallengeByECPoints(n, proof.A, proof.B)
	if err != nil {
		log.Warn("compute challenge failed", "err", err)
	}

	// check pk*z1 = A + X*e.
	left := new(utils.ECPoint).SetFromPublicKey(pk)
	left.ScalarMult(left, proof.Z1)
	right := new(utils.ECPoint).ScalarMult(ct.X, e)
	right.Add(right, proof.A)
	if !left.Equal(right) {
		log.Warn("left not equal right", "left x", left.X, "right x", right.X, "left y", left.Y, "right y", right.Y)
		return false
	}

	// check g*z1 + h*z2 = B + Y*e.
	left = new(utils.ECPoint).ScalarMult(params.G(), proof.Z1)
	tmp := new(utils.ECPoint).ScalarMult(params.H(), proof.Z2)
	left.Add(left, tmp)
	right = new(utils.ECPoint).ScalarMult(ct.Y, e)
	right.Add(right, proof.B)
	if !left.Equal(right) {
		log.Warn("left not equal right", "left x", left.X, "right x", right.X, "left y", left.Y, "right y", right.Y)
		return false
	}

	return true
}
