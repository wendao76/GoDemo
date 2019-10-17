package main

import (
	"fmt"
	"time"
)

func main() {
	var initNum uint32 = 0
	for i:=0;i<100;i++{
		go t1(&initNum)
	}
	time.Sleep(2)
	fmt.Println(initNum)
	fmt.Println("close")
}

func t1(inNum *uint32) {
	*inNum = *inNum + 1
}
