package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadFile(url, FileName string) {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func() { _ = r.Body.Close() }()
	f, err := os.Create(FileName)
	defer func() { _ = f.Close() }()
	n, err := io.Copy(f, r.Body)
	fmt.Println(n, err)
}

func main() {

}
