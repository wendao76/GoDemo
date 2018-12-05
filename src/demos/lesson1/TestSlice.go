package main

import (
	"fmt"
	"strconv"
)

func main() {
	var arr [20]int
	arr = [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	var arr1 = arr[:]
	var slice1 = arr1[:]
	slice12 := append(slice1, 21)
	fmt.Println(slice12)
	fmt.Println(cap(slice12))

	var mapData map[int]string = make(map[int]string)
	mapData[1] = "xiaoming"
	mapData[2] = "xiaoli"
	mapData[3] = "xiaowang"
	for value := range mapData {
		fmt.Println(mapData[value])
	}

	var book Book
	book.title = "三只小猪"
	book.prize = 25
	book.toString()

	var emptySlice = make([]int, 7, 5)
	fmt.Println(emptySlice)

}

type Book struct {
	title string
	prize int
}

func (book *Book) toString() {
	fmt.Println("title=" + book.title + "; prize=" + strconv.Itoa(book.prize))
}
