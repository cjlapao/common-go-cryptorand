package cryptorand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomNumber(t *testing.T) {
	// Arrange + Act
	result := GetRandomNumber(5, 100)
	nextResult := GetRandomNumber(5, 100)

	// Assert
	assert.NotEqualf(t, result, nextResult, "The numbers should not be equal")
	assert.GreaterOrEqualf(t, nextResult, 5, "The number should be greater or equal to 5")
	assert.LessOrEqualf(t, nextResult, 100, "The number should be less or equal to 100")
}
