package main
// https://github.com/yksz/go-design-patterns/blob/master/structure/bridge.go
// structural pattern
// in this pattern decouples implementation class and abstract class by providing a bridge structure between them
// This pattern involves an interface which acts as a bridge which makes the functionality of concrete classes
// independent from interface implementer classes. Both types of classes can be altered structurally without affecting each other

import (
	"fmt"
)

type DrawingAPI interface {
	DrawCircle(x, y, radius float32)
}

type DrawingAPI1 struct{}

func (a *DrawingAPI1) DrawCircle(x, y, radius float32) {
	fmt.Printf("API 1.circle at %f:%f radius %f\n", x, y, radius)
}

type DrawingAPI2 struct{}

func (a *DrawingAPI2) DrawCircle(x, y, radius float32) {
	fmt.Printf("API 2.circle at %f:%f radius %f\n", x, y, radius)
}
/////////////////////////////////////////////////////////////////////
type Shape interface {
	Draw()
	ResizeByPercentage(pct float32)
}

type CircleShape struct {
	x, y, radius float32
	drawingAPI   DrawingAPI
}

func (s *CircleShape) Draw() {
	s.drawingAPI.DrawCircle(s.x, s.y, s.radius)
}

func (s *CircleShape) ResizeByPercentage(pct float32) {
	s.radius *= pct
}

func main() {
	shapes := []Shape{&CircleShape{1, 2, 3, new(DrawingAPI1)}, &CircleShape{5, 7, 11, new(DrawingAPI2)}}
	for _, shape := range shapes {
		shape.ResizeByPercentage(2.5)
		shape.Draw()
	}
	sh := Shape(&CircleShape{15,16,20,new(DrawingAPI1)})
	sh.Draw()
	sh.ResizeByPercentage(2)
	sh.Draw()
}
