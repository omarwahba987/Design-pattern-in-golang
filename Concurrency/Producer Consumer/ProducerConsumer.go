package main

import (
	"fmt"
	"time"
)

func calculateNextInt(prev int) int {
	time.Sleep(1 * time.Second) // pretend this is an expensive operation
	return prev + 1
}

func main() {
	data := make(chan int)
	quit := make(chan interface{})

	// producer
	go func() {
		var i = 0
		for {
			i = calculateNextInt(i)
			select {
			case data <- i:
			case <-quit:
				close(data)
				return
			}
		}
	}()

	// consumer
	for i := range data {
		fmt.Printf("i=%v\n", i)

		if i >= 5 {
			quit <- data
			close(quit)
		}
	}
}

//https://www.instana.com/blog/producer-consumer-in-go/
