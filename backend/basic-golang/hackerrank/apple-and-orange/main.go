package main

import (
	"fmt"
)

// Apple and Orange
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/apple-and-orange/problem

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// TODO: answer here
	
	var samApple int32
	var samOrange int32
	for _, apple := range apples {
		if a+apple >= s && a+apple <= t {
			samApple++
		}
	}
	for _, orange := range oranges {
		if b+orange >= s && b+orange <= t {
			samOrange++
		}
	}
	fmt.Println(samApple, samOrange)
}

func main() {
	var apples = []int32{3, 2, -2}
	var oranges = []int32{2, 1, 5, -6}

	countApplesAndOranges(7, 11, 5, 15, apples, oranges)
}
