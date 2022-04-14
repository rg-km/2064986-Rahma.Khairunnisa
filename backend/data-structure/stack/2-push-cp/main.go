package main

import (
	"errors"
)

// Dari inisiasi stack dengan jumlah maksimal elemen 10, cobalah implementasikan operasi push.

var ErrStackOverflow = errors.New("stack overflow")

type Stack struct {
	// TODO: answer here
	Top  int
    Data []int
}

func NewStack(size int) Stack {
	// TODO: answer here
	return Stack {
		Top: -1,
		Data: make([]int, size)
	}

}

func (s *Stack) Push(Elemen int) error {
	// TODO: answer here
	if s.Top == len(s.Data) {
		return ErrStackOverflow
	}
	s.Top ++
    s.Data = append(s.Data, Elemen)
}
