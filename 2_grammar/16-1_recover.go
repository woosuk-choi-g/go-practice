package main

import (
	"fmt"
	"os"
)

func main() {
	openFile("Invalid.txt") // OPEN ERROR open Invalid.txt: no such file or directory
}

func openFile(fn string) {
	// defer 함수. panic 호출시 실행됨
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
}
