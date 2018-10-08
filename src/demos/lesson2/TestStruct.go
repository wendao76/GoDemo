package main

import (
	"fmt"
	"strconv"
)

func main() {
	var book Book
	book.title = "我自己的书"
	book.prize = 25
	book.printInfo()
	(&book).printStrInfo()
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
