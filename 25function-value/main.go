package main

import "fmt"

func main() {
	goodBye := sayGoodBye
	fmt.Println(goodBye("edwin"))
}

func sayGoodBye(name string) string {
	return "Good bye " + name
}
