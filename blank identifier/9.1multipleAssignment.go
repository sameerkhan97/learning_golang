/* * We've mentioned the blank identifier a couple of times now, in the context of for range loops and maps. The blank identifier can be
     assigned or declared with any value of any type, with the value discarded harmlessl.

    * The blank identifier in multiple assignment :
		* The use of a blank identifier in a for range loop is a special case of a general situation: multiple assignment.
		* If an assignment requires multiple values on the left side, but one of the values will not be used by the program, 
		  a blank identifier on the left-hand-side of the assignment avoids the need to create a dummy variable and makes it 
		  clear that the value is to be discarded.
*/
package main
import "fmt"
func main(){
	// a simplew use of blank identifier
	//here if we use any general variable then it will cause error
	//because its declared but not used
	_ , x:=isPrime(16)
	switch x {
		case true:fmt.Println("its a Prime Number")
		case false:fmt.Println("its not a Prime Number")
	}
}

func isPrime(x int)(i int,flag bool){
	x,flag=1,true
	for i:=2;i<=x/2;i++{
		if(x%i==0){
			x=i
			flag=false;
		}
	}
	return 
}
