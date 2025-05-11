package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func postform() {

	data := make(url.Values)
	data.Add("name", "kaka")
	data.Add("age", "18")
	payload := data.Encode()
	r, _ := http.Post(
		"https://httpbin.org/post",
		"application/x-www-form-urlencoded",
		strings.NewReader(payload),
	)
	defer r.Body.Close()
	content, _ := io.ReadAll(r.Body)
	fmt.Printf("%s", content)

}

func postjson() {
	user := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{
		Name: "kaka",
		Age:  18,
	}
	payload, _ := json.Marshal(user)
	r, _ := http.Post(
		"https://httpbin.org/post",
		"application/json",
		bytes.NewReader(payload),
	)
	defer r.Body.Close()
	content, _ := io.ReadAll(r.Body)
	fmt.Printf("%s", content)
}

func main() {

	postform()
	postjson()
}
