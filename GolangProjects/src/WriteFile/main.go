package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(".\\src\\WriteFile\\Test2.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
	//直接写文件
	file.WriteString("Hello World\n")
	//缓冲写
	writer := bufio.NewWriter(file)
	writer.WriteString("Hello NewWorld\n")
	writer.WriteString("Hello NewWorld\n")
	writer.Flush()
	//一次性写入整个文件
	str := "This is a test"
	str2 := "This is another test"
	str3 := "This is not a exercise"
	os.WriteFile(".\\src\\WriteFile\\Test3.txt", []byte(str), 0666)
	os.WriteFile(".\\src\\WriteFile\\Test3.txt", []byte(str2), 0666)
	os.WriteFile(".\\src\\WriteFile\\Test3.txt", []byte(str3), 0666)
}
