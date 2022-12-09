package main

import "fmt"

func main() {
	edwin := Person{
		Name:   "edwin",
		Addres: "Surabya",
		Age:    22,
	}

	var kesuma Person
	kesuma.Name = "kesuma"
	kesuma.Addres = "Madiun"
	kesuma.Age = 22

	fmt.Println(edwin)
	fmt.Println(kesuma)

	fmt.Println(edwin.Name)
	fmt.Println(edwin.Addres)
	fmt.Println(edwin.Age)

	fmt.Println(kesuma.Name)
	fmt.Println(kesuma.Addres)
	fmt.Println(kesuma.Age)
}

type Person struct {
	Name, Addres string
	Age          int
}
