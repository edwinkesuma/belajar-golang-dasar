package main

import "fmt"

func main() {
	greet("anjing", filter)
}

func greet(name string, function func(string) string) {
	filtered := function(name)
	fmt.Println("Hi ", filtered)
}

func filter(name string) string {
	if name == "anjing" {
		return "******"
	} else {
		return name
	}
}
