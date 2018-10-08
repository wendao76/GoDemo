package main

import "fmt"

func main() {
	var add = func(a int, b int) int {
		return a + b
	}
	var add2 = createAddFunc()

	fmt.Println(add(5, 10))
	fmt.Println(add2(5, 10))
	fmt.Println(testMultiParam(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func createAddFunc() func(a int, b int) int {
	return func(a int, b int) int {
		return a + b
	}
}

//测试函数参数
func testFuncParam(a int, funca func(a int, b int)) int {
	return 1
}

//测试不定参数
func testMultiParam(a ...int) int {
	var returnVal int
	for _, v := range a {
		returnVal += v
	}
	return returnVal
}
