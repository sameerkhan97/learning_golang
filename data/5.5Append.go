/*	* The signature of append is different from our custom Append function above. Schematically, it's like this:
	 		func append(slice []T, elements ...T) []T     */

package main

import "fmt"

func main{
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

	//append function
	a1=append(a1,11,12,13,14)
	fmt.Println("a1 : ",a1)
	fmt.Println("Length(a1) : ",len(a1))
	fmt.Println("Capacity(a1) : ",cap(a1))
	
	a1=append(a1,[]int{15,16,17,18,19,20,21}...) //Go doesnt allow to driectly append slice 				
	//hence using  ... operator in last breakdown slice into individual argeuments
	fmt.Println("a1 : ",a1)
	fmt.Println("Length(a1) : ",len(a1))
	fmt.Println("Capacity(a1) : ",cap(a1))
	
	//declare using make func : make(type,length,capacity)
	x:=make([]int,10,100)
	fmt.Println("Length(x) : ",len(x))
	fmt.Println("Capacity(x) : ",cap(x))
	
	//stack operations in Slice
	//removing elements from begining and end of slice
	z:=[]int{1,2,3,4,5}
	fmt.Println("z :",z);
	z1:=z[1:]			//removing first element
	fmt.Println("removing first element :",z1);
	z2:=z[:len(z)-1]	//removing last element
	fmt.Println("removing last element :",z2);
	z3:=append(z[:2],z[3:]...)	//removing mid element
	fmt.Println("removing middle element :",z3);
	fmt.Print("removing first element :",z1);	//after deleting we will get eror of references

}
