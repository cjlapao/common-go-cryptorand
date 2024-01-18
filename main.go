package cryptorand

import (
	"crypto/rand"
	"math/big"
)

type CryptoRand struct {
}

func New() *CryptoRand {
	return &CryptoRand{}
}

func (c *CryptoRand) Int() (int, error) {
	val, err := rand.Int(rand.Reader, big.NewInt(int64(len(NumericCharacters()))))
	if err != nil {
		return 0, err
	}
	return int(val.Int64()), nil
}

func (c *CryptoRand) GetRandomNumber(min, max int) (int, error) {
	return GetRandomNumber(min, max)
}

func (c *CryptoRand) GetAlphaNumericRandomString(size int) (string, error) {
	source := AlphaNumericCharacters()

	return c.getRandomFromArray(size, source)
}

func (c *CryptoRand) GetRandomString(size int) (string, error) {
	source := AlphaCharacters()

	return c.getRandomFromArray(size, source)
}

func (c *CryptoRand) GetUpperCaseRandomString(size int) (string, error) {
	source := make([]string, 0)
	source = append(source, UpperCaseAlphaCharacters()...)

	return c.getRandomFromArray(size, source)
}

func (c *CryptoRand) GetLowerCaseRandomString(size int) (string, error) {
	source := make([]string, 0)
	source = append(source, LowerCaseAlphaCharacters()...)

	return c.getRandomFromArray(size, source)
}

func (c *CryptoRand) getRandomFromArray(size int, source []string) (string, error) {
	result := ""
	if len(source) > 0 {
		random, err := rand.Int(rand.Reader, big.NewInt(int64(len(source))))
		if err != nil {
			return "", err
		}
		if size > 0 {
			for i := 0; i < size; i++ {
				idx := random.Int64()
				result += source[idx]
			}
		}
	}

	return result, nil
}
