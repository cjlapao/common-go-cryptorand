package rand_test

import (
	"testing"

	"github.com/cjlapao/common-go-cryptorand/rand"
	"github.com/stretchr/testify/assert"
)

func TestGenerateAlphaNumericRandomString(t *testing.T) {
	// Arrange + Act
	result1, err1 := rand.GetAlphaNumericRandomString(5)
	result2, err2 := rand.GetAlphaNumericRandomString(5)

	// Assert
	assert.NoErrorf(t, err1, "Error should be nil")
	assert.NoErrorf(t, err2, "Error should be nil")
	assert.NotEqualf(t, result1, result2, "Strings should not match")
	assert.Lenf(t, result1, 5, "Initial String size should be 5")
	assert.Lenf(t, result1, 5, "Secondary String size should be 5")
}

func TestGenerateRandomString(t *testing.T) {
	// Arrange + Act
	result1, err1 := rand.GetRandomString(15)
	result2, err2 := rand.GetRandomString(15)

	// Assert
	assert.NoErrorf(t, err1, "Error should be nil")
	assert.NoErrorf(t, err2, "Error should be nil")
	assert.NotEqualf(t, result1, result2, "Strings should not match")
	assert.Lenf(t, result1, 15, "Initial String size should be 15")
	assert.Lenf(t, result1, 15, "Secondary String size should be 15")
}

func TestGenerateUpperCaseRandomString(t *testing.T) {
	// Arrange + Act
	result1, err1 := rand.GetUpperCaseRandomString(7)
	result2, err2 := rand.GetUpperCaseRandomString(7)

	// Assert
	assert.NoErrorf(t, err1, "Error should be nil")
	assert.NoErrorf(t, err2, "Error should be nil")
	assert.NotEqualf(t, result1, result2, "Strings should not match")
	assert.Lenf(t, result1, 7, "Initial String size should be 7")
	assert.Lenf(t, result1, 7, "Secondary String size should be 7")
}

func TestGenerateLowerCaseRandomString(t *testing.T) {
	// Arrange + Act
	result1, err1 := rand.GetLowerCaseRandomString(7)
	result2, err2 := rand.GetLowerCaseRandomString(7)

	// Assert
	assert.NoErrorf(t, err1, "Error should be nil")
	assert.NoErrorf(t, err2, "Error should be nil")
	assert.NotEqualf(t, result1, result2, "Strings should not match")
	assert.Lenf(t, result1, 7, "Initial String size should be 7")
	assert.Lenf(t, result1, 7, "Secondary String size should be 7")
}

func TestGenerateNumericRandomString(t *testing.T) {
	// Arrange + Act
	result1, err1 := rand.GetNumericRandomString(7)
	result2, err2 := rand.GetNumericRandomString(7)

	// Assert
	assert.NoErrorf(t, err1, "Error should be nil")
	assert.NoErrorf(t, err2, "Error should be nil")
	assert.NotEqualf(t, result1, result2, "Strings should not match")
	assert.Lenf(t, result1, 7, "Initial String size should be 7")
	assert.Lenf(t, result1, 7, "Secondary String size should be 7")
}
