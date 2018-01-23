package functions

import "testing"
import "github.com/stretchr/testify/assert"

func Test_Add_Should_ReturnASumOfTwoNumbers(t *testing.T) {

	// Arrange

	// Act
	result := Add(5, 4)

	// Assert
	assert.Equal(t, 9, result)
}

func Test_Sub_Should_ReturnADifferenceOfTwoNumbers(t *testing.T) {

	// Arrange

	// Act
	result := Sub(10, 12)

	// Assert
	assert.Equal(t, -2, result)
}

func Test_Swap_Should_SwapVariablesProvided(t *testing.T) {

	// Arrange

	// Act
	a, b := Swap("a", "b")

	// Assert
	assert.Equal(t, "b", a)
	assert.Equal(t, "a", b)
}

func Test_Split_Should_SplitSumIntoTwoParts(t *testing.T) {

	// Arrange

	// Act
	a, b := Split(17)

	// Assert
	assert.Equal(t, 7, a)
	assert.Equal(t, 10, b)
}
