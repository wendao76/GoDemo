package main

import (
	"fmt"
	"strconv"
)

func main() {
	var human Human = Human{"wendao", 29, "F"}
	human.changeName("xiaoming")
	human.toString()
	assertType(human.name)
}

type ToString interface {
	toString() string
	changeName(name string)
}

type Human struct {
	name string
	age  int
	sex  string
}

func (human Human) toString() string {
	returnStr := human.name + ", " + strconv.Itoa(human.age) + ", " + human.sex
	return returnStr
}

func (human *Human) changeName(name string) {
	human.name = name
}

func assertType(sType interface{}) {
	switch sType.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", sType.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", sType.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}
