package proof

import (
	"math/big"
	"testing"

	"github.com/sdct/utils"
	"github.com/stretchr/testify/assert"
)

func TestBuildHashMap(t *testing.T) {
	params := DAggRangeProofParams32()
	buildHashMap(params.H(), 32, 7, 4)
}

func BenchmarkFindPoint(b *testing.B) {
	params := DAggRangeProofParams32()
	h := params.H()

	// biggest number in [0, 2^32)
	r := big.NewInt(2<<31 - 1)
	dst := new(utils.ECPoint).ScalarMult(h, r)
	// make sure the map exists
	loadHashMap(32, 7)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dec := ShanksDlog(h, dst, 32, 7)
		assert.Equal(b, r, dec)
	}
}
