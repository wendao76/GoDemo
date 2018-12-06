package main

import (
	"fmt"
	"strconv"
)

//结构体样例
func main() {
	var book Book
	book.title = "我自己的书"
	book.prize = 25
	book.printInfo()
	(&book).printStrInfo()

	rect1 := new(Rect)

	rect2 := &Rect{}

	rect3 := &Rect{0, 0, 100, 200}

	rect4 := &Rect{width: 100, height: 200}

	rect5 := Rect{1, 2, 3, 4}

	fmt.Println(rect1, rect2, rect3, rect4, rect5)
	fmt.Printf("%v\n%T\n", rect1, rect1)
	fmt.Printf("%v\n%T\n", rect5, rect5)
}

type Book struct {
	title string
	prize int
}

func (book Book) printInfo() {
	fmt.Println(book.title)
	fmt.Println(book.prize)
}

func (book *Book) printStrInfo() {
	fmt.Println(book.title + ", " + strconv.Itoa(book.prize))
}

type Rect struct {
	x, y          float64
	width, height float64
}
