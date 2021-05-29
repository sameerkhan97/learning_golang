package main

import(
	"fmt"
	"time"
)

func main(){
	go say(5)		//both say and go routine say run concurrently
	fmt.Scanln();
	go say(5)
}

func say(x int){
	for i:=0;i<=x;i++{
		fmt.Println(i)
		time.Sleep(time.Second);
	}
}