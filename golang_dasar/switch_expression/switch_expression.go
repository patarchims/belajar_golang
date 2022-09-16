package main

import "fmt"

// Selain if, untuk melakukan percabangan
// Kita juga bisa menggunakan Switch Expression
// Biasanya switch expression digunakan untuk
// Melakukan pengecekan ke kondisi dalam satu variable

func main() {
	name := "Patar"

	switch name {
	case "patar":
		fmt.Println("Hello Patar")
	case "Patar":
		fmt.Println("Hello Patar Chims")
	default:
		fmt.Println("Boleh Kenalan")
	}
}
