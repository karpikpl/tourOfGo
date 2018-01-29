package advanced

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_vertexCreatedWithoutParams_Should_HaveDefaultValues(t *testing.T) {

	// Arrange

	// Act
	v := Vertex{}

	// Assert
	fmt.Println(v)
	assert.Equal(t, "", v.d)
}

func Test_vertexCreatedWithParams_Should_HaveProvidedValues(t *testing.T) {

	// Arrange

	// Act
	v := Vertex{1, 100, "test"}

	// Assert
	fmt.Println(v)
	assert.Equal(t, "test", v.d)
	assert.Equal(t, 1, v.X)
	assert.Equal(t, 100, v.Y)
}

func Test_vertexPointerSyntax_Should_NotRequireAsterisk(t *testing.T) {

	// Arrange
	v := Vertex{d: "the order of named fields is irrelevant."}

	// Act
	p := &v
	p.X = 999

	// Assert
	fmt.Println(v)
	assert.Equal(t, 999, v.X)
}

func Test_createStructWithPointer(t *testing.T) {

	// Arrange
	p := &Vertex{X: 6, Y: 19}

	// Act
	v := *p

	// Assert
	fmt.Println(v)
	assert.Equal(t, v.X, p.X)
}
