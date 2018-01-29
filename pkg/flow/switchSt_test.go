package flow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_whatOs_should_returnAString(t *testing.T) {

	// Arrange

	// Act
	result := whatOs()

	// Assert
	fmt.Println(result)
	assert.NotEmpty(t, result)
}

func Test_buildWhatOs_should_returnOSX_when_darwin(t *testing.T) {

	// Arrange

	// Act
	result := buildWhatOs("darwin")

	// Assert
	assert.Equal(t, result, "Go runs on OS X.")
}

func Test_buildWhatOs_should_returnLinux_when_linux(t *testing.T) {

	// Arrange

	// Act
	result := buildWhatOs("linux")

	// Assert
	assert.Equal(t, result, "Go runs on Linux.")
}

func Test_buildWhatOs_should_returnFoo_when_foo(t *testing.T) {

	// Arrange

	// Act
	result := buildWhatOs("Foo")

	// Assert
	assert.Equal(t, result, "Go runs on Foo.")
}

func Test_testSwith_Should_Not_Call2And3_when_1(t *testing.T) {

	// Arrange
	called := 0
	cond := func(x int) int {
		fmt.Printf("%x called\n", x)
		called++
		return x
	}

	// Act
	testSwitch(cond, 1)

	// Assert
	assert.Equal(t, called, 0, "Ony first condition is executed, so cond should never be called")
}

func Test_testSwith_Should_CallCondOnce_when_2(t *testing.T) {

	// Arrange
	called := 0
	cond := func(x int) int {
		fmt.Printf("Cond %x called\n", x)
		called++
		return x
	}

	// Act
	testSwitch(cond, 2)

	// Assert
	assert.Equal(t, called, 1, "Two conditions are executed, so cond will be colled 1 time")
}

func Test_testSwith_Should_CallCondTwice_when_3(t *testing.T) {

	// Arrange
	called := 0
	cond := func(x int) int {
		fmt.Printf("Cond %x called\n", x)
		called++
		return x
	}

	// Act
	testSwitch(cond, 3)

	// Assert
	assert.Equal(t, called, 2, "All conditions are executed, so cond should be called 2 times")
}

func Test_switchNoCond_Should_ReturnLessThen10_when_2(t *testing.T) {

	// Arrange

	// Act
	result := switchNoCond(2)

	// Assert
	assert.Equal(t, result, "Less than 10.")
}

func Test_switchNoCond_Should_ReturnMoreThen10_when_1000(t *testing.T) {

	// Arrange

	// Act
	result := switchNoCond(1000)

	// Assert
	assert.Equal(t, result, "More than 10.")
}

func Test_switchNoCond_Should_Return10_when_10(t *testing.T) {

	// Arrange

	// Act
	result := switchNoCond(10)

	// Assert
	assert.Equal(t, result, "10")
}
