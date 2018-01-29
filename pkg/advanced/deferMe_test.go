package advanced

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deferEval_should_callDefferedFunctionWith0(t *testing.T) {

	// Arrange
	var defferedParam int
	deferFn := func(x int) {
		defferedParam = x
		fmt.Println("Deffer called with ", x)
	}

	// Act
	result := deferEval(deferFn)

	// Assert
	assert.Equal(t, 1, result, "result shoul be 1")
	assert.Equal(t, 0, defferedParam, "deffered function should be called with 0")
}

func Test_deferOrder_should_callDefferedFunctionWith3210(t *testing.T) {

	// Arrange
	var defferedParams string
	deferFn := func(x int) {
		defferedParams += strconv.Itoa(x)
		fmt.Println("Deffer called with ", x)
	}

	// Act
	deferOrder(deferFn)

	// Assert
	assert.Equal(t, "3210", defferedParams, "deffered function should be called with 3210 - last in, first out")
}

func Test_deferReturn_should_return6(t *testing.T) {

	// Arrange

	// Act
	result := deferReturn()

	// Assert
	assert.Equal(t, 6, result, "deffered function should ater the return value")
}
