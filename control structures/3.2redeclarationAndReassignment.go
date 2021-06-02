
/*   In a := declaration a variable v may appear even if it has already been declared, provided this declaration is in the same scope as the existing declaration
     of v (if v is already declared in an outer scope , the declaration will create a new variable §) , the corresponding value in the initialization is assignable 
     to v, and there is at least one other variable that is created by the declaration.   */

package main

import "fmt"

func main() {

	//the := works as a short declaration form
	x, y := 110, 120
	fmt.Println(x, y)
}

