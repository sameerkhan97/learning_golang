/* A switch can also be used to discover the dynamic type of an interface variable.Such a type switch uses the syntax of a type assertion with the keyword type inside 
   the parentheses If the switch declares a variable in the expression, the variable will have the corresponding type in each clause	*/

package main

import "fmt"

func main(){

	var t interface{}
	t = functionOfSomeType()
	switch t := t.(type) {	// reuse the name in such cases, in effect declaring a new variable with the same name but a different type in each case.
		default:
			fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
		case bool:
			fmt.Printf("boolean %t\n", t)             // t has type bool
		case int:
			fmt.Printf("integer %d\n", t)             // t has type int
		case *bool:
		    	fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
		case *int:
		    	fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}
