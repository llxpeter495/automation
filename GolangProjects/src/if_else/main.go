package main

import (
	"fmt"
)

func main() {
	var a int8
	fmt.Println("请输入成绩")
	fmt.Scan(&a)
	if a > 0 && a < 100 {
		switch {
		case a < 90:
			fmt.Println("a小于90")
		case a < 80:
			fmt.Println("a小于80")
		case a < 70:
			fmt.Println("a小于70")
		case a < 60:
			fmt.Println("a小于60")
		default:
			fmt.Println("不及格")
		}
	} else {
		fmt.Println("非法输入")
	}

}
