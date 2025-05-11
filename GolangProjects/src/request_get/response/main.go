package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io"
	"net/http"
)

func responsebody(r *http.Response) {
	content, _ := io.ReadAll(r.Body)
	fmt.Printf("%s", content)
}

func status(r *http.Response) {
	fmt.Println(r.StatusCode) // http status code
	fmt.Println(r.Status)     // description of http response status

}
func header(r *http.Response) {
	fmt.Println(r.Header.Get("Content-Type"))

}

func encoding(r *http.Response) {
	bufReader := bufio.NewReader(r.Body)
	bytes, _ := bufReader.Peek(1024)
	e, _, _ := charset.DetermineEncoding(bytes, "")
	fmt.Println(e)
	bodyReader := transform.NewReader(bufReader, e.NewDecoder())
	content, _ := io.ReadAll(bodyReader)
	fmt.Printf("%s", content)

}

func main() {
	request, err := http.NewRequest("GET", "https://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, _ := http.Get("https://www.baidu.com")
	responsebody(response)
	status(response)
	header(response)
	encoding(r)
	encoding(response)
	fmt.Println(response.Body)
	fmt.Println(r.Body)
	// 关闭响应体
	response.Body.Close()
}
