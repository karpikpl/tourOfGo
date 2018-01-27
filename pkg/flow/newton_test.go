package flow

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sqrt_Should_returnResultCloseToMathSqrt(t *testing.T) {

	// Arrange

	// Act
	result := sqrt(2)

	// Assert
	assert.InDelta(t, math.Sqrt(2), result, 0.8)
}

func Test_sqrt2_Should_returnResultCloseToMathSqrt(t *testing.T) {

	// Arrange

	// Act
	result, iterations := sqrt2(2)

	// Assert
	fmt.Printf("Found %v after %v iterations\n", result, iterations)
	assert.InDelta(t, math.Sqrt(2), result, 0.01)
}

func Test_sqrt3_Should_returnResultCloseToMathSqrt(t *testing.T) {

	// Arrange

	// Act
	result, iterations := sqrt3(2)

	// Assert
	fmt.Printf("Found %v after %v iterations\n", result, iterations)
	assert.InDelta(t, math.Sqrt(2), result, 0.01)
}
