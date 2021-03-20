package main
// https://github.com/tmrts/go-patterns/blob/master/behavioral/strategy.md
// Strategy behavioral design pattern enables an algorithm's behavior to be selected at runtime
// *with same arguments* <- the difference from builder design pattern
// It defines algorithms, encapsulates them, and uses them interchangeably.

import "fmt"

type Operator interface {
Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type Addition struct{}

func (Addition) Apply(lval, rval int) int {
	return lval + rval
}

type Multiplication struct{}

func (Multiplication) Apply(lval, rval int) int {
	return lval * rval
}

func main() {
	add := Operation{Addition{}}
	fmt.Println(add.Operate(3, 5) )

	add = Operation{Multiplication{}}
	fmt.Println(add.Operate(3, 5))

}
