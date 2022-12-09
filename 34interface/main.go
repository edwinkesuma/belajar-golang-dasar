package main

import "fmt"

func main() {
	edwin := Person{
		Name: "edwin",
	}
	sayHello(edwin)
}

type HasName interface {
	getName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.getName())
}

type Person struct {
	Name string
}

//method struct harus sama seperti yang ada interface
func (person Person) getName() string {
	return person.Name
}
