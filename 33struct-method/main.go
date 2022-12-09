package main

import "fmt"

func main() {
	edwin := Person{
		Name: "edwin",
		Age:  22,
	}
	edwin.sayHi()
}

type Person struct {
	Name string
	Age  int
}

func (person Person) sayHi() {
	fmt.Println("Hi", person.Name)
}
