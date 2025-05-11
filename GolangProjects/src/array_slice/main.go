package main

import (
	"fmt"
	"reflect"
	"slices"
)

func main() {

	s := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(s)
	fmt.Println(s[4:5])
	var a float32
	a = 10.54
	b := fmt.Sprint(a)
	fmt.Println(reflect.TypeOf(b))
	name := "Tom"
	age := 22
	res := fmt.Sprint(name, " is ", age, " years old")
	fmt.Println(res)
	s1 := []int{1, 4, 5}
	s1 = slices.Insert(s1, 1, 100, 80)
	fmt.Println(s1)

	slices.Reverse(s1)
	fmt.Println(s1)
	slices.Sort(s1)
	fmt.Println(s1)
}
