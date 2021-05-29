/*	* Slices hold references to an underlying array.
	* If you assign one slice to another, both refer to the same array.
	* If a function takes a slice argument, changes it makes to the elements of the slice will be visible to the caller, 
	  analogous to passing a pointer to the underlying array            */
	package main
	import "fmt"
	func main(){
		a:=[]int{1,2,3}
		fmt.Println("a : ",a)
		fmt.Println("Length(a) : ",len(a))
		fmt.Println("Capacity(a) : ",cap(a))		//slics has an additional capacity function
		//referncing
		b:=a;
		b[1]=5
		fmt.Println("a : ",a)
		fmt.Println("b : ",b)
	
		//slice Declaration
		a1:=[]int{1,2,3,4,5,6,7,8,9,10}
		b1:=a1[:]	//slice of all element
		c:=a1[3:]	//slice fom 4th element to end
		d:=a1[:6]	//slice of first 6 element
		e:=a1[3:6]	//slice of 4th , 5th and 6th element 
		fmt.Println("a1 : ",a1)
		fmt.Println("Length(a1) : ",len(a1))
		fmt.Println("Capacity(a1) : ",cap(a1))
		fmt.Println("b1 : ",b1)
		fmt.Println("c : ",c)
		 fmt.Println("d : ",d)
		fmt.Println("e : ",e)
	}