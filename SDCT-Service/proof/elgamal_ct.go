package proof

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"math/big"

	log "github.com/inconshreveable/log15"
	"github.com/sdct/utils"
)

// CTParams contains all params to encrypt/decrypt/generae key/msg.
type CTParams interface {
	H() *utils.ECPoint
	Bitsize() int
	KeyParams

	// for auth verify
	Priv() *ecdsa.PrivateKey
}

type ctParams struct {
	h, g    *utils.ECPoint
	bitsize int
	priv    *ecdsa.PrivateKey
}

func newRandomCTParams(curve elliptic.Curve, bitsize int) CTParams {
	ctp := ctParams{}
	ctp.h = utils.NewRandomECPoint(curve)
	ctp.g = utils.NewRandomECPoint(curve)
	ctp.bitsize = bitsize
	priv := MustGenerateKey(&ctp)
	ctp.priv = priv

	return &ctp
}

func (ctp *ctParams) G() *utils.ECPoint {
	return ctp.g
}

func (ctp *ctParams) H() *utils.ECPoint {
	return ctp.h
}

func (ctp *ctParams) Bitsize() int {
	return ctp.bitsize
}

func (ctp *ctParams) Curve() elliptic.Curve {
	return ctp.h.Curve
}

func (ctp *ctParams) Priv() *ecdsa.PrivateKey {
	return ctp.priv
}

// KeyParams contains params to genrate key.
type KeyParams interface {
	G() *utils.ECPoint
	Curve() elliptic.Curve
}

type keyParams struct {
	g *utils.ECPoint
}

func (k *keyParams) G() *utils.ECPoint {
	return k.g
}

func (k *keyParams) Curve() elliptic.Curve {
	return k.g.Curve
}

func newRandomKeyParams(curve elliptic.Curve) KeyParams {
	g := utils.NewRandomECPoint(curve)

	return &keyParams{g}
}

// CTEncPoint respresents encrypted ct tx point on curve.
type CTEncPoint struct {
	X, Y *utils.ECPoint
}

// Sub .
func (c *CTEncPoint) Sub(first, second *CTEncPoint) *CTEncPoint {
	ecPoints := CTEncPoint{}

	ecPoints.X = new(utils.ECPoint).Sub(first.X, second.X)
	ecPoints.Y = new(utils.ECPoint).Sub(first.Y, second.Y)
	return &ecPoints
}

// Copy .
func (c *CTEncPoint) Copy() *CTEncPoint {
	ecPoints := CTEncPoint{}

	ecPoints.X = c.X.Copy()
	ecPoints.Y = c.Y.Copy()
	return &ecPoints
}

// PublicKey represents a public key used in twisted elgamal.
type PublicKey struct {
	elliptic.Curve

	// Point on curve.
	X, Y *big.Int
}

// PrivateKey represents a private key used in twisted elgamal.
type PrivateKey struct {
	// sk
	D *big.Int

	PublicKey
}

// MRTwistedELGamalCTPub represent public points in MRTwistedELGamalCT tx.
type MRTwistedELGamalCTPub struct {
	// X1=pk1*r; X2=pk2*r; X3=pka*r.
	X1, X2, X3 *utils.ECPoint
	// Y = g*r + h*m.
	Y *utils.ECPoint
}

func (mrp *MRTwistedELGamalCTPub) First() *CTEncPoint {
	return &CTEncPoint{
		X: mrp.X1,
		Y: mrp.Y,
	}
}

// MRTwistedELGamalCT defines the structure of 2R1M ciphertext (MR denotes multiple recipients).
type MRTwistedELGamalCT struct {
	MRTwistedELGamalCTPub
	R      *big.Int
	EncMsg []byte
}

func (mr *MRTwistedELGamalCT) Pub() *MRTwistedELGamalCTPub {
	return &mr.MRTwistedELGamalCTPub
}

// TwistedELGamalCT respresents a encrypted transaction encoded in twisted-elgamal format.
type TwistedELGamalCT struct {
	// X, Y
	CTEncPoint
	// Random number used in encryption.
	R *big.Int
	// encrypted msg.
	EncMsg []byte
}

// CopyPublicPoint copy public encrypted ec point.
func (ct *TwistedELGamalCT) CopyPublicPoint() *CTEncPoint {
	ecPoints := CTEncPoint{}

	ecPoints.X = ct.X.Copy()
	ecPoints.Y = ct.Y.Copy()
	return &ecPoints
}

// Sub return new instance = first - second.
func (ct *TwistedELGamalCT) Sub(first, second *TwistedELGamalCT) *TwistedELGamalCT {
	ct.X = new(utils.ECPoint).Sub(first.X, second.X)
	ct.Y = new(utils.ECPoint).Sub(first.Y, second.Y)
	ct.R = nil
	ct.EncMsg = []byte{}

	return ct
}

// Copy returns new instance.
func (ct *TwistedELGamalCT) Copy() *TwistedELGamalCT {
	newCT := TwistedELGamalCT{}
	newCT.X = ct.X.Copy()
	newCT.Y = ct.Y.Copy()
	if ct.R != nil {
		newCT.R = new(big.Int).Set(ct.R)
	}
	if len(ct.EncMsg) > 0 {
		newCT.EncMsg = make([]byte, len(ct.EncMsg))
		copy(newCT.EncMsg, ct.EncMsg)
	}

	return &newCT
}

// copy from crypto/ecdsa.
// randFieldElement returns a random element of the field underlying the given
// curve using the procedure given in [NSA] A.2.1.
func randFieldElement(c elliptic.Curve, rand io.Reader) (k *big.Int, err error) {
	params := c.Params()
	b := make([]byte, params.BitSize/8+8)
	_, err = io.ReadFull(rand, b)
	if err != nil {
		return
	}

	k = new(big.Int).SetBytes(b)
	n := new(big.Int).Sub(params.N, utils.One)
	k.Mod(k, n)
	k.Add(k, utils.One)
	return
}

// MustGenerateKey generates a key pair and panic if err.
// Warn: test purpose only.
func MustGenerateKey(params KeyParams) *ecdsa.PrivateKey {
	key, err := GenerateKey(params)
	if err != nil {
		panic(err)
	}

	return key
}

// GenerateKey generates key pair.
// Warn: The g point from global params is used for generating key pair.
func GenerateKey(params KeyParams) (priv *ecdsa.PrivateKey, err error) {
	curve := params.Curve()
	g := params.G()

	k, err := randFieldElement(curve, rand.Reader)
	if err != nil {
		return
	}

	priv = new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = curve
	priv.D = k
	// Waring: x, y == g * sk.
	priv.PublicKey.X, priv.PublicKey.Y = curve.ScalarMult(g.X, g.Y, k.Bytes())
	return
}

// HexToKey returns key pair by hex.
func HexToKey(hexKey string, params KeyParams) (priv *ecdsa.PrivateKey, err error) {
	b, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, err
	}

	g := params.G()
	priv = new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = params.Curve()
	priv.D = new(big.Int).SetBytes(b)
	priv.PublicKey.X, priv.PublicKey.Y = params.Curve().ScalarMult(g.X, g.Y, b)
	return
}

// EncryptOnChain encrypts msg with random 0.
func EncryptOnChain(params CTParams, pk *ecdsa.PublicKey, msg []byte) (*TwistedELGamalCT, error) {
	msgBit := new(big.Int).SetBytes(msg)

	// Create ct instance.
	ct := new(TwistedELGamalCT)
	ct.X = utils.NewEmptyECPoint(pk.Curve)
	ct.Y = utils.NewEmptyECPoint(pk.Curve)

	// set curve
	curve := pk.Curve

	// set random 0.
	r := new(big.Int).SetUint64(0)

	// for sigma proof purpose.
	ct.R = new(big.Int).Set(r)
	ct.EncMsg = make([]byte, len(msg))
	copy(ct.EncMsg, msg)

	// compute pk * r.(pk ^ r)
	ct.X.SetFromPublicKey(pk)
	ct.X.ScalarMult(ct.X, r)

	// compute g * r + h * m.
	h := params.H()
	ct.Y.ScalarMult(h, msgBit)
	log.Debug("encrypt msg(h^m)", "x", ct.Y.X, "y", ct.Y.Y)
	g := params.G()
	s2X, s2Y := curve.ScalarMult(g.X, g.Y, r.Bytes())
	ct.Y.Add(ct.Y, utils.NewECPoint(s2X, s2Y, curve))

	return ct, nil

}

// EncryptTransfer encrypts msg by two different pk but with same random.
func EncryptTransfer(params CTParams, sender, receiver *ecdsa.PublicKey, msg []byte) (*MRTwistedELGamalCT, error) {
	curve := sender.Curve
	// get random.
	r, err := rand.Int(rand.Reader, curve.Params().N)
	if err != nil {
		return nil, err
	}
	twSender, err := EncryptWithRandom(params, sender, msg, r)
	if err != nil {
		return nil, err
	}
	twReceiver, err := EncryptWithRandom(params, receiver, msg, r)
	if err != nil {
		return nil, err
	}
	twAuth, err := EncryptWithRandom(params, &params.Priv().PublicKey, msg, r)
	if err != nil {
		return nil, err
	}

	if !twReceiver.Y.Equal(twSender.Y) || !twReceiver.Y.Equal(twAuth.Y) {
		return nil, errors.New("twisted elgamal y point not equal")
	}

	return &MRTwistedELGamalCT{
		MRTwistedELGamalCTPub: MRTwistedELGamalCTPub{
			X1: twSender.X,
			X2: twReceiver.X,
			X3: twAuth.X,
			Y:  twSender.Y,
		},
		R:      r,
		EncMsg: msg,
	}, nil
}

// Encrypt encrypts msg.
func Encrypt(params CTParams, pk *ecdsa.PublicKey, msg []byte) (*TwistedELGamalCT, error) {
	curve := pk.Curve
	// get random.
	r, err := rand.Int(rand.Reader, curve.Params().N)
	if err != nil {
		return nil, err
	}

	return EncryptWithRandom(params, pk, msg, r)
}

// EncryptWithRandom encrypts msg in twisted elgamal way.
func EncryptWithRandom(params CTParams, pk *ecdsa.PublicKey, msg []byte, r *big.Int) (*TwistedELGamalCT, error) {
	msgBit := new(big.Int).SetBytes(msg)

	// Create ct instance.
	ct := new(TwistedELGamalCT)
	ct.X = utils.NewEmptyECPoint(pk.Curve)
	ct.Y = utils.NewEmptyECPoint(pk.Curve)

	// set curve
	curve := pk.Curve

	// for sigma proof purpose.
	ct.R = new(big.Int).Set(r)
	ct.EncMsg = make([]byte, len(msg))
	copy(ct.EncMsg, msg)

	// compute pk * r.(pk ^ r)
	ct.X.SetFromPublicKey(pk)
	ct.X.ScalarMult(ct.X, r)

	// compute g * r + h * m.
	h := params.H()
	ct.Y.ScalarMult(h, msgBit)
	g := params.G()
	s2X, s2Y := curve.ScalarMult(g.X, g.Y, r.Bytes())
	ct.Y.Add(ct.Y, utils.NewECPoint(s2X, s2Y, curve))

	return ct, nil
}

// Decrypt decrypts msg in twisted elgamal way.
func Decrypt(params CTParams, sk *ecdsa.PrivateKey, ct *CTEncPoint) []byte {
	encodedMsg := getEncryptedMsg(sk, ct.Copy())
	return decryptEncodedMsg(params, encodedMsg)
}

func getEncryptedMsg(sk *ecdsa.PrivateKey, ct *CTEncPoint) *utils.ECPoint {
	// get public curve.
	curve := sk.PublicKey.Curve
	// compute the inverse of sk.
	skInverse := new(big.Int).ModInverse(sk.D, curve.Params().N)
	// compute X * skInverse(X ^ skInverse) == g * r.
	ct.X.ScalarMult(ct.X, skInverse)
	// use ct.Y - ct.X Point to get encoded msg.
	// get Affine negation formulas of ct.Y.
	encodedMsg := ct.Y.Sub(ct.Y, ct.X)

	return encodedMsg
}

// decryptencodedmsg decrypts and returns original bytes of msg.
// encodeMsg = h * m
func decryptEncodedMsg(params CTParams, encodeMsg *utils.ECPoint) []byte {
	h := params.H().Copy()

	// encode zero.
	if encodeMsg.X.Cmp(utils.Zero) == 0 && encodeMsg.Y.Cmp(utils.Zero) == 0 {
		return []byte{}
	}
	if encodeMsg.Equal(h) {
		return utils.One.Bytes()
	}

	return ShanksDlog(h, encodeMsg, params.Bitsize(), 7).Bytes()
}

// Refresh Ciphertext refreshing algorithm: output a fresh ciphertext for the message encrypted in CT.
func Refresh(params CTParams, sk *ecdsa.PrivateKey, ct *CTEncPoint) (*TwistedELGamalCT, error) {
	// get encrypted msg.
	encodedMsg := getEncryptedMsg(sk, ct.Copy())
	// encrypt.
	pk := sk.PublicKey
	refreshCT := new(TwistedELGamalCT)
	refreshCT.X = utils.NewEmptyECPoint(pk.Curve)
	refreshCT.Y = utils.NewEmptyECPoint(pk.Curve)

	// set curve
	curve := pk.Curve

	// get random.
	r, err := rand.Int(rand.Reader, curve.Params().N)
	if err != nil {
		return nil, err
	}
	// for sigma proof purpose.
	refreshCT.R = new(big.Int).Set(r)
	// don't set the msg.

	// compute pk * r.(pk ^ r)
	refreshCT.X.SetFromPublicKey(&pk)
	refreshCT.X.ScalarMult(refreshCT.X, r)

	// set h * m.
	refreshCT.Y = encodedMsg.Copy()
	// compute g * r.
	g := params.G()
	s2X, s2Y := curve.ScalarMult(g.X, g.Y, r.Bytes())
	// compute g * r + h * m.
	refreshCT.Y.Add(refreshCT.Y, utils.NewECPoint(s2X, s2Y, curve))

	return refreshCT, nil
}
