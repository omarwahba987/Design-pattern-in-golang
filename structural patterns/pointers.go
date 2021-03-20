package main

import (
	"fmt"
	"runtime"
)

type strct struct{
	point 	*string
	str 	string
}

func (c *strct) print(s *string) {
	c.str="xx"
	fmt.Println("c.point>",c.point)
	fmt.Println("c.str>",c.str)
	fmt.Println("s>",s)
}

type strct2 struct {
	pointONStruct strct
	ONStruct	strct
}
func main()  {

	obj := new(strct)
	fmt.Println("1",obj,"--",obj.str,"--",obj.point)
	str:="pp"
	obj2:= &strct{&str,"ss"}
	obj2.print(&str)
	fmt.Println("2",obj2,"--",obj2.str,"--",*obj2.point)
	obj3:=strct{}
	fmt.Println("3",obj3,"--",obj3.str,"--",obj3.point)
	obj4:=new(strct2)
	fmt.Println("4",obj4)
	obj5:=&strct2{}
	fmt.Println("5",obj5)

	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCPU())



}