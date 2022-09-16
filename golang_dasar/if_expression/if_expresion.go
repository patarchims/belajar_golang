package main

import "fmt"

// IF ADALAH SALAH SATU KUNCI YANG DIGUNAKAN
// UNTUK PERCABANGAN, PERCABANGAN ARTINYA
// KITA MENGEKSEKUSI KODE PROGRAM TERTENTU
// KETIKA SUATU KONDISI TERPENUHI
// HAMPIR SEMUA BAHASA PEMROGRAMAN
// MENDUKUNG IF EXPRESSION
func main() {
	name := ""

	if name == "Patar" {
		fmt.Println("Hello " + name)
	} else {
		fmt.Println("Bolehkenalan ?")
	}
}
