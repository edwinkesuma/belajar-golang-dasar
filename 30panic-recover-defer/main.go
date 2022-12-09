package main

import "fmt"

func main() {
	runApp(false)
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan pesan ", message)
	}
	fmt.Println("Aplikasi selesai.")
}
