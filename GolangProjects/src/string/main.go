package main

import (
	"fmt"
	"strings"
)

func main() {

	a := make([]string, 10)
	a = append(a, "hehe")
	fmt.Println(a)
	fmt.Println(len(a))
	var arr = [...]float32{1, 2, 3, 4, 5}

	for i := len(arr); i < 0; i-- {
		fmt.Println(arr[i])
	}
	var letters []int
	letters = append(letters, 12)
	fmt.Println(letters)
	var res string
	res = "hello new world"

	for i := len(res) - 1; i >= 0; i-- {
		fmt.Println(string(res[i]))
	}
	str := "fijf\nfeifie\nqowfj\n"
	fmt.Println(str)
	str = strings.Replace(str, "\n", "", -1)
	fmt.Println(str)
}
