package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("D:\\GolangProjects\\src\\FileRead_go\\Test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	//按行读文件
	reader := bufio.NewReader(file)
	//fmt.Println(reader.ReadLine())
	lines := []string{}
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimRight(line, "\r\n")
		lines = append(lines, line)
	}
	fmt.Println(lines)
	//另外一种方法逐行读文件
	newfile, err := os.Open("D:\\GolangProjects\\src\\FileRead_go\\Test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer newfile.Close()
	newlines := []string{}
	scanner := bufio.NewScanner(newfile)
	for scanner.Scan() {
		newline := strings.TrimRight(scanner.Text(), "\n")
		newlines = append(newlines, newline)
	}
	fmt.Println(newlines)
	//整个文件一起读
	data, err := os.ReadFile(".\\src\\FileRead_go\\Test.txt")
	if err != nil {
		fmt.Println("Failed to open file")
	}
	datanew := string(data)
	fmt.Println(strings.Split(datanew, "\r\n"))

}
