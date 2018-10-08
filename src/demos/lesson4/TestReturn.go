package main

import (
	"fmt"
	"time"
)

func main() {

	t1 := time.Now().UnixNano()
	for i := 0; i < 100000000; i++ {
		returnT1(36)
	}
	t2 := time.Now().UnixNano()

	for i := 0; i < 100000000; i++ {
		returnT2(36)
	}
	t3 := time.Now().UnixNano()

	fmt.Println(t2 - t1)
	fmt.Println(t3 - t2)
}

func returnT1(intVal int) int {
	if intVal < 10 {
		return 1
	}

	if intVal < 20 {
		return 2
	}

	if intVal < 30 {
		return 3
	}

	if intVal < 40 {
		return 4
	}

	if intVal < 50 {
		return 5
	}

	return 4
}

func returnT2(intVal int) int {
	var returnVal int = 4
	if intVal < 10 {
		returnVal = 1
	}
	if intVal < 20 {
		returnVal = 2
	}
	if intVal < 30 {
		returnVal = 3
	}
	if intVal < 40 {
		returnVal = 4
	}
	if intVal < 50 {
		returnVal = 5
	}

	return returnVal
}
