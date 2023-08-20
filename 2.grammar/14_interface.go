package main

import "math"

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area()
		println(a)
	}
}

func main() {
	r := Rect{10., 20.}
	c := Circle{10}

	// 명시된 인터페이스를 구현하는 메소드가 없으면 컴파일 에러
	// Circle.perimeter 를 주석처리하면 아래와 같은 에러가 발생
	// ./14_interface.go:42:14: cannot use c (variable of type Circle) as Shape value in argument to showArea: Circle does not implement Shape (missing method perimeter) (exit status 1)
	showArea(r, c)
}