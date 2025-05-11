package main

import (
	"fmt"
	"sync"
)

var (
	wg   sync.WaitGroup
	once sync.Once
)

func generate(c chan int) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		c <- i
	}
	close(c)
}

func load(c chan int, c1 chan int) {
	defer wg.Done()
	for {
		x, ok := <-c
		if !ok {
			break
		}
		c1 <- x * x
	}
	f := func() {
		close(c1)
	}
	once.Do(f)

}
func extra(n int) int {
	var count int

	for n > 0 {
		count += n % 10
		n = n / 10
	}
	return count

}

func main() {
	a := make(chan int, 20)
	b := make(chan int, 20)
	wg.Add(3)
	go generate(a)
	go load(a, b)
	go load(a, b)
	wg.Wait()
	count := len(b)
	fmt.Println(len(b)) ///有缓存的通道长度为未读取的元素个数
	for i := 0; i <= count; i++ {
		v, ok := <-b
		fmt.Println(v, ok, len(b))
	}
	fmt.Println(extra(367234))

}
