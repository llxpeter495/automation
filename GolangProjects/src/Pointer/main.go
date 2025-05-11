package main

import "fmt"

func selfadd(x *int) int {
	*x++
	return *x
}

func main() {
	var p *int
	a := 10
	p = &a
	selfadd(p)
	fmt.Println(a)

}
