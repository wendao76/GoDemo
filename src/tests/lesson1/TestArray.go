package main

import "fmt"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	arr1Slice := arr1[2:4]

	arr1Slice[0] = arr1Slice[0] + 7
	arr1Slice[1] = arr1Slice[1] + 7
	for _, v := range arr1Slice {
		fmt.Println(v)
	}

	for i := range arr1 {
		fmt.Println(arr1[i])
	}

}
