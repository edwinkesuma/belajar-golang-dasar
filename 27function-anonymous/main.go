package main

import "fmt"

func main() {
	blaclist := func(name string) bool {
		return name == "admin"
	}

	greet("admin", blaclist)

}

func greet(name string, function func(string) bool) {
	if function(name) {
		fmt.Println("You're blocked")
	} else {
		fmt.Println("Hi, ", name)
	}
}
