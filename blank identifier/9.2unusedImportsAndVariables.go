/* * It is an error to import a package or to declare a variable without using it. Unused imports bloat the program and slow compilation, 
     while a variable that is initialized but not used is at least a wasted computation and perhaps indicative of a larger bug.

   * When a program is under active development, however, unused imports and variables often arise and it can be annoying to delete them 
     just to have the compilation proceed, only to have them be needed again later. 
 
   * The blank identifier provides a workaround.
*/ 

//This half-written program has unused imports (fmt and io) and an unused variable (fd), so it will not compile
//To silence complaints about the unused imports, use a blank identifier to refer to a symbol from the imported package.
//Similarly, assigning the unused variable fd to the blank identifier will silence the unused variable error. 
//This version of the program does compile.

package main

import (
	"fmt"
	"log"
	"os"
)

//By convention, the global declarations to silence import errors should come right after the imports and be commented

var _ = fmt.Printf // For debugging; delete when done.
func main() {
	fd, err := os.Open("multipleAssignment.go")
	if err != nil {
		log.Fatal(err)
	}
	_ = fd
}
