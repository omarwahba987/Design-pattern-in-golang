package main

import (
	"fmt"
)

type sandwich interface {
	Operation() string
}

type Concretesandwich struct{
	component sandwich
}

func (c *Concretesandwich) Operation() string {
	return "simple sandwich"
}

type ConcreteDecoratorA struct {
	component sandwich
}

func (d *ConcreteDecoratorA) Operation() string {
	if d.component == nil {
		return ""
	} else {
		return d.component.Operation() + " and katchab"
	}
}

type ConcreteDecoratorB struct {
	component sandwich
}

func (d *ConcreteDecoratorB) Operation() string {
	if d.component == nil {
		return " "
	} else {
		return d.component.Operation() + " and maio"
	}
}

func main() {
	component := &ConcreteDecoratorA{&ConcreteDecoratorB{new(Concretesandwich)}}
	component2 := &Concretesandwich{}
	component3 := &ConcreteDecoratorA{new(Concretesandwich)}
	fmt.Println(component.Operation())
	fmt.Println(component2.Operation())
	fmt.Println(component3.Operation())
}