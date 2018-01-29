package flow

import (
	"fmt"
	"runtime"
)

func whatOs() string {
	return buildWhatOs(runtime.GOOS)
}

func buildWhatOs(name string) string {
	what := "Go runs on "

	switch os := name; os {
	case "darwin":
		what += "OS X."
	case "linux":
		what += "Linux."
	default:
		what += os + "."
	}

	return what
}

type fn func(int) int

func testSwitch(cond fn, num int) {

	testFun := func() int { return num }

	switch testFun() {
	case 1:
		fmt.Println("Case 1 executed.")
	case cond(2):
		fmt.Println("Case 2 executed.")
	case cond(3):
		fmt.Println("Case 3 executed.")
	}
}

func switchNoCond(x int) string {

	// switch doesn't need conditions, can be used like if/else statement
	switch {
	case x > 10:
		return "More than 10."
	case x < 10:
		return "Less than 10."
	}
	return "10"
}
