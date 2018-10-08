package main

import "fmt"

func test1() {
	defer test2()
	defer test3()
	defer test4()
	fmt.Println("test1")
}

func test2() {
	fmt.Println("test2")
}

func test3() {
	fmt.Println("test3")
}

func test4() {
	fmt.Println("test4")
}

func main() {
	test1()
}
