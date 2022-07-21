package cryptorand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAlphaNumericRandomString(t *testing.T) {
	// Arrange + Act
	result1 := GetAlphaNumericRandomString(5)
	result2 := GetAlphaNumericRandomString(5)

	// Assert
	assert.NotEqualf(t, result1, result2, "Strings should not match")
	assert.Lenf(t, result1, 5, "Initial String size should be 5")
	assert.Lenf(t, result1, 5, "Secondary String size should be 5")
}

func TestGenerateRandomString(t *testing.T) {
	// Arrange + Act
	result1 := GetRandomString(15)
	result2 := GetRandomString(15)

	// Assert
	assert.NotEqualf(t, result1, result2, "Strings should not match")
	assert.Lenf(t, result1, 15, "Initial String size should be 15")
	assert.Lenf(t, result1, 15, "Secondary String size should be 15")
}

func TestGenerateUpperCaseRandomString(t *testing.T) {
	// Arrange + Act
	result1 := GetUpperCaseRandomString(7)
	result2 := GetUpperCaseRandomString(7)

	// Assert
	assert.NotEqualf(t, result1, result2, "Strings should not match")
	assert.Lenf(t, result1, 7, "Initial String size should be 7")
	assert.Lenf(t, result1, 7, "Secondary String size should be 7")
}

func TestGenerateLowerCaseRandomString(t *testing.T) {
	// Arrange + Act
	result1 := GetLowerCaseRandomString(7)
	result2 := GetLowerCaseRandomString(7)

	// Assert
	assert.NotEqualf(t, result1, result2, "Strings should not match")
	assert.Lenf(t, result1, 7, "Initial String size should be 7")
	assert.Lenf(t, result1, 7, "Secondary String size should be 7")
}

func TestGenerateNumericRandomString(t *testing.T) {
	// Arrange + Act
	result1 := GetNumericRandomString(7)
	result2 := GetNumericRandomString(7)

	// Assert
	assert.NotEqualf(t, result1, result2, "Strings should not match")
	assert.Lenf(t, result1, 7, "Initial String size should be 7")
	assert.Lenf(t, result1, 7, "Secondary String size should be 7")
}
