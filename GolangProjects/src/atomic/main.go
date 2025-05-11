package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	x  int64
	wg sync.WaitGroup
)

func add() {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
