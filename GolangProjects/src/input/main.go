package main

import (
	"io"
	"os"
)

func main() {
	mystring := ""
	arg := os.Args
	if len(arg) == 1 {
		mystring = "Please give me an argument"
	} else {
		mystring = arg[1]
	}
	io.WriteString(os.Stdout, mystring)
	io.WriteString(os.Stdout, "\n")

}
