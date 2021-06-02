/*	* Go has two allocation primitives, the built-in functions new and make.
	
	* new first. It's a built-in function that allocates memory, but unlike its namesakes in some other languages it does not initialize
	  the memory, it only zeros it. That is, new(T) allocates zeroed storage for a new item of type T and returns its address, a value of
	  type *T. In Go terminology, it returns a pointer to a newly allocated zero value of type T

	* Since the memory returned by new is zeroed, it's helpful to arrange when designing your data structures that the zero value of each 
	 type can be used without further initialization. This means a user of the data structure can create one with new and get right to work.
*/

package main

import "fmt"

type demo struct {
	x int
	t bool
	f float64
}

func main() {
	obj := new(demo)
	fmt.Println(obj.x, obj.t, obj.f) //see values initialized by new
}
