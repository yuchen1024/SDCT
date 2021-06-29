package proof

import (
	"crypto/elliptic"
	"math/big"
	"testing"

	"github.com/sdct/curve"
	"github.com/sdct/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEqualProof(t *testing.T) {
	curve := curve.BN256()
	bitsize := 32
	params := newRandomCTParams(curve, bitsize)
	key, err := GenerateKey(params)
	require.Nil(t, err)

	tests := []struct {
		Msg    []byte
		Custom []*big.Int
	}{
		{
			Msg:    new(big.Int).SetUint64(2).Bytes(),
			Custom: []*big.Int{big.NewInt(1)},
		},
		{
			Msg:    new(big.Int).SetUint64(2).Bytes(),
			Custom: []*big.Int{big.NewInt(1), big.NewInt(2)},
		},
	}

	for _, test := range tests {
		ct, err := Encrypt(params, &key.PublicKey, test.Msg)
		require.Nil(t, err)

		amount := new(big.Int).SetBytes(test.Msg)
		proof, err := GenerateEqualProof(params, amount, ct.CopyPublicPoint(), key, test.Custom...)
		require.Nil(t, err)

		pb := new(utils.ECPoint).SetFromPublicKey(&key.PublicKey)
		actual := VerifyEqualProof(params, ct.CopyPublicPoint(), amount, pb, proof, test.Custom...)
		assert.True(t, actual)
	}
}

func TestDLESigmaProof(t *testing.T) {
	curves := []elliptic.Curve{curve.BN256(), curve.S256()}
	bitsizes := []int{16, 32, 64}
	for _, curve := range curves {
		for _, bitsize := range bitsizes {
			params := newRandomCTParams(curve, bitsize)
			key, err := GenerateKey(params)
			require.Nil(t, err)

			tests := []struct {
				MsgA   []byte
				MsgB   []byte
				Expect bool
				Custom []*big.Int
			}{
				{
					MsgA:   new(big.Int).SetUint64(2).Bytes(),
					MsgB:   new(big.Int).SetUint64(2).Bytes(),
					Expect: true,
					Custom: []*big.Int{big.NewInt(1)},
				},
				{
					MsgA:   new(big.Int).SetUint64(2).Bytes(),
					MsgB:   new(big.Int).SetUint64(1).Bytes(),
					Expect: false,
					Custom: []*big.Int{big.NewInt(1), big.NewInt(2)},
				},
			}

			for _, test := range tests {
				ct1, err := Encrypt(params, &key.PublicKey, test.MsgA)
				require.Nil(t, err)
				ct2, err := Encrypt(params, &key.PublicKey, test.MsgB)
				require.Nil(t, err)
				require.NotEqual(t, ct1.X.X, ct2.X.X)

				proof, err := GenerateDLESigmaProof(params, ct1.CopyPublicPoint(), ct2.CopyPublicPoint(), key, test.Custom...)
				require.Nil(t, err)

				assert.Equal(t, test.Expect, VerifyDLESigmaProof(params, ct1.CopyPublicPoint(), ct2.CopyPublicPoint(), &key.PublicKey, proof, test.Custom...))
			}
		}
	}
}
