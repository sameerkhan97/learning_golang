package main

import (
	"fmt"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go say("Hello") //both say and go routine say run concurrently
	wg.Add(1)
	go say("There")
	wg.Wait()
}

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done()
}
