package main

import "fmt"

func main() {
	firstName, lastName := getName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}

func getName() (firstName, lastName string) {
	firstName = "edwin"
	lastName = "kesuma"
	return
}
