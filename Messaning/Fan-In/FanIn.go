package main

import (
	"fmt"
	"sync"
	// "sync"
)

func main() {

	channelList := make([]chan int, 2) // another way to creat an array of channel
	for i := 0; i < 1; i++ {
		channelList = append(channelList, make(chan int))
	}

	FChannel := make(chan int, 6)
	for i := 0; i <= 5; i++ {
		FChannel <- i // contain values from 0 to 5
	}
	close(FChannel)

	channelList = append(channelList, FChannel)

	SChannel := make(chan int, 5)
	for j := 6; j <= 10; j++ {
		SChannel <- j // contain values from 6 to 10
	}
	channelList = append(channelList, SChannel)
	close(SChannel)

	output := Merge(channelList)

	fmt.Println("merged channels value : ")
	for j := 0; j <= 10; j++ {
		fmt.Println(<-output)
	}
}

// Merge different channels in one channel
func Merge(cs []chan int) <-chan int { // change (cs ...<-chan int) to (cs []chan int)
	var wg sync.WaitGroup

	out := make(chan int, 11)

	// Start an send goroutine for each input channel in cs. send
	// copies values from c to out until c is closed, then calls wg.Done.
	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}
	// Start a goroutine to close out once all the send goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

//https://github.com/tmrts/go-patterns/blob/master/messaging/fan_in.md
//https://medium.com/rungo/anatomy-of-channels-in-go-concurrency-in-go-1ec336086adb                   channel summery
