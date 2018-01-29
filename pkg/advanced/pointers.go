package advanced

import "fmt"

func pointers(x int) int {
	i, j := x, 2701

	p := &i                                               // point to i
	fmt.Println("Value of i through the pointer p: ", *p) // read i through the pointer

	*p = x / 2 // set i through the pointer
	fmt.Println("New value of i: ", i)

	p = &j       // now p points to j
	*p = *p / 37 // divide j through the pointer and set new value for j

	fmt.Println("new value of j: ", j)

	return i
}
