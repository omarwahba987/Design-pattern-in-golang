package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyCoroutine struct {
	line int
	// coroutine.Embeddable
}

func (c *MyCoroutine) Start() {
	for {
		switch msg := c.Recv().(type) {
		case int:
			fmt.Printf("Message %v: Received integer [%v]\n", c.line, msg)
		case string:
			fmt.Printf("Message %v: Received string [%s]\n", c.line, msg)
		}
		c.line++
	}

	fmt.Println("This never prints")
}

func main() {
	r := coroutine.Start(&MyCoroutine{})
	for i := 0; i < 100; i++ {
		r.Send(i)
		r.Send("msg " + strconv.Itoa(i))
	}
	time.Sleep(5 * time.Millisecond)
	r.Stop()
	time.Sleep(5 * time.Millisecond)
	fmt.Println("All done. All messages should have been printed, but not the println at the end of Start.")
}

//https://subscription.packtpub.com/book/application_development/9781783983483/1/ch01lvl1sec10/understanding-goroutines-versus-coroutines
//http://www.golangpatterns.info/concurrency/coroutines
