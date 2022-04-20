package main

import (
	"fmt"
)

func communicate(output chan string) {
	c := make(chan string)
	waitGoroutine := make(chan struct{}) //hanya digunakan untuk menunggu
	greetSteve := ""
	greet := func(name string) {
		fmt.Println("send to channel c")
		//kirim "hello"+name ke channel
		// TODO: answer here
		waitGoroutine < - struct{}{}
		c <- "hello" + name

	}
	go greet("steve")

	receiveGreet := func() {
		// TODO: answer here
		greetSteve = <-c
		
	}
	go receiveGreet()

	<-waitGoroutine
	output <- greetSteve
	fmt.Println(c) //agar variabel c digunakan
}
