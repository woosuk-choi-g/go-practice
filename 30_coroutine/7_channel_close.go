package main

func main() {
	ch := make(chan int, 2)

	ch <- 1
	ch <- 2

	close(ch)

	println(<-ch)
	println(<-ch)

	if _, success := <-ch; !success {
		println("더이상 데이타 없음")
	}

	for i := range ch {
		println(i)
	}
}