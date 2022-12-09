package main

import "fmt"

func main() {
	myMap := map[string]string{
		"firstName": "edwin",
		"lastName":  "kesuma",
	}

	fmt.Println(myMap)
	//bisa digunakan untuk menambah key value di map
	myMap["email"] = "edwin@test.com"
	fmt.Println(myMap)
	fmt.Println(len(myMap))
	fmt.Println(myMap["firstName"])

	map2 := make(map[string]string)
	map2["first"] = "coba"
	map2["last"] = "lagi"
	fmt.Println(map2)
	fmt.Println(len(map2))

	delete(map2, "last")
	fmt.Println(map2)
}
