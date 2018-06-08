package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// takes 9 seconds to run
	// longTask()
	// longTask()
	// longTask()

	channel := make(chan int)
	for i := 1; i <= 3; i++ {
		go longTask(channel)
	}

	// channel waits to receive data and freezes main()
	for i := 1; i <= 3; i++ {
		fmt.Println("Takes", <-channel, "seconds")
	}
}

func longTask(channel chan int) {
	delay := rand.Intn(5)
	fmt.Println("starting long task")
	time.Sleep(time.Duration(delay) * time.Second)
	fmt.Println("end long task")
	channel <- delay
}
