package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var elvesCalories []int
	elvesCounter := 0
	var currentElf int = 0

	fileReader, err := os.Open("inputDay1.txt")
	if err != nil {
		fmt.Println("Something went wrong", err)
	}

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			elvesCalories = append(elvesCalories, currentElf)
			currentElf = 0
			elvesCounter += 1
		} else {
			calories, err := strconv.ParseInt(line, 10, 0)
			if err != nil {
				fmt.Println("Something went wrong", err)
			}

			currentElf += int(calories)
		}
	}

	fileReader.Close()

	fmt.Println(getThreeBiggestTotal(elvesCalories))
}

func getBiggest(inp []int) int {
	var result, tmp int
	for _, element := range inp {
		if element > tmp {
			tmp = element
			result = tmp
		}
	}
	return result
}

func getThreeBiggestTotal(inp []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(inp)))
	result := inp[0] + inp[1] + inp[2]
	return result
}
