package main

import "fmt"

func main() {
	var add = func(a int, b int) int {
		return a + b
	}
	var add2 = createAddFunc()

	fmt.Println(add(5, 10))
	fmt.Println(add2(5, 10))
}

func createAddFunc() func(a int, b int) int {
	return func(a int, b int) int {
		return a + b
	}
}
