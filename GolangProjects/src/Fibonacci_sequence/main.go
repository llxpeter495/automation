package main

import (
	"fmt"
	"sync"
)

var count int
var count2 int64
var fei = map[int64]int64{1: 1, 2: 2}

func feibo(i int64) (result int64) {
	count++
	_, ok := fei[i]
	if ok {
		return fei[i]
	} else {
		fei[i] = feibo(i-1) + feibo(i-2)

	}

	return fei[i]
}
func feibo2(i int64) (result int64) {
	temp := []int64{1, 2}
	var n int64
	if i <= 2 {
		return i
	} else {
		for n = 3; n <= i; n++ {
			temp = append(temp, temp[n-2]+temp[n-3])
			result = temp[n-1]
		}
	}

	return result
}
func feibo3(i int64) (result int64) {
	count2++
	if i <= 2 {
		result = i
	} else {
		result = feibo3(i-1) + feibo3(i-2)
	}

	return result
}
func f1(x int64, c chan int64) {
	c <- feibo(x)
}
func f2(x int64, c chan int64) {
	c <- feibo2(x)
}
func f3(x int64, c chan int64) {
	c <- feibo3(x)

}

func main() {
	var wg sync.WaitGroup
	var b int64
	c1 := make(chan int64)
	c2 := make(chan int64)
	c3 := make(chan int64)

	wg.Add(3)
	fmt.Println("请输入一个整数:")
	fmt.Scan(&b)
	go func() {
		defer wg.Done()
		f1(b, c1)
	}()
	go func() {
		defer wg.Done()
		f2(b, c2)
	}()
	go func() {
		defer wg.Done()
		f3(b, c3)

	}()
	go func() {
		wg.Wait()
		close(c1)
		close(c2)
		close(c3)

	}()

	fmt.Println(<-c1)
	fmt.Println(<-c2)
	fmt.Println(<-c3)
	fmt.Println(count)  ///必须放在代码末尾，保证所有goroutine结束后再打印递归次数
	fmt.Println(count2) ///必须放在代码末尾，保证所有goroutine结束后再打印递归次数
}
