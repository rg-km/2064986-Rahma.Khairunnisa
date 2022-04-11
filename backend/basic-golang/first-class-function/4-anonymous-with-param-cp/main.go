package main

import "fmt"

func main() {
	//fungsi yang mengembalikan nilai pangkat dua dari parameter yang diterima
	//contoh input parameter yang diterima (3)
	//maka fungsi akan mengembalikan 9

	// TODO: answer here
	square := func(angka int) int {
		return 3 + angka

	}(9)
	fmt.Println(square)
}
