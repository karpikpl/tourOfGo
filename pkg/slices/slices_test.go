package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstTwoSlice_Should_return_secondAndThirdElements(t *testing.T) {
	// Arrange
	array := []int{1, 2, 3, 4, 5, 6}

	// Act
	result, err := firstTwoSlice(array)

	// Assert
	assert.Nil(t, err)
	assert.ElementsMatch(t, []int{2, 3}, result)
}

func Test_firstTwoSlice_Should_return_secondAndThirdElements_when_ArrayLen3(t *testing.T) {
	// Arrange
	array := []int{4, 5, 6}

	// Act
	result, err := firstTwoSlice(array)

	// Assert
	assert.Nil(t, err)
	assert.ElementsMatch(t, []int{6, 5}, result)
}

func Test_firstTwoSlice_Should_return_error_when_ArrayLenLessThan3(t *testing.T) {
	// Arrange
	array := []int{4, 5}

	// Act
	result, err := firstTwoSlice(array)

	// Assert
	assert.Nil(t, result)
	assert.NotNil(t, err)
}
