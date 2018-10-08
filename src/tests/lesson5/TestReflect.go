package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := People{"wendao", 30, 2}
	typeObj := reflect.TypeOf(p)
	valObj := reflect.ValueOf(p)
	for i := 0; i < valObj.NumField(); i++ {
		fieldName := typeObj.Field(i).Name
		fmt.Println(fieldName + ":=" + valObj.FieldByName(fieldName).String())
	}
}

type People struct {
	name string
	age  int
	sex  int
}
