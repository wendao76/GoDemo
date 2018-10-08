package main

import (
	"fmt"
	"strconv"
)

func main() {
	error := NewStringError("error happned", 500)
	error.printError()
}

type StringError struct {
	msg  string
	code int
}

func NewStringError(msg string, code int) *StringError {
	return &StringError{msg, code}
}

func (stringError *StringError) printError() {
	fmt.Println(strconv.Itoa(stringError.code) + " " + stringError.msg)
}
