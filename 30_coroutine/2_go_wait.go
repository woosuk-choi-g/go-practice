package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	i  int64
	mu sync.Mutex // 공유 데이터 i를 보호하기 위한 뮤텍스
}

func (c *counter) increment() {
	c.mu.Lock() // i 값을 변경하는 부분(임계 영역)을 뮤텍스로 잠금
	c.i += 1
	c.mu.Unlock() // i 값을 변경 완료 후 뮤텍스 잠금 해제
}

func (c *counter) display() {
	fmt.Println(c.i)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // 모든 CPU 사용
	c := counter{i: 0}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.increment()
		}()
	}

	wg.Wait()
	c.display()
}
