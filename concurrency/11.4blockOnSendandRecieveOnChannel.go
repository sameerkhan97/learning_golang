/* * Any send on a channel is blocked until there is  another goroutine to receive it. 

   * The receive on a channel is blocked unless there is another goroutine to send to that channel.
*/
package main
import (
	"fmt"
	"time"
)
func main(){
	a:=make(chan int)	//defining channel using function make.
	fmt.Println("Sending values to channels")
	go send(a)
	fmt.Println("Recieving values to channels")
	go recieve(a)
	time.Sleep(time.Second*2);
}
func send(ch chan int){
	ch<-1
	fmt.Println("Sending value to channel complete")
}
func recieve(ch chan int){
	time.Sleep(time.Second * 1)	//here recieve function is stop for some seconds hence send() is blocked from executing and after seconds recieve start
	val:=<-ch
	fmt.Println("Value recieved = ",val," in recieve function")
}