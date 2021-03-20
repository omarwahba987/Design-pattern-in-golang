package main
//https://github.com/yksz/go-design-patterns/blob/master/creation/object_pool.go
//https://www.geeksforgeeks.org/object-pool-design-pattern/
//object pool is used in situations where the cost of initializing a class instance is very high.
//Basically, an Object pool is a container which contains some amount of objects. So, when an object is taken from the pool,
// it is not available in the pool until it is put back

import (
	"fmt"
	"time"
)

type Object struct {
	name int
}
//
//func (o *Object)hold() {
//	for i := 0; i < 100000000; i++ {
//
//	}
//}

type Pool chan *Object

func New(total int) *Pool {
	p := make(Pool, total)

	for i := 0; i < total; i++ {
		x:=&Object{i}
		p <- x
	}

	return &p
}
func print(p *Pool){
	for{
		time.Sleep(35*time.Millisecond)
		select {
		case obj := <-*p:
			//print2(p,obj)
			fmt.Println(">",obj.name)
			////time.Sleep(2*time.Second)
			//
			// *p <- obj
		default:
			fmt.Println("dont")
			// No more objects left — retry later or fail

		}
	}
}
func add(p *Pool){
	for  {
		time.Sleep(230*time.Millisecond)
		x:=&Object{}
		*p <- x
	}

}
//func print2(p*Pool,o *Object){
//	fmt.Println(">",">",o.name)
//	time.Sleep(2000 *time.Millisecond)
//	o.hold()
//	*p <- o
//}

func main (){
	//var p Pool
	p:= New(1)
	go print(p)
	go add(p)
	time.Sleep(10 *time.Second)

	//print(p)

	//select {
	//case obj := <-p:
	//	obj( /*...*/ )
	//
	//	p <- obj
	//default:
	//	// No more objects left — retry later or fail
	//	return
	//}
}
