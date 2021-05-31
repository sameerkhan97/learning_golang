/*	* The functions live in the fmt package and have capitalized names: fmt.Printf, fmt.Fprintf, fmt.Sprintf and so on.
	* The string functions (Sprintf etc.) return a string rather than filling in a provided buffer.
	* You don't need to provide a format string. For each of Printf, Fprintf and Sprintf there is another pair of functions, for instance Print and Println.
	  These functions do not take a format string but instead generate a default format for each argument.
	* The Println versions also insert a blank between arguments and append a newline to the output while the Print versions add blanks
	  only if the operand on neither side is a string.
	* For maps, Printf and friends sort the output lexicographically by key.  	*/

package main

import "fmt"

type T struct {
	a int
	b float64
	c string
}

func main(){
	fmt.Printf("Hello %d\n", 23)			//Print versions add blanks only if the operand on neither side is a string.
	fmt.Println("Hello", 23)			//Println insert a blank between arguments and append a newline to the output
	fmt.Println(fmt.Sprint("Hello ", 23))		//Sprintf return a string rather than filling in a provided buffer
	fmt.Printf("%T",12.46)				//prints the type of variable

	//printing structure
	t := T{ 7, -2.35, "abc\tdef" }
	fmt.Printf("%v\n", t) 		
	fmt.Printf("%+v\n", t)				//%+v annotates the fields of the structure with their names
	fmt.Printf("%#v\n", t)				// for any value the alternate format %#v prints the value in full Go syntax.
}	 
