package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Print("dummyCommit")
}

func AddString(fileName string, stringToAdd string) error {
	// TODO: replace this
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()
	if _, err = f.WriteString("\nThis is a new string to a line"); err != nil {
		return err
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	fmt.Printf("%s", data)
	return nil
}
