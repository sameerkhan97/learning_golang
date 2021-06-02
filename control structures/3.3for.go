/*The Go for loop is similar to—but not the same as—C's. It unifies for and while and there is no do-while. There are three forms, only one of which has semicolons.

	// Like a C for
	for init; condition; post { }

	// Like a C while	
	for condition { }

	// Like a C for(;;)
	for { }							
	
	//If you only need the second item in the range (the value), use the blank identifier, an underscore, to discard the first:
	sum := 0
	for _, value := range array {
	    sum += value
	}															*/

package main

import "fmt"

func main() {
	// Simple For loop in GO
	for i := 0; i < 10; i++ {
		fmt.Println("Hello GO ", i)
	}

	//If you're looping over an array, slice, string, or map, or reading from a channel, a range clause can manage the loop.
	primes := [6]int{2, 3, 5, 7, 11, 13} //this si an array ,it will be discussed later On
	for i := range primes {
		fmt.Println(primes[i])
	}
}
