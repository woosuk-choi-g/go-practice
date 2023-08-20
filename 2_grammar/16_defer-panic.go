package main

import "os"

func main() {
	f, err := os.Open("1.txt")
	if err != nil {
		// error 가 발생한 시점에 종료
		panic(err)
	}

	// main 마지막에 파일 close 실행
	defer f.Close()

	bytes := make([]byte, 1024)
	f.Read(bytes)
	println(len(bytes))
}