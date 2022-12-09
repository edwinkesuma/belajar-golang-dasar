package main

import "fmt"

func main() {
	//cara deklarasi variable di go
	var nama string = "Edwin Kesuma"
	var age = 22
	email := "edwin@kesuma.com"
	var (
		firstName = "edwin"
		lastName  = "kesuma"
	)

	fmt.Println(nama)
	fmt.Println(age)
	fmt.Println(email)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
