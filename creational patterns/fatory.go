//-----------------------------
// class diagram discuss: https://stackoverflow.com/questions/53895044/clarifying-uml-class-diagram-of-factory-design-pattern
//   https://medium.com/@haluan/go-factory-design-pattern-e5301e0f3283
/*  we create object without exposing the creation logic to client
and the client use the same common interface
to create new type of object*/
//-------------------------------------------------------------------------------------------------

package main

import (
	"fmt"
)

/*********************************************************************************************/

/****************************/
/********************************interface *********************************************************************/
type Productt interface { ////shape   a
	method()
}

/**************************************************************************************************************/
type ConcreteProduct struct{} ///circle

func (p *ConcreteProduct) method() { //draw circle
	fmt.Println("create three")
	fmt.Println("ConcreteProduct.method()")
}
/////////////////////////////////////////////////////////////
type factory interface {
	factoryMethod() Productt
}

/******************************************************************************************/
type Creator struct { ////
	factory factory
}

func (c *Creator) Operation() {
	//var prod Productt
	//prod.method()
	product := c.factory.factoryMethod()
	fmt.Println("create two:  ")
	product.method()
}

/************************************************/ ////
type ConcreteCreator struct{}

func (c *ConcreteCreator) factoryMethod() Productt {
	fmt.Println("create one")
	return new(ConcreteProduct)
}
/////////////////////////////////////////////////////////////
//type ConcreteCreator2 struct{}
//
//func (c *ConcreteCreator2) factoryMethod() Productt {
//	fmt.Println("create one")
//	return new(ConcreteProduct)
//}
func main() {
	creator := Creator{new(ConcreteCreator)}
	fmt.Println("-------------------------------------creator--------------------------------------------")
	creator.Operation()
}
