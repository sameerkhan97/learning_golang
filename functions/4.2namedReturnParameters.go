/* The return or result "parameters" of a Go function can be given names and used as regular variables, just like the incoming parameters. When named, they are initialized 
   to the zero values for their types when the function begins; if the function executes a return statement with no arguments, the current values of the result parameters
   are used as the returned values.*/

package main

import "fmt"

func main(){
	x,y:=92,302;
	sum:=add(x,y);
	fmt.Println(x," + ",y," = " ,sum);
}
func add(a,b int) (sum int){
	sum=a+b;
	return 
}
