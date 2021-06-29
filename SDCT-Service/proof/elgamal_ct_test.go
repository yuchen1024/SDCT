package proof

import (
	"bytes"
	"crypto/elliptic"
	"math/big"
	"testing"

	"github.com/sdct/curve"
	"github.com/sdct/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyGenerate(t *testing.T) {
	curves := []elliptic.Curve{curve.BN256(), curve.S256()}

	for _, curve := range curves {
		params := newRandomKeyParams(curve)
		// generate key.
		key, err := GenerateKey(params)
		if err != nil {
			panic("generate key failed")
		}

		x, y := curve.ScalarMult(params.G().X, params.G().Y, key.D.Bytes())
		assert.Equal(t, key.X, x)
		assert.Equal(t, key.Y, y)
	}

}

func TestCT(t *testing.T) {
	msgs := []*big.Int{
		big.NewInt(0),
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(-1),
	}

	params := DAggRangeProofParams32()
	loadHashMap(32, 7)

	for _, msg := range msgs {
		// generate key.
		key, err := GenerateKey(params)
		require.Nil(t, err, "generate key failed")

		// Encrypt msg.
		// Just use a certain r for test.
		ct, err := Encrypt(params, &key.PublicKey, msg.Bytes())
		require.Nil(t, err, "encrypt data failed")

		newMsg := Decrypt(params, key, ct.CopyPublicPoint())
		if !bytes.Equal(msg.Bytes(), newMsg) {
			t.Error("encrypt/decrypt msg not equal")
		}

		nct, err := Refresh(params, key, ct.CopyPublicPoint())
		require.Nil(t, err, "refresh failed")
		assert.NotEqual(t, ct.R, nct.R, "randome equal after refresh")
	}
}

func TestEncryptTransfer(t *testing.T) {
	curves := []elliptic.Curve{curve.BN256(), curve.S256()}
	bitsize := 16

	for _, curve := range curves {
		params := newRandomCTParams(curve, bitsize)
		alice, err := GenerateKey(params)
		require.Nil(t, err)
		bob, err := GenerateKey(params)
		require.Nil(t, err)
		msg := new(big.Int).SetUint64(100)
		require.Nil(t, err)

		ct, err := EncryptTransfer(params, &alice.PublicKey, &bob.PublicKey, msg.Bytes())
		require.Nil(t, err)

		// check pk * r == x.
		aliceX := new(utils.ECPoint).ScalarMult(new(utils.ECPoint).SetFromPublicKey(&alice.PublicKey), ct.R)
		require.Equal(t, true, aliceX.Equal(ct.X1), "alice pk*r not equal")

		bobX := new(utils.ECPoint).ScalarMult(new(utils.ECPoint).SetFromPublicKey(&bob.PublicKey), ct.R)
		require.Equal(t, true, bobX.Equal(ct.X2), "bob pk*r not equal")
	}

}
