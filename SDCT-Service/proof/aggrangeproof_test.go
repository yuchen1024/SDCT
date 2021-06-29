package proof

import (
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/sdct/curve"
	"github.com/sdct/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAggRangeProof(t *testing.T) {
	aggsize := 2
	cases := []struct {
		curve   elliptic.Curve
		bitsize int
		v       []*big.Int
		expect  bool
	}{
		{
			curve:   curve.BN256(),
			bitsize: 32,
			v:       []*big.Int{big.NewInt(0), big.NewInt(0)},
			expect:  true,
		},
		{
			curve:   curve.BN256(),
			bitsize: 32,
			v:       []*big.Int{utils.BiggestInt(32), utils.BiggestInt(32)},
			expect:  true,
		},
		{
			curve:   curve.BN256(),
			bitsize: 32,
			v:       []*big.Int{big.NewInt(-1), big.NewInt(0)},
			expect:  false,
		},
		{
			curve:   curve.BN256(),
			bitsize: 32,
			v:       []*big.Int{big.NewInt(0), big.NewInt(-1)},
			expect:  false,
		},
		{
			curve:   curve.BN256(),
			bitsize: 64,
			v:       []*big.Int{big.NewInt(0), big.NewInt(0)},
			expect:  true,
		},
		{
			curve:   curve.BN256(),
			bitsize: 64,
			v:       []*big.Int{utils.BiggestInt(64), utils.BiggestInt(64)},
			expect:  true,
		},
		{
			curve:   curve.BN256(),
			bitsize: 64,
			v:       []*big.Int{big.NewInt(0), big.NewInt(-1)},
			expect:  false,
		},
		{
			curve:   curve.BN256(),
			bitsize: 64,
			v:       []*big.Int{big.NewInt(-1), big.NewInt(0)},
			expect:  false,
		},
		{
			curve:   curve.BN256(),
			bitsize: 16,
			v:       []*big.Int{big.NewInt(0), big.NewInt(0)},
			expect:  true,
		},
		{
			curve:   curve.BN256(),
			bitsize: 16,
			v:       []*big.Int{utils.BiggestInt(16), utils.BiggestInt(16)},
			expect:  true,
		},
		{
			curve:   curve.BN256(),
			bitsize: 16,
			v:       []*big.Int{big.NewInt(-1), big.NewInt(0)},
			expect:  false,
		},
		{
			curve:   curve.BN256(),
			bitsize: 16,
			v:       []*big.Int{big.NewInt(0), big.NewInt(-1)},
			expect:  false,
		},
		{
			curve:   curve.S256(),
			bitsize: 16,
			v:       []*big.Int{big.NewInt(0), big.NewInt(0)},
			expect:  true,
		},
		{
			curve:   curve.S256(),
			bitsize: 16,
			v:       []*big.Int{utils.BiggestInt(16), utils.BiggestInt(16)},
			expect:  true,
		},
		{
			curve:   curve.S256(),
			bitsize: 16,
			v:       []*big.Int{big.NewInt(0), big.NewInt(-1)},
			expect:  false,
		},
		{
			curve:   curve.S256(),
			bitsize: 16,
			v:       []*big.Int{big.NewInt(-1), big.NewInt(0)},
			expect:  false,
		},
		{
			curve:   curve.S256(),
			bitsize: 16,
			v:       []*big.Int{big.NewInt(-1), big.NewInt(-1)},
			expect:  false,
		},
	}

	for _, c := range cases {
		testAggRangeProof(t, c.curve, c.bitsize, aggsize, c.v, c.expect)
	}
}

func testAggRangeProof(t *testing.T, curve elliptic.Curve, bitsize, aggsize int, v []*big.Int, expect bool) {
	params := newRandomAggRangeParams(curve, bitsize, aggsize)
	p, r := newRandomCommitmentsAggRangeProof(params, v)

	proof, err := GenerateAggRangeProof(params, v, r)
	require.Nil(t, err, "generate agg range proof failed")

	assert.Equal(t, expect, VerifyAggRangeProof(params, p, proof), "normal agg range proof verify not expect")
	assert.Equal(t, expect, OptimizedVerifyAggRangeProof(params, p, proof), "optimize agg range proof verify not expect")
	//for simple fake proof.
	if expect {
		proof.t.Sub(proof.t, utils.One)
		assert.Equal(t, false, VerifyAggRangeProof(params, p, proof), "invalid agg range proof pass normal verify")
		assert.Equal(t, false, OptimizedVerifyAggRangeProof(params, p, proof), "invalid agg range proof pass optimized verify")
	}
}

func newRandomCommitmentsAggRangeProof(params AggRangeParams, v []*big.Int) ([]*utils.ECPoint, []*big.Int) {
	g := params.G()
	h := params.H()

	p := make([]*utils.ECPoint, 0)
	r := make([]*big.Int, 0)

	for _, iv := range v {
		random, err := rand.Int(rand.Reader, params.Curve().Params().N)
		if err != nil {
			panic(err)
		}

		point := new(utils.ECPoint).ScalarMult(g, iv)
		point.Add(point, new(utils.ECPoint).ScalarMult(h, random))

		p = append(p, point)
		r = append(r, random)
	}

	return p, r
}

func BenchmarkBN256AggRangeProofVerifyNormal(b *testing.B) {
	v := []*big.Int{big.NewInt(100), big.NewInt(100)}
	benchmarkAggRangeProofVerify(b, curve.BN256(), 32, 2, v, VerifyAggRangeProof)
}

func BenchmarkBN256AggRangeProofVerifyOptimized(b *testing.B) {
	v := []*big.Int{big.NewInt(100), big.NewInt(100)}
	benchmarkAggRangeProofVerify(b, curve.BN256(), 32, 2, v, OptimizedVerifyAggRangeProof)
}

type aggRangeProofVerifyFunc func(params AggRangeParams, v []*utils.ECPoint, proof *AggRangeProof) bool

func benchmarkAggRangeProofVerify(b *testing.B, curve elliptic.Curve, bitsize, aggsize int, v []*big.Int, vf aggRangeProofVerifyFunc) {
	b.StopTimer()
	params := newRandomAggRangeParams(curve, bitsize, aggsize)
	p, r := newRandomCommitmentsAggRangeProof(params, v)

	proof, err := GenerateAggRangeProof(params, v, r)
	require.Nil(b, err, "generate agg range proof failed")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		if !vf(params, p, proof) {
			b.Fatal("agg range proof verify failed")
		}
	}
}
