/* * The usual way to report an error to a caller is to return an error as an extra return value. The canonical Read method is a well-known
     instance; it returns a byte count and an error. But what if the error is unrecoverable? Sometimes the program simply cannot continue.
	
   * There is a built-in function panic that in effect creates a run-time error that will stop the program
   
   * The function takes a single argument of arbitrary type—often a string—to be printed as the program dies
	
   * It's also a way to indicate that something impossible has happened, such as exiting an infinite loop.
*/

package main
import "fmt"
func main(){
	x:=Devision(9,0);
	fmt.Println(x);
}

// A toy implementation of cube root using Newton's method.
func Devision(x ,y float64) float64 {
    if y==0{
		panic("Cannot devide by zero ")
	}
	d:=x/y;
	return d;
}