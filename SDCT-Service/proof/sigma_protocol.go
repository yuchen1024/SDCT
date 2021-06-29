package proof

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	log "github.com/inconshreveable/log15"
	"github.com/sdct/utils"
)

// DLESigmaProof includes items of zero-knowledge proof.
type DLESigmaProof struct {
	A1 *utils.ECPoint
	A2 *utils.ECPoint

	Z *big.Int
}

// GenerateDLESigmaProof generates zero knowledge proof to prove two ciphertexts encrypt the same value under same public key.
func GenerateDLESigmaProof(params KeyParams, ori, refresh *CTEncPoint, sk *ecdsa.PrivateKey, custom ...*big.Int) (*DLESigmaProof, error) {
	// g1 = Y(fresh) - Y(ori)
	g1 := new(utils.ECPoint).Sub(refresh.Y, ori.Y)
	// h1 = X(fresh) - X(ori)
	h1 := new(utils.ECPoint).Sub(refresh.X, ori.X)
	// g2 = g base point.
	g2 := params.G()
	// h2 = pk.
	h2 := new(utils.ECPoint).SetFromPublicKey(&sk.PublicKey)
	// witness = sk.
	w := new(big.Int).Set(sk.D)
	return generateDLESimaProof(g1, h1, g2, h2, w, custom...)
}

// GenerateEqualProof generates a proof to prove amount is same with value in encrypted ct.
func GenerateEqualProof(params BaseParams, amount *big.Int, ct *CTEncPoint, sk *ecdsa.PrivateKey, custom ...*big.Int) (*DLESigmaProof, error) {
	g1 := new(utils.ECPoint).Sub(ct.Y, new(utils.ECPoint).ScalarMult(params.H(), amount))
	h1 := ct.X
	g2 := params.G()
	h2 := new(utils.ECPoint).SetFromPublicKey(&sk.PublicKey)
	w := new(big.Int).Set(sk.D)

	return generateDLESimaProof(g1, h1, g2, h2, w, custom...)
}

// VerifyEqualProof verifies equal proof.
func VerifyEqualProof(params BaseParams, ct *CTEncPoint, amount *big.Int, pk *utils.ECPoint, proof *DLESigmaProof, custom ...*big.Int) bool {
	g1 := new(utils.ECPoint).Sub(ct.Y, new(utils.ECPoint).ScalarMult(params.H(), amount))
	h1 := ct.X
	g2 := params.G()
	h2 := pk.Copy()
	return verifyDLESigmaProof(g1, h1, g2, h2, proof, custom...)
}

func generateDLESimaProof(g1, h1, g2, h2 *utils.ECPoint, w *big.Int, custom ...*big.Int) (*DLESigmaProof, error) {
	curve := g1.Curve
	n := curve.Params().N
	a, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}

	// A1 = g1 * a; A2 = g2 * a.
	A1 := new(utils.ECPoint).ScalarMult(g1, a)
	A2 := new(utils.ECPoint).ScalarMult(g2, a)
	// compute challenge e prime.
	eprime, _ := utils.ComputeChallenge(n, A1.X, A1.Y, A2.X, A2.Y)
	cinput := make([]interface{}, 0)
	for _, c := range custom {
		cinput = append(cinput, c)
	}
	// compute custom input hash.
	hcustom, _ := utils.ComputeChallenge(n, cinput...)
	// compute final challenge.
	e, err := utils.ComputeChallenge(n, eprime, hcustom)
	if err != nil {
		return nil, err
	}

	// compute z = a + e * w.
	z := new(big.Int).Mul(e, w)
	z = z.Mod(z, n)
	z = z.Add(z, a)
	z = z.Mod(z, n)

	// set proof
	proof := new(DLESigmaProof)
	proof.A1 = A1
	proof.A2 = A2
	proof.Z = z

	return proof, nil
}

// VerifyDLESigmaProof checks the proof is valid or not.
func VerifyDLESigmaProof(params KeyParams, ori, refresh *CTEncPoint, pk *ecdsa.PublicKey, proof *DLESigmaProof, custom ...*big.Int) bool {
	// g1 = Y(fresh) - Y(ori)
	g1 := new(utils.ECPoint).Sub(refresh.Y, ori.Y)
	// h1 = X(fresh) - X(ori)
	h1 := new(utils.ECPoint).Sub(refresh.X, ori.X)
	// g2 = g base point.
	g2 := params.G()
	// h2 = pk.
	h2 := new(utils.ECPoint).SetFromPublicKey(pk)

	return verifyDLESigmaProof(g1, h1, g2, h2, proof, custom...)
}

func verifyDLESigmaProof(g1, h1, g2, h2 *utils.ECPoint, proof *DLESigmaProof, custom ...*big.Int) bool {
	curve := proof.A1.Curve
	n := curve.Params().N

	data := []interface{}{proof.A1.X, proof.A1.Y, proof.A2.X, proof.A2.Y}
	for _, c := range custom {
		data = append(data, c)
	}
	// compute e prime challenge
	eprime, _ := utils.ComputeChallenge(n, proof.A1.X, proof.A1.Y, proof.A2.X, proof.A2.Y)
	cinput := make([]interface{}, 0)
	for _, c := range custom {
		cinput = append(cinput, c)
	}
	// compute custom input hash
	hcustom, _ := utils.ComputeChallenge(n, cinput...)
	// compute final challenge.
	e, err := utils.ComputeChallenge(n, eprime, hcustom)
	if err != nil {
		return false
	}

	// check g1 * z == A1 + h1 * e.
	if !checkDLESigmaProof(g1, proof.A1, h1, proof.Z, e) {
		return false
	}
	// check g2 * z == A2 + h2 * e.
	if !checkDLESigmaProof(g2, proof.A2, h2, proof.Z, e) {
		return false
	}

	return true
}

// checkDLESigmaProof checks g * z == A + h * e.
func checkDLESigmaProof(g, A, H *utils.ECPoint, z, e *big.Int) bool {
	// g * z.
	gz := new(utils.ECPoint).ScalarMult(g, z)
	// h * e + A.
	he := new(utils.ECPoint).ScalarMult(H, e)
	expect := new(utils.ECPoint).Add(A, he)

	if !expect.Equal(gz) {
		log.Warn("g * z != A + h * e", "expect x", expect.X, "expect y", expect.Y, "actual x", gz.X, "actual y", gz.Y)
		return false
	}

	return true
}
