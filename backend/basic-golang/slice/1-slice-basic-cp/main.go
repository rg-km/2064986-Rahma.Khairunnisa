package main

import "fmt"

// Disini teman teman akan mencoba untuk
// melakukan penambahan data pada slice.
// Buatlah variable slice dengan tipe data string.
// Lalu masukan nama kalian ke dalam slice.
// Expected outout: ["NamaPanggilan", "Nama Akhir"]
// Contoh [Zein Fahrozi]
// Outputkan jawabannya ya pastikan cap dan len nya adalah 2
func main() {
	// TODO: answer here
	slice = append(slice,"Rahma", "Nailah")
	fmt.Println(len(array), cap(array))

	var slice2 = []string{}
	slice2 = append(slice2, "Rahma", "Nailah")
	fmt.Println(slice2)
}
