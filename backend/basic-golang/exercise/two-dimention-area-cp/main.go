package main

import (
	"fmt"
)

// Check Point:
// Two Dimention Area
// - Input: Select Choice
// - Output: Result Operation

// Contoh:
// Input:
// 1: Rectange Area
// 2: Rectangular Area
// 3: Triangle Area
// 4: Circle Area
// - Enter choice 1, 2, 3, or 4: 1 | 2 | 3 | 4
//   - (1) Rectange Area
//   	- Masukkan sisi : 5
//   - (2) Rectangular Area
// 		- Masukkan panjang : 5
// 		- Masukkan lebar : 10
// 	 - (3) Triangle Area
// 		- Masukkan panjang alas segitiga: 5
// 		- Masukkan tinggi segitiga: 10
// 	 - (4) Circle Area
//      - Masukkan jari-jari : 4

// Output:
// - (1) Luas Persegi adalah : 25
// - (2) Luas Persegi Panjang adalah : 50
// - (3) Luas Segitiga adalah : 25
// - (4) Luas Lingkaran adalah : 50.240000

func main() {
	// TODO: answer here
	var panjang float32
	var lebar float32
	var tinggi float32
	var jari float32
	const phi = 3.14

	var choice int = 0
	var result float32 = 0
	fmt.Println("1: Rectangle Area")
	fmt.Println("2: Rectangle Area")
	fmt.Println("3: Triangle Area")
	fmt.Println("4: Circle Area")
	fmt.Print("Enter Choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Printf("Masukkan sisi: ")
		fmt.Scan(&panjang)
		result = panjang * panjang
		fmt.Printf("Result: %d", result)
		// fmt.Printf("Result: %d", result)
	case 2:
		fmt.Printf("Masukkan Panjang: ")
		fmt.Scan(&panjang)
		fmt.Printf("Masukkan Lebar: ")
		fmt.Scan(&lebar)
		result = panjang * lebar
		fmt.Println("Result: ", result)
	case 3:
		fmt.Printf("Masukkan Panjang: ")
		fmt.Scan(&panjang)
		fmt.Printf("Masukkan Tinggi: ")
		fmt.Scan(&tinggi)
		result = (panjang * tinggi) / 2
		fmt.Println("Result:", result)
	case 4:
		fmt.Printf("Masukkan Jari-jari")
		fmt.Scan(&jari)
		result = phi * jari * jari
		fmt.Println("Result:", result)
	default:
		fmt.Println("Invalid choice")
	}
}
