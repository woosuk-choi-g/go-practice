package main

import "fmt"

func main() {
	// c := make(chan int)
	// c <- 1 // fatal error: all goroutines are asleep - deadlock!  // 수신 루틴이 없으므로 데드락
	// fmt.Println(<-c) // 별도의 Go루틴이 없기 때문에 데드락

	ch := make(chan int, 1) // 수신자가 없더라도 보낼 수 있다.
	ch <- 101
	fmt.Println(<-ch)
}
