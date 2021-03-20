package main

import (
	"fmt"
	"sync"

	"github.com/pwaller/barrier"
)

func main() {
	var w sync.WaitGroup
	defer w.Wait() // Main should wait for its goroutines
	var b barrier.Barrier

	w.Add(1)
	go func() {
		println("2")
		defer w.Done() //  -1
		defer b.Fall()
		println("GO! - Second")
		// <-b.Barrier() // Many goroutines can wait on the barrier
		println("3")
	}()

	w.Add(1)
	go func() {
		println("1")
		println("1")
		println("1")
		defer w.Done() //  -1
		defer b.Fall()
		println("GO! - First")
		// When this goroutine happens to return,
		// all barrier waits can be passed.
		return
	}()

}

func wow() {
	fmt.Println("i am here now")
}
