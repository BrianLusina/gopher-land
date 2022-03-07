// Diffie-Hellman-Merkle key exchange
package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))

func PrivateKey(p *big.Int) *big.Int {
	t := big.NewInt(2)
	pK := new(big.Int).Sub(p, t)
	return pK.Add(t, pK.Rand(r, pK))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	a := new(big.Int).Exp(big.NewInt(g), private, p)
	return a
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	public := PublicKey(private, p, g)
	return private, public
}

func SecretKey(private, public, p *big.Int) *big.Int {
	return new(big.Int).Exp(public, private, p)
}
