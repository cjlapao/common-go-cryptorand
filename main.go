package cryptorand

import "math/rand"

type CryptoRand struct {
	Rand *rand.Rand
}

func New() *CryptoRand {
	rand := Rand()
	return &CryptoRand{
		Rand: rand,
	}
}

func Rand() *rand.Rand {
	var src CryptoSource
	generator := rand.New(src)
	return generator
}
