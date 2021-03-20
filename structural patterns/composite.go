package main
// https://github.com/yksz/go-design-patterns/blob/master/structure/composite.go
// https://sourcemaking.com/design_patterns/composite
// https://www.geeksforgeeks.org/composite-design-pattern/
// structural pattern
// this pattern is used where we need to treat a group of objects in similar way as a single object
// and composes objects in term of a tree structure to represent part as well as whole hierarchy
import (
	"fmt"
)

type Component interface {
	Operation(int)
}

type File struct {
	name string
}

func (f *File) Operation(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
	fmt.Println("File: " + f.name)
}

type Directory struct {
	components []Component
	name       string
}

func (d *Directory) Operation(depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print(" ")
	}
	fmt.Println("Directory: " + d.name)
	for _, component := range d.components {
		component.Operation(depth + 1)
	}
}

func (d *Directory) len() {
	fmt.Println("len. of component list = ",d.name,len(d.components))
}

func (d *Directory) Add(component Component) {
	d.components = append(d.components, component)
}

func (d *Directory) Remove(component Component) {
	for i, v := range d.components {
		if v == component {

			d.remove(i)
			return
		}
	}
}

func (d *Directory) remove(i int) {
	reslice := append(d.components[:i], d.components[i+1:]...)
	newslice := make([]Component, len(reslice))
	copy(newslice, reslice)
	d.components = newslice
}

func (d *Directory) GetChildren() []Component {
	return d.components
}

func main() {
	root := &Directory{name: "root"}
	root.len()
	usr := &Directory{name: "usr"}
	usr.len()
	local := &Directory{name: "local"}
	local.len()
	home := &Directory{name: "home"}
	home.len()
	user1 := &Directory{name: "user1"}
	user1.len()
	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}

	root.Add(usr)
	root.len()
	usr.Add(local)
	usr.len()
	root.Add(home)
	root.len()
	home.Add(user1)
	home.len()
	user1.Add(file1)
	user1.Add(file2)
	user1.len()
	root.Operation(0)
	root.len()
	home.len()
	root.Remove(home)
	root.Operation(0)
	root.len()
}
