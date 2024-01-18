package rand_test

import (
	"testing"

	"github.com/cjlapao/common-go-cryptorand/rand"
	"github.com/stretchr/testify/assert"
)

func TestGetRandomNumber(t *testing.T) {
	// Arrange + Act
	result, err1 := rand.GetRandomNumber(5, 100)
	nextResult, err2 := rand.GetRandomNumber(5, 100)

	// Assert
	assert.NoErrorf(t, err1, "Error should be nil")
	assert.NoErrorf(t, err2, "Error should be nil")
	assert.NotEqualf(t, result, nextResult, "The numbers should not be equal")
	assert.GreaterOrEqualf(t, nextResult, 5, "The number should be greater or equal to 5")
	assert.LessOrEqualf(t, nextResult, 100, "The number should be less or equal to 100")
}
