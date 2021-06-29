package proof

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"

	log "github.com/inconshreveable/log15"
	"github.com/sdct/curve"
	"github.com/sdct/utils"
)

// AggRangeParams contains all params to generate/verify aggregate range proof.
type AggRangeParams interface {
	Bitsize() int
	Aggsize() int
	Curve() elliptic.Curve
	GV() *utils.GeneratorVector
	HV() *utils.GeneratorVector
	U() *utils.ECPoint
	G() *utils.ECPoint
	H() *utils.ECPoint
	Priv() *ecdsa.PrivateKey
}

// Reserve switches value of g/h.
func Reserve(params AggRangeParams) AggRangeParams {
	p, ok := params.(*aggRangeParams)
	if !ok {
		panic("")
	}

	ns := p.Copy()
	ns.g = p.h.Copy()
	ns.h = p.g.Copy()
	return ns
}

type aggRangeParams struct {
	gv, hv           *utils.GeneratorVector
	u, g, h          *utils.ECPoint
	bitsize, aggsize int

	// for auth verify
	priv *ecdsa.PrivateKey
}

func newRandomAggRangeParams(curve elliptic.Curve, bitsize, aggsize int) AggRangeParams {
	arp := aggRangeParams{}
	arp.gv = utils.NewRandomGeneratorVector(curve, bitsize*aggsize)
	arp.hv = utils.NewRandomGeneratorVector(curve, bitsize*aggsize)
	arp.u = utils.NewRandomECPoint(curve)
	arp.g = utils.NewRandomECPoint(curve)
	arp.h = utils.NewRandomECPoint(curve)
	arp.bitsize = bitsize
	arp.aggsize = aggsize
	arp.priv = MustGenerateKey(&arp)

	return &arp
}

// Const params for contracts.
const (
	Bitsize = 32
	N       = 5
	step    = 1
	LRsize  = N + step
)

// DAggRangeProofParams32 returns default params of 32 bit.
func DAggRangeProofParams32() AggRangeParams {
	return DAggRangeProofParamsWithBitsize(32)
}

// DAggRangeProofParamsWithBitsize return default params with bitsize.
func DAggRangeProofParamsWithBitsize(bitsize int) AggRangeParams {
	curve := curve.BN256()
	aggsize := 2
	g := "g generator of twisted elg"
	gpoint := utils.NewECPointByString(g, curve)

	h := "h generator of twisted elg"
	hpoint := utils.NewECPointByString(h, curve)

	gv := utils.NewDefaultGV(curve, bitsize*aggsize)
	hv := utils.NewDefaultHV(curve, bitsize*aggsize)

	u := "u generator of innerproduct"
	upoint := utils.NewECPointByString(u, curve)

	return NewAggRangeParams(gv, hv, upoint, gpoint, hpoint, bitsize, aggsize)
}

// NewAggRangeParams returns a new instance of aggregate range proof params.
func NewAggRangeParams(gv, hv *utils.GeneratorVector, u, g, h *utils.ECPoint, bitsize, aggsize int) AggRangeParams {
	arp := aggRangeParams{}
	arp.gv = gv.Copy()
	arp.hv = hv.Copy()
	arp.u = u.Copy()
	arp.g = g.Copy()
	arp.h = h.Copy()
	arp.bitsize = bitsize
	arp.aggsize = aggsize
	priv := MustGenerateKey(&arp)
	arp.priv = priv

	return &arp
}

// NewRandomAggRangeParams generates a random params for aggregate range proof.
func NewRandomAggRangeParams(curve elliptic.Curve, bitsize, aggsize int) AggRangeParams {
	return newRandomAggRangeParams(curve, bitsize, aggsize)
}

func (arp *aggRangeParams) Bitsize() int {
	return arp.bitsize
}

func (arp *aggRangeParams) Aggsize() int {
	return arp.aggsize
}

func (arp *aggRangeParams) Curve() elliptic.Curve {
	return arp.u.Curve
}

func (arp *aggRangeParams) GV() *utils.GeneratorVector {
	return arp.gv
}

func (arp *aggRangeParams) HV() *utils.GeneratorVector {
	return arp.hv
}

func (arp *aggRangeParams) U() *utils.ECPoint {
	return arp.u
}

func (arp *aggRangeParams) G() *utils.ECPoint {
	return arp.g
}

func (arp *aggRangeParams) H() *utils.ECPoint {
	return arp.h
}

func (arp *aggRangeParams) Priv() *ecdsa.PrivateKey {
	return arp.priv
}

func (arp *aggRangeParams) Copy() *aggRangeParams {
	ns := aggRangeParams{}
	ns.gv = arp.gv.Copy()
	ns.hv = arp.hv.Copy()
	ns.u = arp.u.Copy()
	ns.g = arp.g.Copy()
	ns.h = arp.h.Copy()
	ns.bitsize = arp.bitsize
	ns.aggsize = arp.aggsize
	ns.priv = arp.priv

	return &ns
}

// AggRangeProof aggregates multi range proofs.
type AggRangeProof struct {
	A, S *utils.ECPoint

	T1, T2 *utils.ECPoint

	t, tx, u *big.Int

	ipProof *IPProof
}

type aggRangeProofInput struct {
	points [12]*big.Int
	scalar [5]*big.Int
	l, r   [2 * LRsize]*big.Int
}

// T returns t.
func (aggp *AggRangeProof) T() *big.Int {
	return aggp.t
}

// TX returns tx.
func (aggp *AggRangeProof) TX() *big.Int {
	return aggp.tx
}

// U returns u.
func (aggp *AggRangeProof) U() *big.Int {
	return aggp.u
}

// Len .
func (aggp *AggRangeProof) Len() int {
	return len(aggp.ipProof.l)
}

// Li .
func (aggp *AggRangeProof) Li(i int) *utils.ECPoint {
	return aggp.ipProof.l[i]
}

// L .
func (aggp *AggRangeProof) L() []*utils.ECPoint {
	return aggp.ipProof.l
}

// Ri .
func (aggp *AggRangeProof) Ri(i int) *utils.ECPoint {
	return aggp.ipProof.r[i]
}

// R .
func (aggp *AggRangeProof) R() []*utils.ECPoint {
	return aggp.ipProof.r
}

// AIP .
func (aggp *AggRangeProof) AIP() *big.Int {
	return aggp.ipProof.a
}

// BIP .
func (aggp *AggRangeProof) BIP() *big.Int {
	return aggp.ipProof.b
}

// ToSolidityInput format data to solidity input to test for contracts.
func (aggp *AggRangeProof) ToSolidityInput() *aggRangeProofInput {
	input := aggRangeProofInput{}
	input.points[0] = aggp.A.X
	input.points[1] = aggp.A.Y
	input.points[2] = aggp.S.X
	input.points[3] = aggp.S.Y
	input.points[4] = aggp.T1.X
	input.points[5] = aggp.T1.Y
	input.points[6] = aggp.T2.X
	input.points[7] = aggp.T2.Y

	input.scalar[0] = aggp.t
	input.scalar[1] = aggp.tx
	input.scalar[2] = aggp.u
	input.scalar[3] = aggp.ipProof.a
	input.scalar[4] = aggp.ipProof.b

	for i := 0; i < len(aggp.ipProof.l); i++ {
		input.l[2*i] = aggp.ipProof.l[i].X
		input.l[2*i+1] = aggp.ipProof.l[i].Y
		input.r[2*i] = aggp.ipProof.r[i].X
		input.r[2*i+1] = aggp.ipProof.r[i].Y
	}

	return &input
}

// GenerateAggRangeProof aggregates many(2 current) bullet proof together to reduce proof size.
func GenerateAggRangeProof(params AggRangeParams, v, random []*big.Int) (*AggRangeProof, error) {
	bitsize := params.Bitsize()
	aggsize := params.Aggsize()
	vectorSize := bitsize * aggsize
	n := params.Curve().Params().N

	if len(v) != aggsize {
		return nil, fmt.Errorf("witness v len %d not equal aggregate size %d", len(v), aggsize)
	}

	if len(random) != aggsize {
		return nil, fmt.Errorf("witness r len %d not equal aggregate size %d", len(random), aggsize)
	}

	alVector, err := utils.MultBitVector(v, bitsize)
	if err != nil {
		return nil, err
	}
	// <al, 2 ^ n> == v; ar == al - 1 ^ n.
	al := utils.NewFieldVector(alVector, n)
	ar := al.AllItemsSubOne()

	// pick a random number alpha.
	alpha, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}
	// compute commitment to al and ar.
	// commitA = g vector * al vector + h vector * ar vector + h point * alpha.
	gv := params.GV()
	hv := params.HV()
	h := params.H()
	g := params.G()
	commitA := utils.CommitTwoFieldVector(gv, hv, h, al, ar, alpha)

	// pick binding vector sl, sr.
	sl := utils.NewRandomFieldVector(n, vectorSize)
	sr := utils.NewRandomFieldVector(n, vectorSize)
	// pick another random number rho.
	rho, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}
	// computation same with commitA.
	commitB := utils.CommitTwoFieldVector(gv, hv, h, sl, sr, rho)

	// compute challenge y, z.
	y, err := utils.ComputeChallenge(n, commitA.X, commitA.Y, commitB.X, commitB.Y)
	if err != nil {
		return nil, err
	}
	z, err := utils.ComputeChallenge(n, commitB.X, commitB.Y, commitA.X, commitA.Y)
	if err != nil {
		return nil, err
	}

	// 2^mn.
	mn2 := utils.PowVector(utils.Two, n, vectorSize)
	// y^mn.
	ymn := utils.PowVector(y, n, vectorSize)
	// sr hadamard y^mn.
	rr1 := sr.Hadamard(ymn)
	// compute y^mn (ar + z * 1^mn) + z^(1+j) * (0^(j-1)*n || 2^n || 0^(m-j)*n). (j=[1, m])
	rr0 := ar.AddFieldVector(utils.RepeatItemVector(z, n, vectorSize))
	rr0 = rr0.Hadamard(ymn)
	n2 := mn2.SubFieldVector(0, bitsize)
	for j := 1; j <= aggsize; j++ {
		// 0^((j-1)*n)
		tmpz := utils.RepeatItemVector(utils.Zero, n, bitsize*(j-1))
		// 2^n
		tmpz = tmpz.Append(n2)
		// 0^((m-j)*n)
		tmpz = tmpz.Append(utils.RepeatItemVector(utils.Zero, n, (aggsize-j)*bitsize))
		if tmpz.Size() != vectorSize {
			return nil, fmt.Errorf("tmp z calculate failed expect len %d, actual len %d", vectorSize, tmpz.Size())
		}

		// z^(1+j)
		j1 := new(big.Int).SetUint64(uint64(1 + j))
		zj := new(big.Int).Exp(z, j1, n)
		rr0 = rr0.AddFieldVector(tmpz.Times(zj))
	}

	// compute t0, t1, t2.
	zSquare := new(big.Int).Mul(z, z)
	zSquare.Mod(zSquare, n)

	// compute t0(for check only).
	// t0 = <ll0, rr0>
	zNeg := new(big.Int).Neg(z)
	zNeg.Mod(zNeg, n)
	ll0 := al.AddFieldVector(utils.RepeatItemVector(zNeg, n, vectorSize))
	t0 := ll0.InnerProduct(rr0)
	t0.Mod(t0, n)

	// compute t1.
	// t1 = <ll0, rr1> + <ll1, rro>;
	t1 := ll0.InnerProduct(rr1)
	t1.Add(t1, sl.InnerProduct(rr0))
	t1.Mod(t1, n)

	// t2 = <sl, sr hadamard y^mn>
	// t2 = <ll1, rr1>
	t2 := sl.InnerProduct(rr1)
	t2.Mod(t2, n)

	// commit to t1, t2.
	// pick two random number.
	r1, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}
	r2, err := rand.Int(rand.Reader, n)
	if err != nil {
		return nil, err
	}
	T1 := new(utils.ECPoint).ScalarMult(h, r1)
	T1.Add(T1, new(utils.ECPoint).ScalarMult(g, t1))
	T2 := new(utils.ECPoint).ScalarMult(h, r2)
	T2.Add(T2, new(utils.ECPoint).ScalarMult(g, t2))

	// compute challenge x.
	x, err := utils.ComputeChallenge(n, T1.X, T1.Y, T2.X, T2.Y)
	if err != nil {
		return nil, err
	}
	x2 := new(big.Int).Exp(x, utils.Two, n)

	// compute l, r...
	// l = al - z*1^mn + sl*x.
	l := sl.Times(x)
	l = l.AddFieldVector(ll0)
	// r.
	r := rr0.AddFieldVector(rr1.Times(x))
	t := l.InnerProduct(r)
	t.Mod(t, n)

	// compute r2 * x^2 + r1 * x + ....
	bindingX := new(big.Int).Mul(r2, x2)
	bindingX.Mod(bindingX, n)
	bindingX.Add(bindingX, new(big.Int).Mul(r1, x))
	bindingX.Mod(bindingX, n)
	for j := 1; j <= aggsize; j++ {
		j1 := new(big.Int).SetUint64(uint64(1 + j))
		zj := new(big.Int).Exp(z, j1, n)
		zj.Mul(zj, random[j-1])
		bindingX.Add(bindingX, zj)
	}
	bindingX.Mod(bindingX, n)

	// alpha and rho blind A, S.
	u := new(big.Int).Mul(x, rho)
	u.Mod(u, n)
	u.Add(u, alpha)
	u.Mod(u, n)

	// compute new h generator vector; h' = h * (y ^ -mn).
	hPrime := hv.Hadamard(ymn.ModInverse().GetVector())

	// compute p'. p' = p - h*u. == g*l + h'*r.(this could be apply on inner product).
	newP := gv.Commit(l.GetVector())
	tmpP := hPrime.Commit(r.GetVector())
	newP.Add(newP, tmpP)

	ipparams := NewIPParams(gv, hPrime, params.U())
	ipProof, err := GenIPProof(ipparams, newP, t, l, r)
	if err != nil {
		return nil, err
	}

	proof := AggRangeProof{}
	proof.A = commitA
	proof.S = commitB
	proof.t = t
	proof.T1 = T1
	proof.T2 = T2
	proof.tx = bindingX
	proof.u = u
	proof.ipProof = ipProof

	return &proof, nil
}

// VerifyAggRangeProof verifies aggregate range proofs.
func VerifyAggRangeProof(params AggRangeParams, v []*utils.ECPoint, proof *AggRangeProof) bool {
	n := params.Curve().Params().N
	m := params.Aggsize()
	bitsize := params.Bitsize()
	size := m * bitsize
	y, err := utils.ComputeChallenge(n, proof.A.X, proof.A.Y, proof.S.X, proof.S.Y)
	if err != nil {
		log.Warn("compute challenge y failed", "error", err)
		return false
	}
	ymn := utils.PowVector(y, n, size)

	z, err := utils.ComputeChallenge(n, proof.S.X, proof.S.Y, proof.A.X, proof.A.Y)
	if err != nil {
		log.Warn("compute challenge z failed", "error", err)
		return false
	}
	zNeg := new(big.Int).Neg(z)
	zNeg.Mod(zNeg, n)
	zSquare := new(big.Int).Exp(z, utils.Two, n)

	x, err := utils.ComputeChallenge(n, proof.T1.X, proof.T1.Y, proof.T2.X, proof.T2.Y)
	if err != nil {
		log.Warn("compute challenge x failed", "error", err)
		return false
	}
	x2 := new(big.Int).Exp(x, utils.Two, n)

	h := params.H()
	g := params.G()
	gv := params.GV()
	hv := params.HV()

	// check g*tx + h*t ?= v*(z^2 * z^m) + g*delta + T1*x + T2*x^2. (z^m is a vector)
	zm := utils.PowVector(z, n, m).Times(zSquare)
	expect := utils.NewGeneratorVector(v).Commit(zm.GetVector())

	expect.Add(expect, new(utils.ECPoint).ScalarMult(proof.T1, x))
	expect.Add(expect, new(utils.ECPoint).ScalarMult(proof.T2, x2))
	delta := utils.DeltaMN(y, z, n, m, bitsize)
	expect.Add(expect, new(utils.ECPoint).ScalarMult(g, delta))

	actual := new(utils.ECPoint).ScalarMult(g, proof.t)
	actual.Add(actual, new(utils.ECPoint).ScalarMult(h, proof.tx))

	if !expect.Equal(actual) {
		log.Warn("point not equal", "expect x", expect.X, "expect y", expect.Y, "actual x", actual.X, "actual y", actual.Y)
		return false
	}

	hPrime := hv.Hadamard(ymn.ModInverse().GetVector())
	// compute p point. p = A + S*x + gv*-z + h'*(z*y^mn) + hj'^(z^j+1 * 2^n). (hj=h'[(j-1)*n:j*n-1], j=[1, m])
	p := proof.A.Copy()
	p.Add(p, new(utils.ECPoint).ScalarMult(proof.S, x))
	p.Add(p, new(utils.ECPoint).ScalarMult(gv.Sum(), zNeg))
	p.Add(p, hPrime.Commit(ymn.Times(z).GetVector()))

	n2 := utils.PowVector(new(big.Int).SetUint64(2), n, bitsize)
	for j := 1; j <= m; j++ {
		htmp := hPrime.SubVector((j-1)*bitsize, j*bitsize)
		zj := new(big.Int).Exp(z, new(big.Int).SetUint64(uint64(j+1)), n)
		zjn2 := n2.Times(zj)
		p.Add(p, htmp.Commit(zjn2.GetVector()))
	}

	// compute p'. p' = p - h*u. == g*l + h'*r.(this could be applied on inner product).
	newP := p.Sub(p, new(utils.ECPoint).ScalarMult(h, proof.u))

	return OptimizedVerifyIPProof(NewIPParams(gv, hPrime, params.U()), newP, proof.t, proof.ipProof)
}

// OptimizedVerifyAggRangeProof verifies aggregate range proof.
func OptimizedVerifyAggRangeProof(params AggRangeParams, v []*utils.ECPoint, proof *AggRangeProof) bool {
	n := params.Curve().Params().N
	m := params.Aggsize()
	bitsize := params.Bitsize()
	size := m * bitsize
	y, err := utils.ComputeChallenge(n, proof.A.X, proof.A.Y, proof.S.X, proof.S.Y)
	if err != nil {
		log.Warn("compute challenge y failed", "error", err)
		return false
	}
	ymnInverse := utils.PowVector(new(big.Int).ModInverse(y, n), n, size)

	z, err := utils.ComputeChallenge(n, proof.S.X, proof.S.Y, proof.A.X, proof.A.Y)
	if err != nil {
		log.Warn("compute challenge z failed", "error", err)
		return false
	}
	zNeg := new(big.Int).Neg(z)
	zNeg.Mod(zNeg, n)
	zSquare := new(big.Int).Exp(z, utils.Two, n)

	x, err := utils.ComputeChallenge(n, proof.T1.X, proof.T1.Y, proof.T2.X, proof.T2.Y)
	if err != nil {
		log.Warn("compute challenge x failed", "error", err)
		return false
	}
	x2 := new(big.Int).Exp(x, utils.Two, n)

	h := params.H()
	g := params.G()
	gv := params.GV()
	hv := params.HV()

	// check g*tx + h*t ?= v*(z^2 * z^m) + g*delta + T1*x + T2*x^2. (z^m is a vector)
	zm := utils.PowVector(z, n, m).Times(zSquare)
	expect := utils.NewGeneratorVector(v).Commit(zm.GetVector())

	expect.Add(expect, new(utils.ECPoint).ScalarMult(proof.T1, x))
	expect.Add(expect, new(utils.ECPoint).ScalarMult(proof.T2, x2))
	delta := utils.DeltaMN(y, z, n, m, bitsize)
	expect.Add(expect, new(utils.ECPoint).ScalarMult(g, delta))

	actual := new(utils.ECPoint).ScalarMult(g, proof.t)
	actual.Add(actual, new(utils.ECPoint).ScalarMult(h, proof.tx))
	if !expect.Equal(actual) {
		log.Warn("point not equal", "expect x", expect.X, "expect y", expect.Y, "actual x", actual.X, "actual y", actual.Y)
		return false
	}

	right := new(utils.ECPoint).ScalarMult(proof.S, x)
	right.Add(right, proof.A)

	xj := make([]*big.Int, 0)
	xj2 := make([]*big.Int, 0)
	xj2Inv := make([]*big.Int, 0)

	for i := 0; i < len(proof.ipProof.l); i++ {
		l := proof.ipProof.l[i]
		r := proof.ipProof.r[i]
		tmpx, err := utils.ComputeChallenge(n, l.X, l.Y, r.X, r.Y)
		if err != nil {
			log.Warn("compute challenge for l, r failed", "err", err)
			return false
		}

		xj = append(xj, tmpx)
		tmpx2 := new(big.Int).Mul(tmpx, tmpx)
		tmpx2.Mod(tmpx2, n)
		xj2 = append(xj2, tmpx2)

		tmpx2Inv := new(big.Int).ModInverse(tmpx2, n)
		xj2Inv = append(xj2Inv, tmpx2Inv)

		tmpp := new(utils.ECPoint).ScalarMult(l, tmpx2)
		right.Add(right, tmpp)
		tmpp = new(utils.ECPoint).ScalarMult(r, tmpx2Inv)
		right.Add(right, tmpp)
	}

	// scalar mul, add.
	tl := make([]*big.Int, size)
	tr := make([]*big.Int, size)
	rl := make([]*big.Int, size)
	ll := make([]*big.Int, size)

	challengeLen := len(proof.ipProof.l)
	n2 := utils.PowVector(new(big.Int).SetUint64(2), n, bitsize)
	for i := 0; i < size; i++ {
		if i == 0 {
			for j := 0; j < len(proof.ipProof.l); j++ {
				if j == 0 {
					tl[i] = new(big.Int).Set(xj[j])
				} else {
					tl[i].Mul(tl[i], xj[j])
					tl[i].Mod(tl[i], n)
				}

			}

			tr[i] = new(big.Int).Set(tl[i])
			tl[i] = tl[i].ModInverse(tl[i], n)
		} else {
			k := utils.GetBiggestPos(i, challengeLen)
			tl[i] = new(big.Int).Mul(tl[i-utils.Pow(k-1)], xj2[challengeLen-k])
			tl[i].Mod(tl[i], n)

			tr[i] = new(big.Int).Mul(tr[i-utils.Pow(k-1)], xj2Inv[challengeLen-k])
			tr[i].Mod(tr[i], n)
		}

		ll[i] = new(big.Int).Set(tl[i])
		rl[i] = new(big.Int).Set(tr[i])

		ll[i] = ll[i].Mul(ll[i], proof.ipProof.a)
		ll[i].Mod(ll[i], n)
		ll[i].Add(ll[i], z)
		ll[i].Mod(ll[i], n)

		rl[i] = rl[i].Mul(rl[i], proof.ipProof.b)
		rl[i].Mod(rl[i], n)

		zj := new(big.Int).Exp(z, new(big.Int).SetUint64(uint64(i/bitsize+2)), n)

		index := i % bitsize
		zjn2 := new(big.Int).Mul(zj, n2.Get(index))
		zjn2.Mod(zjn2, n)
		rl[i].Sub(rl[i], zjn2)
		rl[i].Mod(rl[i], n)
		rl[i].Mul(rl[i], ymnInverse.Get(i))
		rl[i].Mod(rl[i], n)
		rl[i].Sub(rl[i], z)
		rl[i].Mod(rl[i], n)
	}

	xu, err := utils.ComputeChallenge(n, proof.t)
	if err != nil {
		log.Warn("compute challenge for xu failed", "err", err)
		return false
	}

	left := gv.Commit(ll)
	left.Add(left, hv.Commit(rl))
	uBase := params.U()

	xabt := new(big.Int).Mul(proof.ipProof.a, proof.ipProof.b)
	xabt.Mod(xabt, n)
	xabt.Sub(xabt, proof.t)
	xabt.Mod(xabt, n)
	xabt.Mul(xabt, xu)
	xabt.Mod(xabt, n)

	left.Add(left, new(utils.ECPoint).ScalarMult(uBase, xabt))
	left.Add(left, new(utils.ECPoint).ScalarMult(h, proof.u))

	return left.Equal(right)
}
