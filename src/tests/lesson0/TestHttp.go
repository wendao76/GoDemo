package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("go lang test")
	var result, err = http.Get("http://www.baidu.com")
	fmt.Print(err)
	fmt.Print(result.Header.Get("X-Ua-Compatible"))
	fmt.Print(120)
}
