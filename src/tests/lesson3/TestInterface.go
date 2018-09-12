package main

import (
	"fmt"
	"strconv"
)

func main() {
	var human Human = Human{"wendao", 29, "F"}
	human.toString()
}

type ToString interface {
	toString() string
}

type Human struct {
	name string
	age  int
	sex  string
}

func (human Human) toString() string {
	returnStr := human.name + ", " + strconv.Itoa(human.age) + ", " + human.sex
	fmt.Println(returnStr)
	return returnStr
}
