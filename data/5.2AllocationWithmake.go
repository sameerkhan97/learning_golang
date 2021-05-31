/*  * Go has two allocation primitives, the built-in functions new and make.
	* the new does not initialize the memory, it only zeros it.But sometimes the zero value isn't good enough 
	* The built-in function make(T, args) serves a purpose different from new(T).
	* make creates slices, maps, and channels only, and it returns an initialized (not zeroed) value of type T (not *T).	*/

package main

import "fmt"

func main(){
	v:=make([]int,10,100)
	// Allocates an array of 100 ints and then creates a slice structure with length 10 and 
	// a capacity of 100 pointing at the first 10 elements of the array. 

	var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
	var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

	// Unnecessarily complex:
	var p *[]int = new([]int)
	*p = make([]int, 100, 100)

	// Idiomatic:
	v := make([]int, 100)
}
//NOTE : Remember that make applies only to maps, slices and channels and does not return a pointer.  
 
