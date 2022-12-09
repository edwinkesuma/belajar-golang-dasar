package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{
		City:     "Surabaya",
		Province: "Jawa Timur",
		Country:  "",
	}

	changeToIndonesia(&address1)
	fmt.Println(address1)
}
