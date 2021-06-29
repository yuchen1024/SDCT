package sdctsys

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"math/big"

	"github.com/sdct/proof"
	"github.com/sdct/utils"
)

// Account represents the account used in sdct system.
type Account struct {
	name    string
	sk      *ecdsa.PrivateKey
	balance *proof.CTEncPoint

	m      *big.Int
	nonce  uint64
	params proof.CTParams
}

// CreateTestAccount creates account for test purpose.
func CreateTestAccount(params proof.CTParams, name string, balance *big.Int) *Account {
	a := Account{}
	a.name = name
	a.m = new(big.Int).Set(balance)
	a.nonce = 0
	a.params = params
	key := proof.MustGenerateKey(params)

	a.sk = key
	ct, err := proof.Encrypt(params, &a.sk.PublicKey, balance.Bytes())
	if err != nil {
		panic(err)
	}
	a.balance = ct.CopyPublicPoint()

	return &a
}

// UpdateBalance updates user's encrypted balance and nonce.
func (a *Account) UpdateBalance(nonce *big.Int, ct [4]*big.Int) {
	a.balance.X = utils.NewECPoint(ct[0], ct[1], a.sk.Curve)
	a.balance.Y = utils.NewECPoint(ct[2], ct[3], a.sk.Curve)
	a.nonce = nonce.Uint64()

	// decrypt balance
	balance := proof.Decrypt(a.params, a.sk, a.balance)

	a.m = new(big.Int).SetBytes(balance)
}

// Info prints account infomation.
func (a *Account) Info() {
	fmt.Println("-----------------------------------------------------------------------------")
	fmt.Println("New SDCT account created")
	fmt.Println("Account name:", a.name)
	fmt.Printf("Account pubkey: %x\n", elliptic.Marshal(a.params.Curve(), a.sk.X, a.sk.Y))
}

func (a *Account) Priv() *ecdsa.PrivateKey {
	return a.sk
}

func (a *Account) M() *big.Int {
	return a.m
}

func (a *Account) Name() string {
	return a.name
}
