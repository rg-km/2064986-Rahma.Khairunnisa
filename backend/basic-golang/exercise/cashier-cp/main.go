package main

import (
	"fmt"
)

// Check Point:
// Cashier
// - Input: Nama Menu, Add Again
// - Output: Total Harga yang harus dibayar

// Contoh:
// Menu makanan:
// 1 . Menu: bakso ,  Price : 20000
// 2 . Menu: burger ,  Price : 15000
// 3 . Menu: sate ,  Price : 25000
// 4 . Menu: sosis ,  Price : 20000
// 5 . Menu: soto ,  Price : 25000

// Input:
// - Form:
//   - Masukan nama menu yang mau dipesan: bakso

// Output:
// Menu telah ditambahkan ke Ordered Menu:
// Menu: bakso ,  Price : 20000

// Input:
// - Ingin memesan menu lain?(yes/no): if no (break) --> if yes (add again)

// Output:
// Data Order Updated:
//  - Menu: sate ,  Price : 25000
//  - Menu: burger ,  Price : 15000
//  - Menu: sosis ,  Price : 20000
//  - Total harga makanan yang harus anda bayar:  60000

func main() {
	foodMenu := map[string]int64{
		"bakso":  20000,
		"burger": 15000,
		"sate":   25000,
		"sosis":  20000,
		"soto":   25000,
	}
	var pilihan, addAgain string
	orderMenu := make(map[string]int64)

	// TODO: answer here
	
	for {
		fmt.Println("Menu Makanan : ")
		for k, v := range foodMenu {
			fmt.Printf("Menu : %s, Price : %s \n ", k, v)
		}

		fmt.Print("Masukkan Nama Menu Makanan yang Ingin Dipesan : ")
		fmt.Scan(&pilihan)

		orderMenu[pilihan] = foodMenu[pilihan]

		fmt.Println()
		fmt.Println("Data Ordered Menu Updated : ")
		fmt.Printf("Menu : %s, Price : %d \n", pilihan, foodMenu[pilihan])

		fmt.Printf("Ingin menambah data baru?(yes/no): ")
		fmt.Scan(&addAgain)

		if addAgain == "no" {
			break
		}
	}
	var total int64
	fmt.Println("Data Ordered Menu Updated!:")
	for k, v := range orderMenu {
		fmt.Println("- Menu : ", k, ", ", "Price :", v)
		total = total + v
	}
	fmt.Printf("Total harga makanan yang harus anda bayar :  %d", total)
}
