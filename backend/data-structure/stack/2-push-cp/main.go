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

}

func (s *Stack) Push(Elemen int) error {
	// TODO: answer here
	s.Top += 1
    s.Data = append(s.Data, Elemen)
}
