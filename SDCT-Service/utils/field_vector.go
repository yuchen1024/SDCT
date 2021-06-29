package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"

	log "github.com/inconshreveable/log15"
)

// FieldVector respresents vector scalar array.
type FieldVector struct {
	items []*big.Int
	// order n.
	n *big.Int
}

// NewFieldVector creates instance of field vector.
func NewFieldVector(items []*big.Int, n *big.Int) *FieldVector {
	newItems := make([]*big.Int, len(items))
	for i, item := range items {
		newItems[i] = new(big.Int).Set(item)
	}

	return newFieldVector(newItems, new(big.Int).Set(n))
}

// newFieldVector returns instance of field vector.
func newFieldVector(items []*big.Int, n *big.Int) *FieldVector {
	f := FieldVector{}
	f.items = items
	f.n = n

	return &f
}

// NewRandomFieldVector generates field vector randomly.
// Warning: Just for test purpose.
func NewRandomFieldVector(order *big.Int, n int) *FieldVector {
	f := FieldVector{}
	f.n = new(big.Int).Set(order)
	f.items = make([]*big.Int, 0)
	for i := 0; i < n; i++ {
		tmp, err := rand.Int(rand.Reader, order)
		if err != nil {
			// no sense for test purpose.
			panic(err)
		}

		f.items = append(f.items, tmp)
	}

	return &f
}

// RepeatItemVector returns a field vector whose items are same.
func RepeatItemVector(item, order *big.Int, n int) *FieldVector {
	f := FieldVector{}
	f.n = new(big.Int).Set(order)
	f.items = make([]*big.Int, 0)
	newItem := new(big.Int).Set(item)
	newItem.Mod(newItem, order)
	for i := 0; i < n; i++ {
		f.items = append(f.items, new(big.Int).Set(newItem))
	}

	return &f
}

// PowVector returns a field vector(1, y^1, y^2,..., y^n-1)
func PowVector(itemBase, order *big.Int, n int) *FieldVector {
	f := FieldVector{}
	f.items = make([]*big.Int, 0)
	f.n = new(big.Int).Set(order)
	for i := 0; i < n; i++ {
		exponent := new(big.Int).SetUint64(uint64(i))
		item := new(big.Int).Exp(itemBase, exponent, order)
		f.items = append(f.items, item)
	}

	return &f
}

// HalfLeft returns half of items on left.
func (f *FieldVector) HalfLeft() *FieldVector {
	return f.SubFieldVector(0, f.Size()/2)
}

// HalfRight returns half of items on right.
func (f *FieldVector) HalfRight() *FieldVector {
	size := f.Size()
	return f.SubFieldVector(size/2, size)
}

// Size returns len of underlying items.
func (f *FieldVector) Size() int {
	return len(f.items)
}

// SubFieldVector returns sub items by start/end index.
func (f *FieldVector) SubFieldVector(start, end int) *FieldVector {
	if start < 0 || end > f.Size() {
		panic(fmt.Sprintf("field vector index start %d, end %d out of range", start, end))
	}

	newItems := make([]*big.Int, 0)
	for _, item := range f.items[start:end] {
		newItems = append(newItems, new(big.Int).Set(item))
	}

	return newFieldVector(newItems, new(big.Int).Set(f.n))
}

// First returns first item in field.
func (f *FieldVector) First() *big.Int {
	return new(big.Int).Set(f.Get(0))
}

// Get returns item by index.
func (f *FieldVector) Get(i int) *big.Int {
	return new(big.Int).Set(f.items[i])
}

// InnerProduct computes <items, other>.
func (f *FieldVector) InnerProduct(other *FieldVector) *big.Int {
	if f.Size() != other.Size() {
		panic(fmt.Sprintf("field vector size %d != %d", f.Size(), other.Size()))
	}

	res := new(big.Int)
	for i, item := range f.items {
		tmp := new(big.Int).Mul(item, other.Get(i))
		res.Add(res, tmp)
		res.Mod(res, f.n)
	}

	return res.Mod(res, f.n)
}

// GetVector returns underlying items.
func (f *FieldVector) GetVector() []*big.Int {
	newItems := make([]*big.Int, 0)
	for _, item := range f.items {
		newItems = append(newItems, new(big.Int).Set(item))
	}

	return newItems
}

// AllItemsSub returns a new field vector whose items = ori item - d mod n.
func (f *FieldVector) AllItemsSub(d *big.Int) *FieldVector {
	newItems := make([]*big.Int, 0)
	for _, item := range f.items {
		newItem := new(big.Int).Sub(item, d)
		newItem.Mod(newItem, f.n)

		newItems = append(newItems, newItem)
	}

	return newFieldVector(newItems, new(big.Int).Set(f.n))
}

// AllItemsSubOne returns a new field vector whose items = ori item - 1 mod n.
func (f *FieldVector) AllItemsSubOne() *FieldVector {
	newItems := make([]*big.Int, 0)
	one := new(big.Int).SetUint64(1)
	for _, item := range f.items {
		newItem := new(big.Int).Sub(item, one)
		newItem.Mod(newItem, f.n)

		newItems = append(newItems, newItem)
	}

	return newFieldVector(newItems, new(big.Int).Set(f.n))
}

// ModInverse returns new field vector whose item = modInverse(ori item).
func (f *FieldVector) ModInverse() *FieldVector {
	newItems := make([]*big.Int, 0)

	for _, item := range f.items {
		newItem := new(big.Int).ModInverse(item, f.n)
		newItems = append(newItems, newItem)
	}

	return newFieldVector(newItems, new(big.Int).Set(f.n))
}

// Times compute item * x and returns new instance.
func (f *FieldVector) Times(x *big.Int) *FieldVector {
	newItems := make([]*big.Int, 0)

	for _, i := range f.items {
		t := new(big.Int).Mul(i, x)
		t.Mod(t, f.n)
		newItems = append(newItems, t)
	}

	return newFieldVector(newItems, new(big.Int).Set(f.n))
}

// Copy returns a copy of current.
func (f *FieldVector) Copy() *FieldVector {
	return NewFieldVector(f.items, f.n)
}

// Append appends another filed vector to current.
func (f *FieldVector) Append(another *FieldVector) *FieldVector {
	newItems := make([]*big.Int, 0)
	for _, item := range f.items {
		newItems = append(newItems, item)
	}
	for _, item := range another.GetVector() {
		newItems = append(newItems, item)
	}

	return newFieldVector(newItems, new(big.Int).Set(f.n))
}

// AddFieldVector computes fi + other.
func (f *FieldVector) AddFieldVector(other *FieldVector) *FieldVector {
	if f.Size() != other.Size() {
		panic(fmt.Sprintf("filed vector size not equal %d != %d", f.Size(), other.Size()))
	}

	newItems := make([]*big.Int, 0)
	for i, item := range f.items {
		t := new(big.Int).Add(item, other.Get(i))
		t.Mod(t, f.n)
		newItems = append(newItems, t)
	}

	return newFieldVector(newItems, new(big.Int).Set(f.n))
}

// Hadamard returns field vector (a1*b1, a2*b2, ..., an*bn).
func (f *FieldVector) Hadamard(other *FieldVector) *FieldVector {
	if f.Size() != other.Size() {
		panic(fmt.Sprintf("filed vector size not equal %d != %d", f.Size(), other.Size()))
	}

	newItems := make([]*big.Int, 0)
	for i, item := range f.items {
		t := new(big.Int).Mul(item, other.Get(i))
		t.Mod(t, f.n)
		newItems = append(newItems, t)
	}

	return newFieldVector(newItems, new(big.Int).Set(f.n))
}

// Sum returns a1 + a2 + .... + an.
func (f *FieldVector) Sum() *big.Int {
	res := new(big.Int).SetUint64(0)

	for _, item := range f.items {
		res.Add(res, item)
		res.Mod(res, f.n)
	}

	return res
}

// Log .
func (f *FieldVector) Log() {
	for _, item := range f.items {
		log.Debug("vector", "v", item)
	}
}
