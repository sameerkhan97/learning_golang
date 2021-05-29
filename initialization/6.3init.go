/* * Each source file can define its own niladic init function to set up whatever state is required
   * Each file can have multiple init functions
   * Multiple init() functio are be called one after another
   * init() is called after all the variable declarations in the package have evaluated their initializers,
     and those are evaluated only after all the imported packages have been initialized. 
   * init() is used to do task thgat we want to do before everythinng else.
	*/

package main
import "fmt"
func init(){
	fmt.Println("In init 1 ")
}
func init(){
	fmt.Println("In init 2 ")
}
func main(){
	a:=10;
	fmt.Println("In Main ",a)
} 
func init(){
	fmt.Println("In init 3 ")
}
