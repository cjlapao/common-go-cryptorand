package cryptorand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomInt(t *testing.T) {
	// Arrange
	random := Rand()

	// Act
	result := random.Int()
	nextResult := random.Int()

	assert.NotEqualf(t, nextResult, result, "The numbers should not be equal in the next run")
}

func TestNewCryptoRand(t *testing.T) {
	// Arrange
	r := New()

	//act
	result := r.Rand.Int()
	nextResult := r.Rand.Int()

	assert.NotEqualf(t, nextResult, result, "The numbers should not be equal in the next run")
}
