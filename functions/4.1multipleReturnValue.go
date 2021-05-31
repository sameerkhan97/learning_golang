//Functions and methods can return multiple values in GO

package main

import "fmt"

func main(){
	msg1:="hello"
	msg2:="Sam"	
	sayMessage(&msg1,&msg2)
	fmt.Println(msg1,msg2);
	s:=sum("Sum ",1,2,3,4,5)
	fmt.Println("sum is : ",s)
	//Anonymus Function
	f:=func(){
		fmt.Println("HEllo Go ")
	}
	f()
	d,err:=devide(5.0,2.0)
	if(err!=nil){
		fmt.Println(err)
		return 
	}
	fmt.Println(d)
	f()
}

//returning a value
func sum(msg string,value ...int) ( int) {	//this ... abrevetation should be onely one and should be at last 
	fmt.Println(value)
	for _,v:=range value{
		result+=v;
	}
	return result
}

//returning multiple value
func devide(a,b float64) (float64,error){
	if b==0.0{
		return 0.0,fmt.Errorf("Cannot devide by zero ")
	}
	return a/b,nil
}

//pass by reference
func sayMessage(msg1,msg2 *string){
	fmt.Println(*msg1,*msg2)
	*msg2="zaid"
	fmt.Println(*msg1,*msg2)
}
