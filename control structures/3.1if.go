package main

import "fmt"

func main() {
	x := 5
	//simple if looks like this
	if x > 0 {
		fmt.Println("Bigger than Zero")
	}

	//Since if and switch accept an initialization statement, it's common to see one used to set up a local variable.
	if y := 10; y < 11 {
		fmt.Println("lesser than Eleven")
	}
}
