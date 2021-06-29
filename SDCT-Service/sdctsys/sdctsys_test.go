package sdctsys

import (
	"crypto/elliptic"
	"errors"
	"math/big"
	"testing"

	"github.com/sdct/curve"
	"github.com/sdct/proof"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSDCTSystem(t *testing.T) {
	curves := []elliptic.Curve{curve.BN256(), curve.S256()}
	bitsizes := []int{16, 32, 64}

	for _, curve := range curves {
		for _, bitsize := range bitsizes {
			params := proof.NewRandomAggRangeParams(curve, bitsize, 2)

			cases := []struct {
				aliceAmount *big.Int
				bobAmount   *big.Int
				send        *big.Int
				expect      bool
			}{
				{
					big.NewInt(100),
					big.NewInt(100),
					big.NewInt(50),
					true,
				},
				{
					big.NewInt(50),
					big.NewInt(50),
					big.NewInt(51),
					false,
				},
				{
					big.NewInt(0),
					big.NewInt(50),
					big.NewInt(51),
					false,
				},
			}

			for _, c := range cases {
				testsdctctx(t, params, c.aliceAmount, c.bobAmount, c.send, c.expect)
			}

		}
	}

}

func testsdctctx(t *testing.T, params proof.AggRangeParams, senderAmount, receiver, amount *big.Int, expect bool) {
	alice := CreateTestAccount(params, "alice", senderAmount)
	bob := CreateTestAccount(params, "bob", receiver)

	token := big.NewInt(0)
	ctx, err := CreateConfidentialTx(params, alice, &bob.sk.PublicKey, amount, token)
	require.Nil(t, err, "generate confidential tx failed", err)

	assert.Equal(t, expect, VerifyConfidentialTx(params, ctx), "confidential tx verify failed")
}

func BenchmarkCreateConfidentialTxBN256(b *testing.B) {
	params, alice, bob, amount := getTestConfig()
	token := big.NewInt(0)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if _, err := CreateConfidentialTx(params, alice, &bob.sk.PublicKey, amount, token); err != nil {
			b.Fatal(err)
		}
	}

}

func BenchmarkVerifyConfidentialTxBN256(b *testing.B) {
	params, alice, bob, amount := getTestConfig()
	token := big.NewInt(0)
	ctx, err := CreateConfidentialTx(params, alice, &bob.sk.PublicKey, amount, token)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !VerifyConfidentialTx(params, ctx) {
			b.Fatal(errors.New("failed"))
		}
	}
}

func getTestConfig() (proof.AggRangeParams, *Account, *Account, *big.Int) {
	senderAmount := new(big.Int).SetUint64(10000)
	receiver := new(big.Int).SetUint64(20000)
	amount := new(big.Int).SetUint64(9000)
	curve := curve.BN256()
	bitsize := 32
	params := proof.NewRandomAggRangeParams(curve, bitsize, 2)

	alice := CreateTestAccount(params, "alice", senderAmount)
	bob := CreateTestAccount(params, "bob", receiver)

	return params, alice, bob, amount
}
