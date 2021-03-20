package main

// https://github.com/tmrts/go-patterns/blob/master/structural/proxy.md
// structural pattern
// The proxy pattern provides an object that controls access to another object, intercepting all calls.
// here from function ObjDo i can control witch action can be allow do done



import "fmt"

// To use proxy and to object they must implement same methods
type IObject interface {
	ObjDo(action string)
	ObjDont(action string)
}

// Object represents real objects which proxy will delegate data
type Object struct {
	action string
}

// ObjDo implements IObject interface and handel's all logic
func (obj *Object) ObjDo(action string) {
	// Action behavior
	fmt.Println("I can, ", action)
}
func (obj *Object) ObjDont(action string) {
	// Action denied
	fmt.Println("I can not, ", action)
}

// ProxyObject represents proxy object with intercepts actions
type ProxyObject struct {
	object *Object
}

// ObjDo are implemented IObject and intercept action before send in real Object
func (p *ProxyObject) ObjDoo(action string) {
	//if p.object == nil {
	//	p.object = new(Object)
	//}
	if action == "Walk" {
		p.object.ObjDo(action)
	}else {
		p.object.ObjDont(action)
	}
}

func main() {
	subject := new(ProxyObject)
	subject.ObjDoo("Walk")
	subject.ObjDoo("Run")
}