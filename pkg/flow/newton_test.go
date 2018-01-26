package flow

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sqrt_Should_returnResultCloseToMathSqrt(t *testing.T) {

	// Arrange

	// Act
	result := sqrt(2)

	// Assert
	assert.InDelta(t, math.Sqrt(2), result, 0.01)
}
