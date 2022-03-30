package main

import "fmt"

//interface adalah tipe data

//method signature

type circle struct {
	Radius int
	//2 method, menghitung luas, menghitung keliling
}
func (c Circle) Luas() float32{
	return math.Pi * float32(c.Radius) * float32(c.Radius)
}

func(c.Circle) Keliling() float32{
	return 2 * math.Pi * float32(c.Radius)
}

type Square struct {
	Width int
	Height int
	// 2 method, menghitung luas
}

func(s Square) Luas() float32 {
	return float32(s.Width) * float32(s.Height)
}

func(s Square) Keliling() float32 {}