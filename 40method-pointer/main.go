package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) getMaried() {
	man.Name = "Mr. " + man.Name
}

func main() {
	edwin := Man{
		Name: "edwin",
	}
	fmt.Println(edwin.Name)

	edwin.getMaried()
	fmt.Println(edwin.Name)
}
