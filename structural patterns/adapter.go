package main
// https://github.com/yksz/go-design-patterns/blob/master/structure/adapter.go
// structural pattern
// Adapter pattern works as a bridge between two incompatible interfaces
// Convert the interface of a class into another interface clients expect.
// Adapter lets classes work together that couldn't otherwise because of incompatible interfaces.
// let a class perform as a nother class and do its functionalties
import (
	"fmt"
)


type Adaptee struct{}

func (a *Adaptee) OldMethod() {
	fmt.Println("1")
}
//////////////////////////////////
type Target interface {
	RequiredMethod()
}

type Adapter struct {
	adaptee *Adaptee
}

func NewAdapter() *Adapter {
	return &Adapter{new(Adaptee)}
}

func (a *Adapter) RequiredMethod() {
	a.adaptee.OldMethod()
	fmt.Println("2")

}

func main() {
	target := NewAdapter()
	target.RequiredMethod()
}