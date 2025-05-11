package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	person := map[string]string{"name": "Peter", "age": "30"}
	man, _ := json.Marshal(person)
	file, _ := os.OpenFile(".\\src\\JSON\\test.json", os.O_WRONLY|os.O_APPEND, 0666)

	file.WriteString(string(man) + "\n")

	f, _ := os.Open(".\\src\\JSON\\test.json")
	student := map[string]string{}
	decodeer := json.NewDecoder(f)
	err := decodeer.Decode(&student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(student)
	data := map[string]interface{}{}
	newfile, _ := os.ReadFile("D:\\GolangProjects\\src\\JSON\\test2.json")
	if err := json.Unmarshal(newfile, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)
	fmt.Println(data["hobby"])
	jdata, _ := json.MarshalIndent(data, " ", "")
	fmt.Println(string(jdata))

}
