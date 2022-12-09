package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	mySlice := []string{"edwin", "kesuma"}
	for _, value := range mySlice {
		fmt.Println(value)
	}

	myMap := map[string]string{
		"firstName": "edwin",
		"lastName":  "kesuma",
	}

	for key, value := range myMap {
		fmt.Printf("The value of %v is %v\n", key, value)
	}
}
