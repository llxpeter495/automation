package main

import (
	"fmt"
	"sync"
)

func main() {
	a := sync.Map{}
	a.Store("Sam", 100)
	fmt.Println(a)

}
