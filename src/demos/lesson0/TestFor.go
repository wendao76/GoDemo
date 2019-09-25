package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i = 1
	for i=1; i< 100; i++{
		fmt.Println(i)
	}
	fmt.Println("break")
agc:
	for true{
		fmt.Println("hello world")
		break agc
	}

	var testStr = "this is my name"
	for idx,val := range testStr{
		fmt.Println(strconv.Itoa(idx) + ":" +(string(val)))
	}

	var str = fmt.Sprintln("a" + "b")
	fmt.Println(str)

}

