package flow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ifStatement_Should_ReturnRandomBiggerThan10(t *testing.T) {

	// Arrange

	// Act
	result, iterations := ifStatement(1)
	fmt.Printf("Received %v after %v iterations\n", result, iterations)

	//Assert
	assert.Condition(t, func() bool { return result > 0.5 }, "random value greater than 0.5 - %v", result)
}
