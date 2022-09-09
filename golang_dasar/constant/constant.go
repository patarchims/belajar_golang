package main

import "fmt"

// Constant adalah variable yang nilainya tidak bisa diubah
// setelah pertama kali diberi nilai
// Cara pembuatan constant mirip dengan dengan variable,
// yang membedakan hanya kata kunci yang digunakan adalah const,
// bukan var.
// Saat pembuatan constant, kita wajib langsung menginisialisasikan
// datanya

// Deklarasi mutile constant
// sama seperti variable
func main() {
	const firstname string = "Patar"
	const lastName string = "Chims"

	fmt.Println(firstname + lastName)

	const (
		name string = "My Name"
		last string = "Patar Chims"
	)

	fmt.Println(name + last)
}
