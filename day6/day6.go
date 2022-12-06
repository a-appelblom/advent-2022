package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1 := false

	fileReader, err := os.Open("inputDay6.txt")
	if err != nil {
		fmt.Println("Something went wrong", err)
	}

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)
	if part1 {
		for fileScanner.Scan() {
			var lastFour []rune
			line := fileScanner.Text()
			counter := 0
			for _, char := range line {
				counter++
				if len(lastFour) < 3 {
					lastFour = append(lastFour, char)
				} else {
					lastFour = append(lastFour, char)
					match := hasDuplicates(lastFour)
					fmt.Println(match)
					if !match {
						fmt.Println("no matched: ", lastFour, "counter: ", counter)
						break
					}
					lastFour = lastFour[1:]
				}
			}
		}
	}

	if !part1 {
		for fileScanner.Scan() {
			var lastFourteen []rune
			line := fileScanner.Text()
			counter := 0
			for _, char := range line {
				counter++
				if len(lastFourteen) < 13 {
					lastFourteen = append(lastFourteen, char)
				} else {
					lastFourteen = append(lastFourteen, char)
					match := hasDuplicates(lastFourteen)
					fmt.Println(match)
					if !match {
						fmt.Println("no matched: ", lastFourteen, "counter: ", counter)
						break
					}
					lastFourteen = lastFourteen[1:]
				}
			}
		}
	}

	fileReader.Close()
}

func hasDuplicates(chars []rune) bool {
	fmt.Println(chars)
	for i := 0; i < len(chars); i++ {
		for j := i + 1; j < len(chars); j++ {
			if chars[i] == chars[j] {
				return true
			}
		}
	}
	return false
}
