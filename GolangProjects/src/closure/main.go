package main

import "fmt"

func getfoo(a func() int) func() int {
	n := 0
	a = func() int {
		n++
		return n
	}
	return a
}

func external(a int) int {
	x := func() int {
		y := a * a
		return y
	}()
	return x
}
func main() {
	fmt.Println(external(10))
	var f func(int64) int64
	f = func(i int64) (result int64) {
		if i <= 2 {
			result = i
		} else {
			result = f(i-1) + f(i-2)
		}
		return result
	}

	fmt.Println(f(10))

}
