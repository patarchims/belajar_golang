package main

import "fmt"

// Tipe data slice
// Adalah potongan dari Array
// slice mirip dengan Array, yang membedakan adalah
// ukuran slice bisa berubah
// Slice dan array selalu terkoneksi, dimana slice adalah data
// yang mengakses sebagaian dari data di Array

// DETAIL TIPE DATA SLICE
// Tipe data slice memiliki tipe data
// Yaitu pointer, lenght dan capacity
// Pointer adalah petunjuk data pertama di array para slice
// length adalah panjang dari slice, dan
// Capacity adalah kapasitas dimana length tidak
// boleh lebih dari capacity

func main() {
	var months = [...]string{
		"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	var slice1 = months[4:7]

	fmt.Println(slice1)
	months[5] = "Diubah"
	fmt.Println(slice1)
}
