package main

import "fmt"

func main() {
	/*
		Reverse Word:
		Example: halo -> olah
	*/
	word := "halo"
	res := ReverseWord(word)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := ReverseWordCorrect(arr)
	// fmt.Println(resCorrect)
}

func ReverseWord(word string) string {
	n := len(word)
	temp := []byte(word)

	for i := 0; i < n/2; i++ {
		temp[i], temp[n-i-1] = temp[n-i-1], temp[i]
	}

	return string(temp)
}

func ReverseWordCorrect(word string) string {
	//return 0 // TODO: replace this
	n := len(word)
	temp := []byte(word)

	for i := 0; i < n/2; i++ {
		temp[i], temp[n-i-1] = temp[n-i-1], temp[i]
	}
	return string(temp)
}
