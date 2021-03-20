package main


// https://golangbyexample.com/chain-of-responsibility-design-pattern-in-golang/
// https://github.com/yksz/go-design-patterns/blob/master/behavior/chain_of_responsibility.go
// the problem: is when a tightly coupled objects that makes it hard to coordinate between the object that sends
// a request and the other objects that receives request to process.

// the solution: CoR lets you create a chain of request handlers. For every incoming request,
// it is passed through the chain and each of the handler:
// Processes the request or skips the processing.
// Decides whether to pass the request to the next handler in the chain or not.

import (
	"fmt"
)

type Handler interface {
	Request(flag bool)
}

type ConcreteHandlerA struct {
	next Handler
}

func (h *ConcreteHandlerA) Request(flag bool) {
	fmt.Println("ConcreteHandlerA.Request()")
	if flag {
		h.next.Request(flag)
	}
}

type ConcreteHandlerB struct {
	next Handler
}

func (h *ConcreteHandlerB) Request(flag bool) {
	fmt.Println("ConcreteHandlerB.Request()")
}

func main() {
	handlerA := &ConcreteHandlerA{new(ConcreteHandlerB)}
	handlerA.Request(false)
}