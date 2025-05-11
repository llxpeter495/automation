package main

import "fmt"

type man struct {
	age     uint8
	height  uint16
	weight  uint16
	name    string
	company employee
	health  uint8
}
type employee struct {
	salary     uint16
	title      string
	department string
	hired      bool
}

func (m *man) run() {
	fmt.Printf("%s is running\n", m.name)
}
func (m *man) jump() {
	fmt.Println("Someone is jumping")
}
func (m *man) attacked() {
	fmt.Printf("%s被攻击了\n", m.name)
	m.health = m.health - 10

}
func (e *employee) fired() {
	fmt.Println("You are fired")
}

type action interface {
	jump()
	run()
}

var count int
var count2 int
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
	if n <= 2 {
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
func main() {

	//for i := 0; i <= 100; i++ {
	//	fmt.Println(i)
	//}
	//str := "golang is the best programming language"
	//for i, v := range str {
	//	fmt.Println(i, string(v))
	//}

	//start := time.Now()
	//fmt.Println(feibo(45))
	//end := time.Now()
	//dur := end.Sub(start)
	//start2 := time.Now()
	//fmt.Println(feibo3(45))
	//end2 := time.Now()
	//dur2 := end2.Sub(start2)
	//fmt.Println(dur.Nanoseconds())
	//fmt.Println(dur2.Nanoseconds())

	//var s = []int{}
	//fmt.Println(s)
	//s = append(s, 1)
	//fmt.Println(s)
	//var ss = []interface{}{}
	//ss = append(ss, "dddwdwdw", 123234, false)
	//fmt.Println(ss)

	var Tom, Tony man

	Tom.age = 20
	Tom.weight = 60
	Tom.height = 180
	Tom.name = "Tom"
	fmt.Println(Tom)
	Tony.name = "Tony"
	Tony.health = 100
	Tony.company.fired()
	Tony.run()
	Tony.attacked()
	fmt.Println(Tony.health)
	var Sam action
	Sam = &man{20, 30, 30, "Sam", employee{department: " IT"}, 100}
	Sam.jump()
	Sam.run()

	var p int
	aa := &p
	*aa = 10
	fmt.Println(p)

	var test = func() {
		fmt.Println("hello world")
	}
	test()
	func() {
		fmt.Println("this is a function")
	}()

}
