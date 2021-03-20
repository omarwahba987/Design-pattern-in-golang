package main
// https://github.com/yksz/go-design-patterns/blob/master/behavior/visitor.go
//It is used when we have to perform an operation on a group of similar kind of Objects with nonsimiler implementation .
// With the help of visitor pattern, we can move the operational logic from the objects to another class.

import (
	"fmt"
)


type Element interface {
	Accept(Visitor)
}

type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	fmt.Println("ConcreteElementA.Accept()")
	visitor.VisitA(e)
}

type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	fmt.Println("ConcreteElementB.Accept()")
	visitor.VisitB(e)
}

type Visitor interface {
	VisitA(*ConcreteElementA)
	VisitB(*ConcreteElementB)
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitA(element *ConcreteElementA) {
	fmt.Println("ConcreteVisitor.VisitA()")
}

func (v *ConcreteVisitor) VisitB(element *ConcreteElementB) {
	fmt.Println("ConcreteVisitor.VisitB()")
}

func main() {
	visitor := new(ConcreteVisitor)
	elementA := new(ConcreteElementA)
	elementB := new(ConcreteElementB)
	elementA.Accept(visitor)
	elementB.Accept(visitor)
}