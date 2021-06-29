package curve

import (
	"crypto/elliptic"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/crypto"
)

// S256 returns curve used in btc.
func S256() elliptic.Curve {
	return crypto.S256()
}

// NoCGOS256 returns curve.
func NoCGOS256() elliptic.Curve {
	return btcec.S256()
}

// BN256 returns curve alt bn128.
func BN256() elliptic.Curve {
	return &BN128{}
}
