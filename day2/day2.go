package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// rock = 1, A, X
	// Paper = 2, B, Y
	// Scissors = 3, C, Z
	part1 := false
	var total int = 0

	fileReader, err := os.Open("inputDay2.txt")
	if err != nil {
		fmt.Println("Something went wrong", err)
	}

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)
	if part1 {
		for fileScanner.Scan() {
			roundScore := 0

			line := fileScanner.Text()
			opponent, player := splitter(line)
			roundScore += compare(opponent, player)
			roundScore += getMoveScore(player)

			total += roundScore

			fmt.Println(opponent, player, roundScore)
		}
		fmt.Println("Total: ", total)
	}

	if !part1 {
		for fileScanner.Scan() {
			roundScore := 0

			line := fileScanner.Text()
			opponent, result := splitter2(line)

			move := getIdealMove(opponent, result)
			fmt.Println("Move: ", move)
			roundScore += compare(opponent, move)
			roundScore += getMoveScore(move)

			total += roundScore
		}

		fmt.Println("Total: ", total)
	}

	fileReader.Close()
}

func splitter2(str string) (string, string) {
	tmp := strings.Split(str, " ")
	p := getOpponentMove(tmp[0])
	return p, tmp[1]
}

func splitter(str string) (string, string) {
	var opponent, player string
	tmp := strings.Split(str, " ")

	opponent = getOpponentMove(tmp[0])
	player = getPlayerMove(tmp[1])

	return opponent, player
}

func getOpponentMove(a string) string {
	if a == "A" {
		return "rock"
	} else if a == "B" {
		return "paper"
	} else if a == "C" {
		return "scissor"
	} else {
		return "wrong"
	}
}

func getPlayerMove(a string) string {
	if a == "X" {
		return "rock"
	} else if a == "Y" {
		return "paper"
	} else if a == "Z" {
		return "scissor"
	} else {
		return "wrong"
	}
}

func compare(o, p string) int {
	fmt.Println(o, p)
	if (p == "rock" && o == "paper") || (p == "paper" && o == "scissor") || (p == "scissor" && o == "rock") {
		fmt.Println("win")
		return 0
	} else if (o == "rock" && p == "paper") || (o == "paper" && p == "scissor") || (o == "scissor" && p == "rock") {
		fmt.Println("Lose")
		return 6
	} else if o == p {
		fmt.Println("Draw")
		return 3
	} else {
		fmt.Println("Compare: Error")
		return 0
	}
}

func getMoveScore(p string) int {
	if p == "rock" {
		return 1
	} else if p == "paper" {
		return 2
	} else if p == "scissor" {
		return 3
	} else {
		return 0
	}
}

func getIdealMove(opponent, outcome string) string {
	fmt.Println(outcome)
	if outcome == "Y" {
		return opponent
	} else if outcome == "X" {
		if opponent == "rock" {
			return "scissor"
		}
		if opponent == "paper" {
			return "rock"
		}
		if opponent == "scissor" {
			return "paper"
		}
	} else if outcome == "Z" {
		if opponent == "rock" {
			return "paper"
		}
		if opponent == "paper" {
			return "scissor"
		}
		if opponent == "scissor" {
			return "rock"
		}
	}
	return "Error"
}
