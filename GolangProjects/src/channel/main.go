package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f1(ch chan int) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		ch <- i
	}
	close(ch)
}
func f2(ch1, ch2 chan int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
}

func main() {
	a := make(chan int, 20)
	b := make(chan int, 20)
	wg.Add(2)
	go f1(a)
	go f2(a, b)
	wg.Wait()
	for i := range b {
		fmt.Println(i)
	}

}
