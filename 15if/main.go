package main

import "fmt"

func main() {
	name := "Dwi"
	if name == "edwin" {
		fmt.Println("Hallo edwin!")
	} else if name == "kesuma" {
		fmt.Println("Hallo kesuma")
	} else {
		fmt.Println("Boleh kenalan?")
	}

	if i := 4; i < 5 {
		fmt.Println("i kurang dari 5")
	}
}
