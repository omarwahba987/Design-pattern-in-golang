package main

import "fmt"

func main() {
	mainChannel := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		mainChannel <- i // contain values from 0 to 10
	}
	close(mainChannel)

	output := Split(mainChannel, 2)

	fmt.Println(" length of output channel is : ", len(output))
	fmt.Println(" length of the first one is : ", len(output[0]))
	fmt.Println(" length of the second one is : ", len(output[1]))

	for _, c := range output {
		for value := range c {
			fmt.Println(value)
		}
	}

}

// Split a channel into n channels that receive messages in a round-robin fashion.
func Split(ch <-chan int, n int) []chan int {
	cs := make([]chan int, 0)
	for i := 0; i < n; i++ {
		cs = append(cs, make(chan int, 5))
	}

	// Distributes the work in a round robin fashion among the stated number
	// of channels until the main channel has been closed. In that case, close
	// all channels and return.
	distributeToChannels := func(ch <-chan int, cs []chan int) {
		// Close every channel when the execution ends.
		defer func(cs []chan int) {
			for _, c := range cs {
				close(c)
			}
		}(cs)

		for {
			for _, c := range cs {
				select {
				case val, ok := <-ch:
					if !ok {
						return
					}

					c <- val
				}
			}
		}
	}

	go distributeToChannels(ch, cs)

	return cs
}

//https://github.com/tmrts/go-patterns/blob/master/messaging/fan_out.md
//https://www.techopedia.com/definition/13205/round-robin
//https://www.geeksforgeeks.org/unidirectional-channel-in-golang/         (Undirected channel)
