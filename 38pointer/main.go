package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	address2 := &address1
	address3 := address1

	address2.City = "Madiun"
	address3.City = "Gresik"

	fmt.Println(address1)
	fmt.Println(*address2)
	fmt.Println(address3)

}
