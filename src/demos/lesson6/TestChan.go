package main

import "fmt"

var count int = 0

func Add(sign chan int) {

	for i := 1; i < 10000000000; i = i + 1 {
		count = count + 1
	}
	fmt.Print(count)
	sign <- 1
}

func main() {
	var sign chan int = make(chan int)
	go Add(sign)
	<-sign

}
