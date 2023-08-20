package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	// "time"
)

func main() {
	// rand.Seed(time.Now().Unix()) // rand seed 는 매번 다른 값을 주지 않으면 결과가 동일함
	// Go 1.20 부터는 Seed 를 호출하지않아도 됨
	i := rand.Intn(15)

	fmt.Println(i)

	if i < 10 {
		fmt.Println("A")
	} else if i == 10{
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
	
	// if condition
	str := "Hello world!"
	// ioutil.WriteFile is deprecated: As of Go 1.16, this function simply calls [os.WriteFile].deprecated(default)
	// if err := os.WriteFile("5_controll.txt", []byte(str), 0644); err != nil {
	if err := ioutil.WriteFile("5_controll.txt", []byte(str), 0644); err != nil {
		fmt.Println(err)
	}

	switch i {
	case 0, 1:
		fmt.Println("A")
	case 2, 3, 4:
		fmt.Println("B")
	case 5:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}