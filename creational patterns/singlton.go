/////  --------------- https://www.geeksforgeeks.org/singleton-design-pattern/
//The singleton pattern is a design pattern that restricts the instantiation of a class to one object
//for example a single DB connection shared by multiple objects as creating a separate DB connection
//for every object may be costly.

package main

import "fmt"

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() *singleton {

	if instance == nil {

		instance = new(singleton)

	}

	return instance

}

func (s *singleton) AddOne() int {

	s.count++

	return s.count

}
func main() {
	singleton1 := GetInstance()

	fmt.Println(singleton1.AddOne())
	fmt.Println(singleton1.AddOne())
	fmt.Println(instance)
	singleton2 := GetInstance()
	fmt.Println("********************", singleton2)
}

