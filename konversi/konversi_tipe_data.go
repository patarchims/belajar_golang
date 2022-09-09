package main

import "fmt"

// Konversi Tipe Data
// Di golang kadang kita butuh melakukan konversi
// tipe data dari satu tipe ke tipe lain
// misalnya kita mengkonversi tipe data int32 ke int64, dan lain lain

func main() {
	var nilai32 int32 = 10000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

}
