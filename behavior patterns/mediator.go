package main
//https://github.com/yksz/go-design-patterns/blob/master/behavior/mediator.go
//https://www.tutorialspoint.com/design_pattern/mediator_pattern.htm
//used to reduce communication complexity between multiple objects or classes.
// This pattern provides a mediator class which normally handles all the communications between different classes and supports
// easy maintenance of the code by loose coupling
import (
	"container/list"
	"fmt"
)

type Mediator interface {
	AddColleague(colleague Colleague)
	Consultation(colleague Colleague)
}

type ConcreteMediator struct {
	colleagues *list.List
}

func NewConcreteMediator() *ConcreteMediator {
	return &ConcreteMediator{list.New()}
}

func (m *ConcreteMediator) AddColleague(colleague Colleague) {
	m.colleagues.PushBack(colleague)
}

func (m *ConcreteMediator) Consultation(colleague Colleague) {
	for e := m.colleagues.Front(); e != nil; e = e.Next() {
		if e.Value == colleague {
			switch e.Value.(type) {
			case *ConcreteColleagueA:
				fmt.Println("ConcreteColleagueA consulted")
			case *ConcreteColleagueB:
				fmt.Println("ConcreteColleagueB consulted")
			}
		}
	}
}

type Colleague interface {
	NeedsAdvice()
}

type ConcreteColleagueA struct {
	mediator Mediator
}

func (c *ConcreteColleagueA) NeedsAdvice() {
	c.mediator.Consultation(c)
}

type ConcreteColleagueB struct {
	mediator Mediator
}

func (c *ConcreteColleagueB) NeedsAdvice() {
	c.mediator.Consultation(c)
}

func main() {
	mediator := NewConcreteMediator()
	colleagueA := &ConcreteColleagueA{mediator}
	colleagueB := &ConcreteColleagueB{mediator}
	mediator.AddColleague(colleagueA)
	mediator.AddColleague(colleagueB)
	colleagueA.NeedsAdvice()
	colleagueB.NeedsAdvice()
}