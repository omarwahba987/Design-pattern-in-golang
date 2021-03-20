//// Prototype patterns is required, when object creation is time consuming, and costly operation
///Clone is the simplest approach to implement prototype pattern. However,
//   it is your call to decide how to copy existing object based on your business model.
///
package main

import (
	"fmt"
)

type Prototype interface {
	Clone()
}

/*********************************************************************************/
type ConcretePrototype struct {
	name string
}

func (p *ConcretePrototype) Clone() *ConcretePrototype {
	return &ConcretePrototype{p.name}
	// return &(*p)
}

/************************************************************************************/
func (p *ConcretePrototype) String() string {
	return "ConcretePrototype [name=" + p.name + "]"
}
func main() {
	prototype := &ConcretePrototype{"prototype1"}
	p := prototype
	fmt.Println(prototype)

	//prototype.name = "fff"
	f := prototype.Clone()
	prototype.name = "kkkkkk"
	fmt.Println("----", p)
	fmt.Println(prototype)
	fmt.Println(prototype.Clone())
	fmt.Println("44444", f)
}
