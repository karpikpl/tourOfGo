package slices

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sliceLiterals(t *testing.T) {

	// Arrange
	array := [3]int{5, 15, 20}
	slice := []int{5, 15, 20}
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
	}

	// Act
	arrayType := reflect.TypeOf(array).String()
	sliceType := reflect.TypeOf(slice).String()

	// Assert
	fmt.Println(s)
	assert.Equal(t, len(array), len(slice))
	assert.Equal(t, "[]int", sliceType)
	assert.Equal(t, "[3]int", arrayType)
}

func Test_sliceDefaults(t *testing.T) {

	// Arrange
	array := [5]int{67, 1, 5, 6, 9}

	// Act
	s1 := array[0:5]
	s2 := array[:5]
	s3 := array[0:]
	s4 := array[:]

	// Assert
	assert.ElementsMatch(t, array, s1)
	assert.ElementsMatch(t, array, s2)
	assert.ElementsMatch(t, array, s3)
	assert.ElementsMatch(t, array, s4)
}
