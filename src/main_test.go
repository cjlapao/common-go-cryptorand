package cryptorand

import (
	"testing"

	"github.com/cjlapao/common-go-cryptorand/constants"
	"github.com/stretchr/testify/assert"
)

func TestRandomInt(t *testing.T) {
	// Arrange
	random := New()

	// Act
	result, err := random.Int()
	nextResult, nextErr := random.Int()

	// Assert
	assert.NoError(t, err, "Error should be nil")
	assert.NoError(t, nextErr, "Error should be nil")
	assert.NotEqualf(t, nextResult, result, "The numbers should not be equal in the next run")
}

func TestCryptoRand_GetAlphaNumericRandomString(t *testing.T) {
	// Arrange
	c := New()
	size := 10

	// Act
	result, err := c.GetAlphaNumericRandomString(size)

	// Assert
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, size, len(result), "Result length should be equal to size")
	for _, char := range result {
		source := make([]string, 0)
		source = append(source, constants.LowerCaseAlphaCharacters()...)
		source = append(source, constants.UpperCaseAlphaCharacters()...)
		source = append(source, constants.NumericCharacters()...)
		assert.Contains(t, source, string(char), "Result should contain only alphanumeric characters")
	}
}
func TestCryptoRand_GetRandomNumber(t *testing.T) {
	// Arrange
	c := New()
	min := 0
	max := 100

	// Act
	result, err := c.GetRandomNumber(min, max)

	// Assert
	assert.NoError(t, err, "Error should be nil")
	assert.GreaterOrEqual(t, result, min, "Result should be greater than or equal to min")
	assert.LessOrEqual(t, result, max, "Result should be less than or equal to max")
}

func TestCryptoRand_GetRandomString(t *testing.T) {
	// Arrange
	c := New()
	size := 10

	// Act
	result, err := c.GetRandomString(size)

	// Assert
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, size, len(result), "Result length should be equal to size")
	for _, char := range result {
		source := make([]string, 0)
		source = append(source, constants.LowerCaseAlphaCharacters()...)
		source = append(source, constants.UpperCaseAlphaCharacters()...)
		assert.Contains(t, source, string(char), "Result should contain only alphanumeric characters")
	}
}
func TestCryptoRand_GetUpperCaseRandomString(t *testing.T) {
	// Arrange
	c := New()
	size := 10

	// Act
	result, err := c.GetUpperCaseRandomString(size)

	// Assert
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, size, len(result), "Result length should be equal to size")
	for _, char := range result {
		assert.Contains(t, constants.UpperCaseAlphaCharacters(), string(char), "Result should contain only uppercase alphabetic characters")
	}
}
func TestCryptoRand_GetLowerCaseRandomString(t *testing.T) {
	// Arrange
	c := New()
	size := 10

	// Act
	result, err := c.GetLowerCaseRandomString(size)

	// Assert
	assert.NoError(t, err, "Error should be nil")
	assert.Equal(t, size, len(result), "Result length should be equal to size")
	for _, char := range result {
		assert.Contains(t, constants.LowerCaseAlphaCharacters(), string(char), "Result should contain only lowercase alphabetic characters")
	}
}
