package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
	// close(doneChan)
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true // use arrow for show from/to what the data is flowing
	close(doneChan)
}

func main() {
	// 1st way
	// done := make(chan bool) // create a channel
	// go greet("Nice to meet you!", done)
	// go greet("How are you?", done)
	// go slowGreet("How ... are ... you ...?", done) // use go to initiate a go routine
	// go greet("I hope you're liking the course!", done)
	// <- done // wait for done to emit something
	// <- done // using the same channel for many routines needs to end all of them
	// <- done
	// <- done

	// 2nd way
	// dones := make([]chan bool, 4) // create a channel
	// dones[0] = make(chan bool)
	// go greet("Nice to meet you!", dones[0])
	// dones[1] = make(chan bool)
	// go greet("How are you?", dones[1])
	// dones[2] = make(chan bool)
	// go slowGreet("How ... are ... you ...?", dones[2]) // use go to initiate a go routine
	// dones[3] = make(chan bool)
	// go greet("I hope you're liking the course!", dones[3])
	// for _, done := range dones {
	// 	<- done
	// }

	// 3rd way
	done := make(chan bool) // create a channel
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done) // use go to initiate a go routine
	go greet("I hope you're liking the course!", done)
	for range done {
		// fmt.Println(doneChan)
	}
}