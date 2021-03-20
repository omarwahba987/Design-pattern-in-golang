package main

// https://github.com/yksz/go-design-patterns/blob/master/structure/flyweight.go
// structural pattern
// this pattern provides way to decrease object count thus improving application required objects structure.
// Flyweight pattern is used when we need to create a large number of similar objects
// to less number of objects to reduces the memory usage


import (
	"fmt"
)

type FlyweightFactory struct {
	pool map[string]*Flyweight
}

func (f *FlyweightFactory) GetFlyweight(str string) *Flyweight {
	flyweight := f.pool[str]				// if not == nil then its exist and return the old one and didnt make a new object
	fmt.Println("exist before >> ", flyweight)
	if flyweight == nil {
		fmt.Println("new: " + str)
		f.pool[str] = &Flyweight{str}
	}
	return flyweight
}

type Flyweight struct {
	str string
}

func main() {
	factory := &FlyweightFactory{make(map[string]*Flyweight)}
	factory.GetFlyweight("a")
	factory.GetFlyweight("p")
	factory.GetFlyweight("p")
	factory.GetFlyweight("l")
	factory.GetFlyweight("e")
	factory.GetFlyweight("!")
	factory.GetFlyweight("o")
	factory.GetFlyweight("r")
	factory.GetFlyweight("a")
	factory.GetFlyweight("n")
	factory.GetFlyweight("g")
	factory.GetFlyweight("e")
	factory.GetFlyweight("!")
}
