package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan looping pada template.
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Peringkat ke-n: Nama", contoh: "Peringkat ke-1: Roger"

type UserRank struct {
	Name  string
	Email string
	Rank  int
}

type Leaderboard struct {
	Users []*UserRank
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	// TODO: answer here
	textTemplate := `{{range .Users}}Peringkat ke-{{.Rank}}: {{.Name}} {{end}}`

	tmpl, err := template.New("test").Parse(textTemplate)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer

	err = tmpl.Execute(&b, leaderboard)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil

}

func main() {
	users := []*UserRank{
		{
			Name: "Roger",
			Rank: 1,
		},
		{
			Name: "Tony",
			Rank: 2,
		},
		{
			Name: "Bruce",
			Rank: 3,
		},
		{
			Name: "Natasha",
			Rank: 4,
		},
		{
			Name: "Clint",
			Rank: 5,
		},
	}

	leaderboardObject := Leaderboard{
		Users: users,
	}

	b, err := ExecuteToByteBuffer(leaderboardObject)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
