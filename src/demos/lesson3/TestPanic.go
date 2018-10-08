package main

import "fmt"

func main() {
	afunc()
	fmt.Println("call:main")
}

func afunc() {
	defer cfunc()
	defer bfunc()
	panic("error：call:afunc")
	fmt.Println("recovery")
}

func bfunc() {
	msg := recover()
	fmt.Println(msg)
}

func cfunc() {
	fmt.Println("call:cfunc")
}
