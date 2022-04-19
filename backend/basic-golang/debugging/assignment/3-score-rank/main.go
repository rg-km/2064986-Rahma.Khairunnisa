package main

import "fmt"

func main() {
	/*
		Score Rank:
		90-100: A
		80-89: B
		70-79: C
		60-69: D
		0-59: E
	*/
	res := ScoreRank(50)
	fmt.Println(res)

	// Try correct answer:
	// resCorrect := ScoreRankCorrect(arr)
	// fmt.Println(resCorrect)
}

func ScoreRank(score int) string {
	if score > 100 || score < 0 {
		return "INVALID"
	}
	var rank string
	switch {
	case score < 60:
		rank = "E"
	case score < 70:
		rank = "D"
	case score < 80:
		rank = "C"
	case score < 90:
		rank = "B"
	case score <= 100:
		rank = "A"
	default:
		return "INVALID"
	}

	return rank
}

func ScoreRankCorrect(score int) string {
	//return "" // TODO: replace this
	if score > 100 || score < 0 {
		return "INVALID"
	}
	var rank string
	switch {
	case score < 60:
		rank = "E"
	case score < 70:
		rank = "D"
	case score < 80:
		rank = "C"
	case score < 90:
		rank = "B"
	case score <= 100:
		rank = "A"
	default:
		return "INVALID"
	}

	return rank
}
