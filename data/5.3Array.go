/*	* Arrays are useful when planning the detailed layout of memory and sometimes can help avoid allocation, 
      but primarily they are a building block for slices,
	* There are major differences between the ways arrays work in Go and C. In Go,
			~Arrays are values. Assigning one array to another copies all the elements.
			~In particular, if you pass an array to a function, it will receive a copy of the array, not a pointer to it.
			~The size of an array is part of its type. The types [10]int and [20]int are distinct						*/ 

package main
import "fmt"
func main(){
	//simple declaration of an array in GO
	v:= [10]int{1,2,3,4,5,6,7,8,9,10}	
	fmt.Println("v :",v);
	
	y:=sum(&v)			//passing array to a function by reference
	fmt.Println("Sum : ",y)
	fmt.Println("v :",v)

	//copying an array
	x:=v
	fmt.Println("x :",x)
	x[2]=99				//changing value in copy array , will it effect on v?
	fmt.Println("v :",v)
	fmt.Println("x :",x)
}
func sum(a *[10]int) (sum int){
	for i:=0;i<len(a);i++{
		sum=sum+a[i];
	}
	a[0]=100;			//changiing values of array in function to check if it effect v
	return 
}
