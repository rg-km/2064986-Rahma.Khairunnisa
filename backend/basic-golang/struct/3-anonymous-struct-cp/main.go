package main

import "fmt"

func main() {
	//membuat rectangle dengan anonymous struct
	//field dari struct ini sama seperti rectangle sebelumnya
	// TODO: answer here
	newRectangle := struct {
		Width int
		Length int
	}{
		Width: 3,
		Length: 4,
	}

	fmt.Printf("rectangle dengan lebar %d dan panjang %d", newRectangle.Width, newRectangle.Length)
}
