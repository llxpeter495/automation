package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg     sync.WaitGroup
	x      int
	rwlock sync.RWMutex
)

func read() {
	defer wg.Done()
	rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond * 900)
	rwlock.RUnlock()
}
func write() {
	defer wg.Done()
	rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 900)
	rwlock.Unlock()
}
func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go write()
	}
	time.Sleep(time.Second * 2)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
}
