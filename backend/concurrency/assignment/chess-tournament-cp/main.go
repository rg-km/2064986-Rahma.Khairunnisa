package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}
func main() {
	startTournament()
}

func playMatch(i int) {
	fmt.Printf("match #%d started\n", i+1)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("match #%d finished\n", i+1)
}

func startTournament() {
	// TODO: answer here
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			playMatch(i)
			c <- i
		}(i)
	}
	
}
for i := 0; i < 10; i++ {
	<-c
}
fmt.Printf("tournament finished in %s\n", time.Since(start))
