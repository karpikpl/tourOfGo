package advanced

import "fmt"

type fn func(int)

// A deferred function's arguments are evaluated when the defer statement is evaluated.
func deferEval(fun fn) int {
	i := 0
	defer fun(i)
	i++
	return i
}

// Deferred function calls are executed in Last In First Out order after the surrounding function returns.
func deferOrder(fun fn) {
	for i := 0; i < 4; i++ {
		defer fun(i)
	}
}

// Deferred functions may read and assign to the returning function's named return values.
func deferReturn() (i int) {
	defer func() {
		fmt.Printf("Altering return value (%v) by adding 5\n", i)
		i += 5
	}()
	fmt.Println("Returning 1")
	return 1
}
