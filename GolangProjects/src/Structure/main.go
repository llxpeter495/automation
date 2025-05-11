package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	name   string
	hair   string
	height int16
	age    int8
	course []string
	sex    string
}

func (p *person) eat() {
	fmt.Printf("%s在吃饭", p.name)
}

func main() {
	Mycar := struct {
		Make  string
		Model string
	}{
		Make:  "Tesla",
		Model: "Model 3",
	}
	fmt.Println(Mycar)

	Tom := person{hair: "yellow", height: 178, age: 32, course: []string{"c", "java"}}
	fmt.Println(Tom)
	var Tony person
	Tony.course = make([]string, 3, 3)
	Tony.sex = "Male"
	Tony.name = "Tony"
	res, _ := json.Marshal(Tom.course)
	fmt.Println(string(res))
	var Sam person
	Sam = person{"Sam", "black", 178, 30, []string{"math", "art"}, "male"}
	v, _ := json.Marshal(Sam)
	fmt.Println(string(v))
	Tony.eat()

}
