package main

import (
	"fmt"
)

func main() {
	// Membuat sebuah map dengan tipe data string sebagai kunci dan int sebagai nilai
	data := make(map[string]int)

	// Meminta input dari pengguna
	var nama string
	var umur int

	for i := 0; i < 3; i++ {
		fmt.Print("Masukkan nama: ")
		fmt.Scanln(&nama)

		fmt.Print("Masukkan umur: ")
		fmt.Scanln(&umur)

		// Menambahkan data ke dalam map
		data[nama] = umur
	}

	// Mencetak data dari map
	fmt.Println("Data map:", data)
}