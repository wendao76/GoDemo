package main

import "fmt"

func main() {
	var girl Girl
	girl.speak("hello world")
	girl.Eating("blake")
}

type Person struct {
}

type Animal struct {
}

type eat interface {
	Eating(foodName string) string
}

type eat2 interface {
	Eating(foodName string) string
}

type Girl struct {
	Person
}

func (person *Person) speak(words string) {
	fmt.Println(words)
}

func (animal *Animal) eat(food string) {
	fmt.Println("I eat " + food)
}

func (person *Person) Eating(food string) string {
	fmt.Println("I eat " + food)
	return food
}

func (person *Person) eat(food string) {
	fmt.Println("I eat " + food)
}
