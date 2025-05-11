package main

import (
	"fmt"
	"maps"
	"strconv"
	"strings"
)

func main() {
	var emp = map[string]string{"name": "sam", "age": "20"}
	emp["major"] = "engineer"
	fmt.Println(emp)
	stu := make(map[string]int8)
	stu["math"] = 20
	fmt.Println(stu)
	s := []string{}
	for i := 0; i < 20; i++ {
		s = append(s, strconv.Itoa(i))
	}
	fmt.Println(s)
	resnew := strings.Join(s, "-")
	fmt.Println(resnew)
	animal := map[string]string{"cat": "yellow", "dog": "green", "elephant": "white"}
	fmt.Println(animal)
	k := []string{}
	for i := range animal {
		k = append(k, i)
	}
	fmt.Println(k)
	inv := map[string]string{}
	for i := 1; i < 255; i++ {
		inv[fmt.Sprint("switch", i)] = fmt.Sprint("10.8.6.", i)
	}
	fmt.Println(inv)
	brr := maps.Keys(inv)
	for i := range brr {
		fmt.Println(i)
	}
}
