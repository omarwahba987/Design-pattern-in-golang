package main
// https://www.geeksforgeeks.org/template-method-design-pattern/
// https://github.com/yksz/go-design-patterns/blob/master/behavior/template_method.go
//In Template pattern, an abstract class exposes defined way(s)/template(s) to execute its methods.
// Its subclasses can override the method implementation as per need but the invocation is to be in the same way
// as defined by an abstract class. This pattern comes under behavior pattern category.
import (
	"fmt"
)

type AbstractClass struct {
	template template
}

func (c *AbstractClass) TemplateMethod() {
	c.template.method1()
	c.template.method2()
}

type template interface {
	method1()
	method2()
}

type ConcreteClass struct{}

func (c *ConcreteClass) method1() {
	fmt.Println("ConcreteClass.method1()")
}

func (c *ConcreteClass) method2() {
	fmt.Println("ConcreteClass.method2()")
}

type ConcreteClass2 struct{}

func (c *ConcreteClass2) method1() {
	fmt.Println("ConcreteClass2.method1()")
}

func (c *ConcreteClass2) method2() {
	fmt.Println("ConcreteClass2.method2()")
}
func main() {
	abstractClass := AbstractClass{new(ConcreteClass)}
	abstractClass2 := AbstractClass{new(ConcreteClass2)}
	abstractClass.TemplateMethod()
	abstractClass2.TemplateMethod()
}