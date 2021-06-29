package curve

import (
	"testing"

	"github.com/sdct/utils"
	"github.com/stretchr/testify/assert"
)

func TestBN128Curve(t *testing.T) {
	c := BN128{}
	n := c.Params().N
	x, y := c.ScalarBaseMult(n.Bytes())

	assert.Equal(t, x, y)
	assert.Equal(t, x.Cmp(utils.Zero), 0)

	assert.True(t, c.IsOnCurve(c.Params().Gx, c.Params().Gy))
}
