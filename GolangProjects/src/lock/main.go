package main

import (
	"fmt"
	"sync"
)

var (
	x    int
	wg   sync.WaitGroup
	lock sync.Mutex
)

func add() {
	defer wg.Done()
	for i := 0; i < 50000; i++ {
		lock.Lock() ///保护全局变量一次只被一个goroutine操作
		x = x + 1
		lock.Unlock()
	}

}

func main() {

	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
