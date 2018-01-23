package variables

import "fmt"

// The var statement declares a list of variables
var c, python, java bool

// decaration can include initializers (type can be skipped)
var i, j, text = 1, 2, "this is pretty cool but readability is questionable"

func varStatment() {
	var i int
	fmt.Println(i, c, python, java)
	fmt.Println(i, j, text)
}

func declarations() {
	// short assignment statement
	golang, node, java := true, "YES!", "NOOOOOO"
	fmt.Println(golang, node, java)
}
