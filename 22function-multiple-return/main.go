package main

import "fmt"

func main() {
	firstname, _ := greet()
	fmt.Println(firstname)
}

func greet() (string, string) {
	return "edwin", "kesuma"
}
