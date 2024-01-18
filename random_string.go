package cryptorand

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func GetAlphaNumericRandomString(size int) (string, error) {
	return randomString(size, true, true, true)
}

func GetRandomString(size int) (string, error) {
	return randomString(size, true, true, false)
}

func GetUpperCaseRandomString(size int) (string, error) {
	return randomString(size, false, true, false)
}

func GetLowerCaseRandomString(size int) (string, error) {
	return randomString(size, true, false, false)
}

func GetNumericRandomString(size int) (string, error) {
	return randomString(size, false, false, true)
}

func randomString(size int, includeLowerCase bool, includeUpperCase bool, includeNumeric bool) (string, error) {
	source := make([]string, 0)
	result := make(([]byte), size)

	if includeLowerCase {
		source = append(source, LowerCaseAlphaCharacters()...)
	}

	if includeUpperCase {
		source = append(source, UpperCaseAlphaCharacters()...)
	}

	if includeNumeric {
		source = append(source, NumericCharacters()...)
	}
	if len(source) > 0 {
		if size > 0 {
			for i := 0; i < size; i++ {
				idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(source))))
				if err != nil {
					return "", err
				}
				result[i] = source[idx.Int64()][0]
			}
		}
	}

	return strings.Join([]string{string(result)}, ""), nil
}
