package main

import (
	"bytes"
	"html/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan function pada template.
// Lengkapi function CalculateScore yang berfungsi untuk menjumlahkan total score dari users
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Nama: Score", contoh: "Roger: 1000"
// Pada bagian terakhir, cetak "Total Score: total", contoh: "Total Score: 50000"

type UserRank struct {
	Name  string
	Email string
	Rank  int
	Score int
}

type Leaderboard struct {
	Users []*UserRank
}

func CalculateScore(leaderboard Leaderboard) int {
	// TODO: answer here
	total := 0

	for _, user := range leaderboard.Users {
		total += user.Score
	}
	return total
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	// TODO: answer here
	funcMap := template.FuncMap{
		"CalculateScore": CalculateScore,
	}

	textTemplate := `{{ range .Users }}{{ .Name }}: {{ .Score }}{{ end }}Total Score: {{ CalculateScore . }}`

	var b bytes.Buffer

	tmpl, err := template.New("test").Funcs(funcMap).Parse(textTemplate)
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(&b, leaderboard)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil

}

func main() {
	users := []*UserRank{
		{
			Name:  "Roger",
			Rank:  1,
			Score: 1000,
		},
		{
			Name:  "Tony",
			Rank:  2,
			Score: 900,
		},
		{
			Name:  "Bruce",
			Rank:  3,
			Score: 800,
		},
		{
			Name:  "Natasha",
			Rank:  4,
			Score: 700,
		},
		{
			Name:  "Clint",
			Rank:  5,
			Score: 600,
		},
	}

	leaderboardObject := Leaderboard{
		Users: users,
	}

	b, err := ExecuteToByteBuffer(leaderboardObject)
	if err != nil {
		panic(err)
	}

	println(string(b))
}
