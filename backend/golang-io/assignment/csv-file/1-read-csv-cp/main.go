package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	fmt.Print("dummy commit")
}

func CSVToMap(data map[string]string, fileName string) (map[string]string, error) {
	// TODO: answer here
	file, err := os.Open(fileName)
	if err != nil (
		return nil, err
	)

	reader := csv.NewReader(file)

	for{
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		data[record[0]] = record[1]

	}
	return data, err

}
