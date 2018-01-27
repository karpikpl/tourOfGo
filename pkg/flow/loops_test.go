package flow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sum_Should_Return45(t *testing.T) {

	// Arrange

	// Act
	result := sum()

	// Assert
	assert.Equal(t, 45, result)
}

func Test_sumOptional_Should_Return45(t *testing.T) {

	// Arrange

	// Act
	result := sumOptional()

	// Assert
	assert.Equal(t, 1024, result)
}

func Test_infiniteLoop_Should_Return46(t *testing.T) {

	// Arrange

	// Act
	result := infiniteLoop()

	// Assert
	assert.Equal(t, 46, result)
}
