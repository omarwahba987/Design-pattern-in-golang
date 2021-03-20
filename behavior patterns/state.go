package main

// https://github.com/yksz/go-design-patterns/blob/master/behavior/state.go
// https://en.wikipedia.org/wiki/Loose_coupling
// the problem: is when there's an object that needs to alter its behavior according to to its internal state.
//
//  the solution: we will create objects which represent various states and a context object whose behavior
// varies as its state object changes.
import (
	"fmt"
)

type State interface {
	Handle()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("ConcreteStateA.Handle()")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("ConcreteStateB.Handle()")
}


type Context struct {
	state State
}

func (c *Context) Request() {
	c.state.Handle()
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) GetState() State {
	return c.state
}



func main() {
	context := Context{new(ConcreteStateA)}
	context.Request()
	context.SetState(new(ConcreteStateB))
	context.Request()
}