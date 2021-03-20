package main

import "fmt"

//Count return 1-99 no.
func Count(start int, end int) chan int {
	ch := make(chan int)

	go func(ch chan int) {
		for i := start; i <= end; i++ {
			// Blocks on the operation
			ch <- i
		}

		close(ch)
	}(ch)

	return ch
}

func main() {
	fmt.Println("No bottles of juice on the wall")

	for i := range Count(1, 99) {
		fmt.Println("Pass it around, put one up,", i, "bottles of juice on the wall")
		// Pass it around, put one up, 1 bottles of juice on the wall
		// Pass it around, put one up, 2 bottles of juice on the wall
		// ...
		// Pass it around, put one up, 99 bottles of juice on the wall
	}

	fmt.Println(100, "bottles of juice on the wall")
}

// Yields a sequence of values one at a time
// https://github.com/tmrts/go-patterns/blob/master/concurrency/generator.md
