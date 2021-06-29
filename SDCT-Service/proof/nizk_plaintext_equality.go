package proof

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	log "github.com/inconshreveable/log15"
	"github.com/sdct/utils"
)

// BaseParams contains all params to generate normal proof.
type BaseParams interface {
	G() *utils.ECPoint
	H() *utils.ECPoint
	Priv() *ecdsa.PrivateKey
}

type baseParams struct {
	g, h *utils.ECPoint
	priv *ecdsa.PrivateKey
}

func (bp *baseParams) G() *utils.ECPoint {
	return bp.g
}

func (bp *baseParams) H() *utils.ECPoint {
	return bp.h
}

func (bp *baseParams) Priv() *ecdsa.PrivateKey {
	return bp.priv
}

// PTEqualityProof .
type PTEqualityProof struct {
	A1, A2, A3, B *utils.ECPoint

	Z1, Z2 *big.Int
}

// GeneratePTEqualityProof generates proof to prove msg in two plaintext is the same.
func GeneratePTEqualityProof(params BaseParams, pk1, pk2 *ecdsa.PublicKey, ct *MRTwistedELGamalCT) (*PTEqualityProof, error) {
	curve := ct.X1.Curve
	n := curve.Params().N

	// get random.
	a, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}
	b, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}

	proof := PTEqualityProof{}
	// A1=pk1^a; A2=pk2^a; A3=pkauth^a.
	proof.A1 = new(utils.ECPoint).SetFromPublicKey(pk1)
	proof.A1.ScalarMult(proof.A1, a)
	proof.A2 = new(utils.ECPoint).SetFromPublicKey(pk2)
	proof.A2.ScalarMult(proof.A2, a)
	proof.A3 = new(utils.ECPoint).SetFromPublicKey(&params.Priv().PublicKey)
	proof.A3.ScalarMult(proof.A3, a)

	// B = g*a + h*b.
	g := params.G()
	h := params.H()
	proof.B = new(utils.ECPoint).ScalarMult(g, a)
	proof.B.Add(proof.B, new(utils.ECPoint).ScalarMult(h, b))

	e, err := utils.ComputeChallengeByECPoints(curve.Params().N, proof.A1, proof.A2, proof.B)
	if err != nil {
		return nil, err
	}

	r := new(big.Int).Set(ct.R)
	v := new(big.Int).SetBytes(ct.EncMsg)
	// Z1 = a + e * r.
	proof.Z1 = new(big.Int).Mul(e, r)
	proof.Z1.Mod(proof.Z1, n)
	proof.Z1.Add(proof.Z1, a)
	proof.Z1.Mod(proof.Z1, n)

	// Z2 = b + e * v.
	proof.Z2 = new(big.Int).Mul(e, v)
	proof.Z2.Mod(proof.Z2, n)
	proof.Z2.Add(proof.Z2, b)
	proof.Z2.Mod(proof.Z2, n)

	return &proof, nil
}

// VerifyPTEqualityProof verifies proof.
func VerifyPTEqualityProof(params BaseParams, pk1, pk2 *ecdsa.PublicKey, ct *MRTwistedELGamalCTPub, proof *PTEqualityProof) bool {
	curve := pk1.Curve
	n := curve.Params().N
	// compute challenge.
	e, err := utils.ComputeChallengeByECPoints(n, proof.A1, proof.A2, proof.B)
	if err != nil {
		log.Warn("verify pteequality proof failed(compute challenge)", "err", err)
		return false
	}
	// check pk1 * z1 == A1 + X1 * e.
	if !checkTwistedELGamalCTX(new(utils.ECPoint).SetFromPublicKey(pk1), proof.A1, ct.X1, proof.Z1, e) {
		return false
	}
	// check pk2 * z1 == A2 + X2 * e.
	if !checkTwistedELGamalCTX(new(utils.ECPoint).SetFromPublicKey(pk2), proof.A2, ct.X2, proof.Z1, e) {
		return false
	}
	// check pkauth * z1 == A3 + X3*e.
	if !checkTwistedELGamalCTX(new(utils.ECPoint).SetFromPublicKey(&params.Priv().PublicKey), proof.A3, ct.X3, proof.Z1, e) {
		return false
	}

	h := params.H()
	g := params.G()
	// Check g * z1 + h * z2 == B + Y * e.
	if !checkTwistedELGamalCTY(g, h, proof.B, ct.Y, proof.Z1, proof.Z2, e) {
		return false
	}

	return true
}

// checkTwistedELGamalCTX checks pk * z == A + X * e or not(check X).
func checkTwistedELGamalCTX(pk, A, X *utils.ECPoint, z, e *big.Int) bool {

	// compute x * e + A
	actual := new(utils.ECPoint).ScalarMult(X, e)
	actual.Add(actual, A)
	// compute pk * z.
	expect := new(utils.ECPoint).ScalarMult(pk, z)

	if !expect.Equal(actual) {
		log.Warn("pk * z != A + X * e", "expect x", expect.X, "actual x", actual.X, "expect y", expect.Y, "actual y", actual.Y)
		return false
	}

	return true
}

// checkTwistedELGamalCTY checks g*z1 + h*z2 == B + Y*e(check Y).
func checkTwistedELGamalCTY(g, h, B, Y *utils.ECPoint, z1, z2, e *big.Int) bool {
	// compute g * z1 + h * z2.
	gzb := new(utils.ECPoint).ScalarMult(g, z1)
	actual := new(utils.ECPoint).ScalarMult(h, z2)
	actual.Add(actual, gzb)
	// compute B + Y * e.
	expect := new(utils.ECPoint).ScalarMult(Y, e)
	expect.Add(expect, B)
	if !expect.Equal(actual) {
		log.Warn("g * z1 + h * z2 != B + Y * e", "expect x", expect.X, "actual x", actual.X, "expect y", expect.Y, "actual y", actual.Y)
		return false
	}

	return true
}
