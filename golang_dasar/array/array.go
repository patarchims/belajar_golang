package main

import "fmt"

// Tipe data Array
// Array adalah tipe data yang berisikan kumpulan data
// dengan tipe data yang sama
// Saat membuat array, kita perlu menentukan jumlah data yang
// ditampung oleh Array tersebut
// Daya tampung Array tidak bisa bertambah setelah Array dibuat

func main() {
	var name [3]string
	name[0] = "Patar "
	name[1] = "Eko"
	name[2] = "Indra"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	// Di golang bisa membuat Array langsung
	// saat deklarasi variable
	var values = [3]int{
		90, 232, 434,
	}
	fmt.Println(values)

	// Function Array
	// len => Untuk mendapatkan panjang Array
	// array[index] => untuk mendapatkan poisisi index
	// array[index] = value => untuk mengubah data di posisi index

	fmt.Println(len(values))

}
