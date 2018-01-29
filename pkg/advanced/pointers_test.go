package advanced

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pointers_Should_Return10_When_20(t *testing.T) {

	// Arrange
	val := 20

	// Act
	result := pointers(val)

	// Assert
	assert.Equal(t, 10, result)
}
