package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func printbody(r *http.Response) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s", content)
}
func requestbypara() {
	request, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	para := make(url.Values)
	para.Add("name", "Tom")
	para.Add("age", "18")
	fmt.Println(para)
	request.URL.RawQuery = para.Encode()

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

func requestbyhead() {
	request, err := http.NewRequest("GET", "https://httpbin.org/get", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("user-agent", "Chrome")
	request.Header.Add("Content-Type", "Application/xml")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

}

func main() {
	requestbypara()
}
