package main

import (
	"fmt"
	"time"
)

var count int = 0

func Add(sign chan int) {
	fmt.Println("start Add")
	<-sign
	fmt.Println("recieve msg")
	fmt.Println("end Add")
}

/**
 备注：
	chan<-  写数据
	<-chan 读数据
 */
func main() {
	var sign chan int = make(chan int)
	go Add(sign)
	time.Sleep(5 * time.Second)
	fmt.Println("send msg")
	sign <- 1
	time.Sleep(5 * time.Second)
	fmt.Println("over")
}
