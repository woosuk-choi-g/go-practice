package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // rand seed 는 매번 다른 값을 주지 않으면 결과가 동일함
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
	if err := ioutil.WriteFile("5_controll.txt", []byte(str), 0644); err != nil {
		fmt.Println(err)
	}
}