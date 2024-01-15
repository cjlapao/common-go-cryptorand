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
		random := Rand()
		if size > 0 {
			for i := 0; i < size; i++ {
				idx := random.Intn(len(source))
				result += source[idx]
			}
		}
	}

	return result, nil
}
