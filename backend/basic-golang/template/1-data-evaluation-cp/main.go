package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Dari contoh yang telah diberikan, cobalah untuk mempraktikkan data evaluation pada template.
// Lengkapi function ExecuteToByteBuffer sehingga objek dari struct Account dapat ter-render ke dalam template.
// Gunakan bytes.Buffer sebagai io.Writer pada template.
// Lengkapi juga textTemplate, sehingga variabel dari objek Account dapat ter-render ke dalam template.
// Contoh:
// acc := {Name: "Tony", Number: "1002321", Balance: 1000}
// Akun Tony dengan Nomor 1002321 memiliki saldo sebesar $1000.

type Account struct {
	Name    string
	Number  string
	Balance int
}

func ExecuteToByteBuffer(account Account) ([]byte, error) {
	// TODO: answer here
	textTemplate := "Akun {{.Name}} dengan Nomor {{.Number}} memiliki saldo sebesar ${{.Balance}}."

	t, err := template.New("account").Parse(textTemplate)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer

	err = t.Execute(&b, account)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil

}
func main() {
	acc := Account{
		Name:    "Tony",
		Number:  "1002321",
		Balance: 1000,
	}
	
	b, err := ExecuteToByteBuffer(acc)

	
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	
}
