package main

import "fmt"

// Type Declaration
// Type declaration adalah kemampuan membuat tipe data baru
// dari tipe data yang sudah ada
// Type Declaration digunakan untuk membuat alias terhadap
// tipe data yang sudah ada dengan tujuan agar lebih mudah dimengerti

func main() {
	type NoKTP string
	var ktpPatar NoKTP = "123131313"

	fmt.Println(ktpPatar)
}
