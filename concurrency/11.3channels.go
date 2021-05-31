/* * Channel is a data type in Go which provides synchrounization and communication between goroutines. They can be thought of as
     pipes which is used by goroutines to communicate. 
 Declaring Channels:	 
   * Each channel variable can hold data only of a particular type. Go uses special keyword "chan" while declaring a channel.
     Default value of a channel is nil
   * To define the channel we can use the inbuilt function make. A channel is internally represented by a hchan struct who has many elements.
   * When using make, an instance of hchan struct is created and all the fields are initialized to their default values.
Operations on Channel :
   ~ Send Operation:
   * The send operation is used to send data to the channel
    Eg ch <- val , ch is the channel variable , val is what being sent to the channel , Note that data type of val and data type of channel should match.
   * Any send on a channel is blocked until there is  another goroutine to receive it. 
   ~ Receive Operation:
   * The receive operation is used to read data from the channel. Eg:- val := <- ch 
   * A very important point to note about the receive operation is that a particular value sent to the channel 
     can only be received once in any of the goroutine
   * The receive on a channel is also blocked unless there is another goroutine to send to that channel.
Buffered Channel :
   * In Go you can also create a buffered channel. A buffered channel has some capacity to hold data hence for a buffered channel:
   * Send on a buffer channel only blocks if the buffer is full
		Receive is only block is channel is empty
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
}
func recieve(ch chan int){
	val:=<-ch
	fmt.Println("Value recieved = ",val," in recieve function")
}
