/* * When panic is called, including implicitly for run-time errors such as indexing a slice out of bounds or failing a type assertion, it immediately stops execution 
     of the current function and begins unwinding the stack of the goroutine, running any deferred functions along the way.If that unwinding reaches the top of the 
     goroutine's stack, the program dies.
   * It is possible to use the built-in function recover to regain control of the goroutine and resume normal execution.
   * A call to recover stops the unwinding and returns the argument passed to panic. Because the only code that runs while unwinding
     is inside deferred functions.
NOTE : recover is only useful inside deferred functions.If the recover function is not within the defer function then it will not stop panic */

package main

import "fmt"

func main(){
	a:=[]string{"one","two"}
	checkAndPrint(a,2)
}
//checkAndPrint which checks and prints slice element at an index passed in the argument
func checkAndPrint(a []string,index int){	
	//defer function named handleOutIfBounds as well at the start of the function checkAndPrint.
	//This function  contains  the  call to recover function
	defer handleOutOfBound();
	
	if index>len(a)-1{
		//if the index passed is greater than the length of the array then the program panics.
		panic("Out of Bound aaccess for slice")
	}
	fmt.Println(a[index]);
}
func handleOutOfBound(){
	if r:=recover();r!=nil{
		//The recover function will catch the panic and we can also print the message from the panic. 
		fmt.Println("Recovering from Panic : ",r)
	}
}
