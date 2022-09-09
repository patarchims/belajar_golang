package main

import "fmt"

// Variable adalah tempat untuk menyimpan data
// Variable digunakan agar kita mengakses data yang sama dimanapun kita mau
// Di Go-lang variable hanya bisa menyimpan tipe data yang sama
// jika ingin menyimpan data yang berbeda beda jenis, kita harus
// membuat beberapa variable
// Untuk membuat variable, kita menggunakan kata kunci var,
// lalu diikuti dengan nama variable dan tipe datanya

func main() {
	var name string

	name = "Patar chims"

	fmt.Println(name)

	// teknik kedua
	nikname := "Patar Chims"

	fmt.Println(nikname)
}
