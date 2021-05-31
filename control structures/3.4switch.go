/* Go's switch is more general than C's. The expressions need not be constants or even integers, the cases are evaluated top to bottom until a match is found, 
  and if the switch has no expression it switches on true. It's therefore possible—and idiomatic—to write an if-else-if-else chain as a switch. */

package main

import "fmt"

func main(){

	//switch statement
	//Go only runs the selected case, not all the cases that follow. In effect, 
	//the break statement is provided automatically in Go after every case.
	
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
		case today + 0:
			fmt.Println("Today.")
		case today + 1:
			fmt.Println("Tomorrow.")
		case today + 2:
			fmt.Println("In two days.")
		default:
			fmt.Println("Too far away.")
	}
	//cases can be presented in comma-separated lists.
	switch c {
		case ' ', '?', '&', '=', '#', '+', '%':
			return true
		}
	}
}
