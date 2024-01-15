package cryptorand

import (
	"crypto/rand"
	"math/big"
)

func GetRandomNumber(min, max int) (int, error) {
	bg := big.NewInt(int64(max) - int64(min))

	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		return 0, err
	}

	return int(n.Int64() + int64(min)), nil
}
