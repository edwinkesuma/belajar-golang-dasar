package main

import "fmt"

func main() {
	value := newMap("")

	if value != nil {
		fmt.Println("Ada data")
	} else {
		fmt.Println("Tidak ada data")
	}
}

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"Name": name,
		}
	}
}
