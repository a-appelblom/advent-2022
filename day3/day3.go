package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	part1 := false
	priorityTotal := 0

	fileReader, err := os.Open("inputDay3.txt")
	if err != nil {
		fmt.Println("Something went wrong", err)
	}

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)
	if part1 {
		for fileScanner.Scan() {
			line := fileScanner.Text()
			size := len(line)
			comp1 := line[:size/2]
			comp2 := line[size/2:]
			match := getMatch(comp1, comp2)
			priorityTotal += getPriority(match)
		}
		fmt.Println(priorityTotal)
	}

	if !part1 {
		var one string
		var two string
		var three string
		runsCounter := 0
		for fileScanner.Scan() {
			line := fileScanner.Text()
			if len(one) == 0 {
				one = line
			} else if len(two) == 0 {
				two = line
			} else if len(three) == 0 {
				three = line
				badge := getBadge(one, two, three)
				priorityTotal += getPriority(badge)
				one = ""
				two = ""
				three = ""
				runsCounter++
			}
		}
		fmt.Println(runsCounter)
		fmt.Println(priorityTotal)
	}

	fileReader.Close()
}

func getBadge(a, b, c string) rune {
	tmp1 := getMatch2(a, b)
	tmp2 := getMatch2(b, c)
	// tmp3 := getMatch2(a, c)
	res := getMatch(tmp2, tmp1)
	return res
}

func getMatch2(a, b string) string {
	var matches string
	for _, char := range a {
		if strings.ContainsRune(b, char) {
			matches += string(char)
		}
	}
	return matches
}

func getMatch(a, b string) rune {
	for _, char := range a {
		if strings.ContainsRune(b, char) {
			return char
		}
	}
	return 0
}

func getPriority(character rune) int {
	var a = 'a'
	var z = 'z'
	var A = 'A'
	var Z = 'Z'
	if character >= a && character <= z {
		return int(character) - 96
	}
	if character >= A && character <= Z {
		return int(character) - 64 + 26
	}
	fmt.Println("Error")
	return 0
}
