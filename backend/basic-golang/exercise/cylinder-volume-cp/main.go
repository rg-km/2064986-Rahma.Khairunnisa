package main

import (
	"fmt"
)

// Check Point:
// Menghitung volume tabung
// - Input: jari-jari, tinggi
// - Output: volume tabung

// Contoh:
// Input:
// - Masukkan jari-jari alas tabung: 2
// - Masukkan tinggi tabung : 20
// Output:
// - Jadi volumenya adalah : 251.200012

func main() {
	// TODO: answer here
	var v, r, t float64
	const phi = 3.14

	fmt.Print("Masukkan Jari-jari alas tabung :")
	fmt.Scanln(&r)
	fmt.Print("Masukkan Tinggi Tabung :")
	fmt.Scanln(&t)

	v = phi * r * r * t

	fmt.Println("Jadi Volumenya adalah : ", v)
}
