package main

import (
	"fmt"
	"time"
)

func main() {
	go t1()
	fmt.Println(120)
	time.Sleep(1)
	fmt.Println("main")
}

func t1() {
	fmt.Println("t1")
	time.Sleep(500)
}
