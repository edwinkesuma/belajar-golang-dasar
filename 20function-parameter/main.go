package main

import "fmt"

func main() {
	sayHi("edwin", "kesuma")
}

func sayHi(firstName string, lastName string) {
	fmt.Println("Hi ", firstName, lastName)
}
