/* * Library routines must often return some sort of error indication to the caller.
   * Go's multivalue return makes it easy to return a detailed error description alongside the normal return value
   * For example, as we'll see, os.Open doesn't just return a nil pointer on failure, 
     it also returns an error value that describes what went wrong.
   * By convention, errors have type error, a simple built-in interface.
		
  		type error interface {
    		Error() string
		}
*/

package main
import (
	"fmt"
	"errors"
)
func main(){
	//sumOf() calculate sum of elements between given range
	//it will give error when start of range is greater than end of range  
	total,err:=sumOf(10,2);

	if err!=nil{
		fmt.Println("There is an Error ",err)
	}else{
		fmt.Println("Sum = ",total)
	}
}
func sumOf(start,end int) (int , error){
	if start>end{
		return 0,errors.New("Start is greater than end")
	}
	total:=0;
	for i:=start;i<=end;i++{
		total+=i;
	}
	return total,nil
}