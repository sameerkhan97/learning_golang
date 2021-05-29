/* *  A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space.
   * It is lightweight, costing little more than the allocation of stack space. And the stacks start small, so they are cheap, 
     and grow by allocating (and freeing) heap storage as required.
   * Goroutines are multiplexed onto multiple OS threads so if one should block, such as while waiting for I/O, others continue to run.
   * Prefix a function or method call with the go keyword to run the call in a new goroutine
   * When the call completes, the goroutine exits, silently.

*/ 


package main

import(
	"fmt"
	"time"
)

func main(){
	go say("Hello")		//both say and go routine say run concurrently
	go say("There")
	time.Sleep(time.Second)	//main function would wait here for a second before ending and goroutine will get some time to complete
}

func say(s string){
	for i:=0;i<3;i++{
		fmt.Println(s)
		time.Sleep(time.Millisecond);
	}
}