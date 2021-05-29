/* * Go's defer statement schedules a function call (the deferred function) to be run immediately 
     before the function executing the defer returns
   
   * Deferring a call to a function such as Close has two advantages. First, it guarantees that you will never forget to close the file,
     a mistake that's easy to make if you later edit the function to add a new return path. Second, it means that the close sits near the open,
     which is much clearer than placing it at the end of the function.

   * The arguments to the deferred function (which include the receiver if the function is a method) 
     are evaluated when the defer executes, not when the call executes

	
   * Deferred functions are executed in LIFO order                  
   * Function with difer keyword are call when xurrent function's execution completes  */


package main
import "fmt"
func main(){
	//simple example of defer function
	for i:=1;i<=10;i++{
		defer fmt.Println(i);	//show that Deferred functions are executed in LIFO order 
	}

	//some complex example 
	a();
}
func a(){
	fmt.Println("a begins")
	defer b()
	fmt.Println("a ends")
}
func b(){
	fmt.Println("in b")
}