package main

import "fmt"

func main() {
	var mapNameAge map[string]int = make(map[string]int)
	mapNameAge["aaa"] = 20
	mapNameAge["bbb"] = 30

	delete(mapNameAge, "aaa")
	for name, age := range mapNameAge {
		fmt.Println(name)
		fmt.Println(age)
	}

}
