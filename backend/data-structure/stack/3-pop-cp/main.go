package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi pop.

var ErrStackUnderflow = errors.New("stack underflow")

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
	if s.Top == len(s.Data) {
        return errors.New("stack overflow")
    } else {
        s.Top += 1
        s.Data[s.Top] = Elemen
    }
    return nil
}

func (s *Stack) Pop() (int, error) {
	// TODO: answer here
	if s.IsEmpty() {
        return 0, errors.New("stack underflow")
    } else {
        poppedValue := s.Data[s.Top]
        s.Top -= 1
        s.Data = s.Data[:len(s.Data)-1]
        return poppedValue, nil
    }	
}
