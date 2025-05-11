package main

import "fmt"



func main() {
	s := []int{}
	s1 := []int{}

	for i := 1; i <= 100; i++ {
		s1 = append(s1, i)
	}
	for i := 2; i <= 100; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				s = append(s, i)
				s1[i-1] = 0
				break

			}

		}

	}
	fmt.Println(s)
	fmt.Println(s1)

	s2 := []int{}
	for i := 0; i < len(s1); i++ {
		if s1[i] != 0 {
			s2 = append(s2, s1[i])
		}
	}
	fmt.Println(s2)
	for i := range 10 { ///1.22版本后新增加的功能
		fmt.Println(i)
	}
}
