package main

import "fmt"

type person struct {
	name string
	age int
}

type dict struct {
	data map[int]string
}

func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d
}

func main() {
	p := person{}
	p.name = "Lee"
	p.age = 10
	fmt.Println(p)

	var p1 person
	p1 = person{"Bob", 20}
	p2 := person{name: "Sean", age: 50}
	fmt.Println(p1)
	fmt.Println(p2)

	p3 := new(person) // 포인터를 리턴
	p3.name = "Lee" // 포인터도 . 을 이용해 접근 가능
	fmt.Println(p3)

	dic := newDict()
	dic.data[1] = "A"
	fmt.Println(dic)
}