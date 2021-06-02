/* * In Go you can also create a buffered channel. A buffered channel has some capacity to hold data hence for a buffered channel:
			~ Send on a buffer channel only blocks if the buffer is full
			~ Receive is only block is channel is empty
   
   * This is the syntax for creating a buffered channel using the make function
			a = make(chan , capacity)
     The second argument specifies the capacity of the channel. Unbuffered channel is of zero capacity.
*/

package main

import "fmt"

func main() {
	//The channel is created with a capacity of one. Hence sending to the channel is not blocked and the value is stored in the channel's buffer
	ch := make(chan int, 1)
	ch <- 1
	//ch<-1
	//this line will cause error coz second sent to the channel is blocked because buffer is full and hence it results
	// in a deadlock situation because the program cannot proceed and that is why as you can see the output is
	// o/p :- fatal error: all goroutines are asleep - deadlock!
	fmt.Println("Sending value to channnel complete")

	val := <-ch
	val = <-ch
	//we tried a second receive from the channel and it resulted in a deadlock situation because the program
	//cannot proceed as the channel is empty and there is nothing to receive.
	fmt.Printf("Receiving Value from channel finished. Value received: %d\n", val)
}

