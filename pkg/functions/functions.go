package functions

// functions that start with capital case are exported

func Add(x int, y int) int {
	return x + y
}

// shorter method definition when same type
func Sub(x, y int) int {
	return x - y
}

// multiple results
func Swap(x, y string) (string, string) {
	return y, x
}

// naked return
// Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
