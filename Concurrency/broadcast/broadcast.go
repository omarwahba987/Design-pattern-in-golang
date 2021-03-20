package main

import (
	"log"
	"sync"

	broadcast "github.com/dustin/go-broadcast"
	"github.com/mitchellh/mapstructure"
)

//Message boradcasted
type Message struct {
	y string
	x int
}

var w sync.WaitGroup

// var bar barrier.Barrier
var ch chan interface{}

func main() {

	b := broadcast.NewBroadcaster(100)

	w.Add(1)
	go workerOne(b)

	d := Message{"message :", 0}

	// defer close(ch)
	for i := 0; i < 5; i++ {
		d.x = i
		log.Printf("Sending %v", d)

		b.Submit(d)
	}

	w.Wait()
}

func workerOne(b broadcast.Broadcaster) {
	ch = make(chan interface{})
	b.Register(ch)
	var mes Message
	for {
		v, _ := <-ch
		mapstructure.Decode(v, &mes)
		if mes.x != 4 {
			log.Printf("workerOne() reading : %v", v)
		} else {
			log.Printf("i am here")
			b.Unregister(ch)
			w.Done()
			return
		}
	}
}

func workerTwo(b broadcast.Broadcaster) {
	defer w.Done() //  -1
	ch := make(chan interface{})
	b.Register(ch)

	defer b.Unregister(ch)

	log.Printf("workerTwo read :  %v\n", <-ch)
}
