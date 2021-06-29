package proof

import (
	"crypto/elliptic"
	"math/big"
	"testing"

	"github.com/sdct/curve"
	"github.com/stretchr/testify/require"
)

func TestCTValidProof(t *testing.T) {
	curves := []elliptic.Curve{curve.BN256(), curve.S256()}
	bitsizes := []int{16, 32, 64}

	for _, curve := range curves {
		for _, bitsize := range bitsizes {
			params := newRandomCTParams(curve, bitsize)
			// generate key.
			alice := MustGenerateKey(params)

			msg := new(big.Int).SetUint64(1000)
			ct, err := Encrypt(params, &alice.PublicKey, msg.Bytes())
			require.Nil(t, err, "encrypt msg failed", err)

			proof, err := GenerateCTValidProof(params, &alice.PublicKey, ct)
			require.Nil(t, err, "generate ct valid proof failed", err)

			result := VerifyCTValidProof(params, &alice.PublicKey, ct.CopyPublicPoint(), proof)
			require.Equal(t, result, true, "verify a valid ctvalid proof failed")
		}
	}
}
