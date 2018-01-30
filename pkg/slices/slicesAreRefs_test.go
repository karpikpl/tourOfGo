package slices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_slices_Should_beReferences(t *testing.T) {

	// Arrange
	names := [4]string{"John", "Bob", "Alice", "Barrack"}

	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// Act
	b[0] = "Foo Bar"

	// assert
	fmt.Println(names)
	fmt.Println(a, b)
	assert.Equal(t, b[0], "Foo Bar")
	assert.Equal(t, a[1], "Foo Bar")
	assert.Equal(t, names[1], "Foo Bar")
}
