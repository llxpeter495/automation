package main

import "fmt"

func add(x int) (sum int) {
	if x == 1 {
		sum = x
		return
	} else {
		sum = x + add(x-1)
		return
	}

}
func feibo(n int) (num int) {
	if n == 1 || n == 2 {
		return 1
	} else {
		num = feibo(n-1) + feibo(n-2)
		return num
	}
}
func feibo2(n int) (arr []int) {
	arr = append(arr, 0)
	for i := 1; i <= n; i++ {
		if i == 1 || i == 2 {
			arr = append(arr, 1)

		} else {
			arr = append(arr, arr[i-1]+arr[i-2])
		}

	}
	return
}

func test(a int, b ...any) {

	fmt.Println(a, b)

}

func main() {

}
