package main

// https://github.com/tmrts/go-patterns/blob/master/structural/decorator.md
// structural pattern
// Decorator structural pattern allows extending the function of an existing object dynamically without altering its internals.
// Decorators provide a flexible method to extend functionality of objects.
// here we extend double function and can increase its functionality without modify it
// and here is an example for a function can return diffierent functions with same header

import (
	"log"
)

type Objectt func(int) int

func LogDecorate(fn Objectt) Objectt {
	return func(n int) int {
		log.Println("Starting the execution with the integer", n)

		result := fn(n)

		log.Println("Execution is completed with the result", result)

		return result
	}
}
func Double(n int) int {
	return n * 2
}
func Triple(n int) int {
	return n * 3
}



func main() {


	f := LogDecorate(Double)
	f(5)
	x:= LogDecorate(Triple)
	x(5)
	x(3)
	//fmt.Println(x(3))


}