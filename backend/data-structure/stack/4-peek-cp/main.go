package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi peek.

var ErrEmptyStack = errors.New("stack is empty")

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

func (s *Stack) Peek() (int, error) {
	// TODO: answer here
	if s.Top == -1 {
        return 0, errors.New("stack is empty")
    } else {
        return s.Data[s.Top], nil
    }
}
