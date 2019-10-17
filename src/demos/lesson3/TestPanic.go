package main

import "fmt"

/**
备注：
	defer 是栈式结构， 先进后出
 */
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
	fmt.Println("error-msg")
	fmt.Println(msg)
	fmt.Println("call:bfunc")
}

func cfunc() {
	fmt.Println("call:cfunc")
}
