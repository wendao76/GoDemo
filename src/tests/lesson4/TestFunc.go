package main

import "fmt"

func main() {
	var add = func(a int, b int) int {
		return a + b
	}

	fmt.Println(add(5, 10))
}
