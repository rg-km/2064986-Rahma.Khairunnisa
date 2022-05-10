package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Dari contoh yang diberikan, cobalah untuk mengimplementasikan sebuah http client sederhana.
// Buatlah sebuah http client dan lakukan request GET ke API https://www.metaweather.com/api/.
// Buatlah request get ke endpoint /api/location/(woeid)/(date)/ dengan nilai woeid 1047378.
// Untuk date, gunakan format YYYY/MM/dd

func main() {
	// TODO: answer here
	c := http.Client{}

	date := "2022/05/09"
	woeid := "1047378"
	url := "https://www.metaweather.com/api/location/" + woeid + "/" + date + "/"

	resp, err := c.Get(url)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}

	fmt.Printf(string(body))
}
