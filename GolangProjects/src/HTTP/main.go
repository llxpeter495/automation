package main

import (
	"fmt"
	"io"
	"net/http"
)

func get() {
	request, err := http.NewRequest("GET", "https://ai.ashuiai.com/", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
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

	// 关闭响应体
	response.Body.Close()

	// 打印响应内容
	fmt.Println(string(body))
}

func post() {
	request, err := http.NewRequest("POST", "https://httpbin.org/post", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
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

	// 关闭响应体
	response.Body.Close()

	// 打印响应内容
	fmt.Println(string(body))
}
func put() {
	request, err := http.NewRequest("PUT", "https://httpbin.org/put", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
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

	// 关闭响应体
	response.Body.Close()

	// 打印响应内容
	fmt.Println(string(body))
}
func delete() {
	request, err := http.NewRequest("DELETE", "https://httpbin.org/delete", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
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

	// 关闭响应体
	response.Body.Close()

	// 打印响应内容
	fmt.Println(string(body))
}

func patch() {
	request, err := http.NewRequest("PATCH", "https://httpbin.org/patch", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	request.Header.Add("content-type", "Application/json")
	request.Header.Add("user-agent", "Chrome")
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

	// 关闭响应体
	response.Body.Close()

	// 打印响应内容
	fmt.Println(string(body))
}

func main() {
	patch()
	
}
