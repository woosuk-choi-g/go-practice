package main

import (
	"log"
	"os"
)

type error interface {
	Error() string
}

func main() {
	f, err := os.Open("C:\\temp\\1.txt")
	if err != nil {
		log.Printf("%#v", f)
		log.Printf("%#v", err)
		log.Fatal(err.Error()) // 2023/08/20 04:19:02 open C:\temp\1.txt: no such file or directory
	}
	println(f.Name())

	// _, err := otherFunc()
	// switch err.(type) {
	// default:
	// 	println("ok")
	// case MyError:
	// 	log.Print("Log my error")
	// case error:
	// 	log.Faltal(err.Error())
	// }
}
