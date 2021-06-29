package curve

import (
	"crypto/elliptic"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
)

// Marshaler can marshal a point to a big.int.
type Marshaler interface {
	Marshal(x, y *big.Int) []byte
}

// BN128 implements elliptic.curve.
// elliptic curve y²=x³+3
type BN128 struct {
}

func fromString(s string) *big.Int {
	r, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("invalid hex in source file: " + s)
	}
	return r
}

// Params implements curve.
func (bn128 *BN128) Params() *elliptic.CurveParams {
	params := new(elliptic.CurveParams)
	// x, y order.
	params.P = fromString("21888242871839275222246405745257275088696311157297823662689037894645226208583")
	// scalar order
	params.N = fromString("21888242871839275222246405745257275088548364400416034343698204186575808495617")
	params.Gx = new(big.Int).SetUint64(1)
	params.Gy = new(big.Int).SetUint64(2)
	params.B = new(big.Int).SetUint64(3)
	params.BitSize = 256

	return params
}

// IsOnCurve implements curve.
func (bn128 *BN128) IsOnCurve(x, y *big.Int) bool {
	// point to g1 will try to convert x,y to g1 point and
	// will check it on curve or not.
	if _, err := PointToG1(x, y); err != nil {
		return false
	}
	return true
}

// Add implements curve.
func (bn128 *BN128) Add(x1, y1, x2, y2 *big.Int) (x, y *big.Int) {
	a := MustPointToG1(x1, y1)
	b := MustPointToG1(x2, y2)
	p := new(bn256.G1).Add(a, b)

	return MustG1ToPoint(p)
}

// Double implements curve.
func (bn128 *BN128) Double(x1, y1 *big.Int) (x, y *big.Int) {
	p := MustPointToG1(x1, y1)
	p.ScalarMult(p, new(big.Int).SetUint64(2))

	return MustG1ToPoint(p)
}

// ScalarMult implements curve.
func (bn128 *BN128) ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int) {
	p := MustPointToG1(x1, y1)
	newP := new(bn256.G1).ScalarMult(p, new(big.Int).SetBytes(k))
	x, y = MustG1ToPoint(newP)
	return
}

// ScalarBaseMult implements curve.
func (bn128 *BN128) ScalarBaseMult(k []byte) (x, y *big.Int) {
	p := new(bn256.G1).ScalarBaseMult(new(big.Int).SetBytes(k))
	return MustG1ToPoint(p)
}

// Marshal marshal x, y point to byte slice.
func (bn128 *BN128) Marshal(x, y *big.Int) []byte {
	return MustPointToG1(x, y).Marshal()
}

// Unmarshal sets e to the result of converting the output of Marshal back into
// a group element and then returns e.
func (bn128 *BN128) Unmarshal(m []byte) (x, y *big.Int) {
	p := new(bn256.G1)
	_, err := p.Unmarshal(m)
	if err != nil {
		panic(err)
	}

	return MustG1ToPoint(p)
}

// MustG1ToPoint converts g1 struct to x,y point. panic if error.
func MustG1ToPoint(p *bn256.G1) (x, y *big.Int) {
	x, y, err := G1ToPoint(p)
	if err != nil {
		panic(err)
	}

	return
}

// G1ToPoint converts g1 struct to x,y point.
func G1ToPoint(p *bn256.G1) (x, y *big.Int, err error) {
	data := p.Marshal()
	const numBytes = 256 / 8
	if len(data) < 2*numBytes {
		err = errors.New("bn256: not enough data")
		return
	}

	x = new(big.Int).SetBytes(data[0:numBytes])
	y = new(big.Int).SetBytes(data[numBytes:])
	return
}

// MustPointToG1 converts point x,y to G1 struct. panic if error.
// based on ethereum contracts op.
func MustPointToG1(x, y *big.Int) *bn256.G1 {
	p, err := PointToG1(x, y)
	if err != nil {
		panic(err)
	}

	return p
}

// PointToG1 convert point x,y to G1 struct.
// based on ethereum contracts op.
func PointToG1(x, y *big.Int) (*bn256.G1, error) {
	// make sure each value is a 256-bit number.(32 byte)
	xBytes := common.BytesToHash(x.Bytes()).Bytes()
	yBytes := common.BytesToHash(y.Bytes()).Bytes()

	data := make([]byte, 64)
	copy(data[0:32], xBytes)
	copy(data[32:], yBytes)

	p := new(bn256.G1)
	if _, err := p.Unmarshal(data); err != nil {
		return nil, err
	}

	return p, nil
}
