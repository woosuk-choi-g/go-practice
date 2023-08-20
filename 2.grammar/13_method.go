package main

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.width * r.height
}

// 포인터 Receiver
func (r *Rect) aread2() int {
	r.width++
	return r.width * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()
	println(area)

	area2 := rect.aread2()
	println(rect.width, area2)
}