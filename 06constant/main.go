package main

import "fmt"

func main() {
	//deklarasi constant
	//constant tidak harus dipakai seperti variable biasa
	const age int = 22
	const (
		firstName = "edwin"
		lastName  = "kesuma"
	)

	fmt.Println(age)
	fmt.Println(firstName)
	fmt.Println(lastName)

}
