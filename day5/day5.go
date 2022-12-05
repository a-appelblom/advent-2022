package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1 := false

	fileReader, err := os.Open("inputDay5.txt")
	if err != nil {
		fmt.Println("Something went wrong", err)
	}

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)

	diagram := readDiagram(fileScanner)
	fmt.Println("Start", diagram)

	if part1 {
		for fileScanner.Scan() {
			line := fileScanner.Text()
			amount, from, to := numberExtractor(line)

			for i := 0; i < amount; i++ {
				// x, a = a[len(a)-1], a[:len(a)-1]
				var crate string
				var movedCrates []string
				crate, diagram[from] = diagram[from][len(diagram[from])-1], diagram[from][:len(diagram[from])-1]

				movedCrates = append(movedCrates, crate)
				diagram[to] = append(diagram[to], movedCrates...)
			}
		}
	}

	if !part1 {
		for fileScanner.Scan() {
			line := fileScanner.Text()
			amount, from, to := numberExtractor(line)

			var movedCrates []string
			movedCrates, diagram[from] = diagram[from][len(diagram[from])-amount:], diagram[from][:len(diagram[from])-amount]

			diagram[to] = append(diagram[to], movedCrates...)
		}
	}
	fmt.Println("End", diagram)
	fmt.Println(diagram[1])
	fmt.Println(diagram[2])
	fmt.Println(diagram[3])
	fmt.Println(diagram[4])
	fmt.Println(diagram[5])
	fmt.Println(diagram[6])
	fmt.Println(diagram[7])
	fmt.Println(diagram[8])
	fmt.Println(diagram[9])

	fileReader.Close()
}

func moveCrate(from, to []string) ([]string, []string) {
	var crate string
	crate = from[len(from)-1]
	from = append(from, from[:len(from)-2]...)
	to = append(to, crate)

	fmt.Println("moved")
	return from, to
}

func numberExtractor(s string) (int, int, int) {
	var amount, from, to int64

	splitted := strings.Split(s, " ")

	amount, _ = strconv.ParseInt(splitted[1], 10, 0)
	from, _ = strconv.ParseInt(splitted[3], 10, 0)
	to, _ = strconv.ParseInt(splitted[5], 10, 0)

	return int(amount), int(from), int(to)
}

func readDiagram(fileScanner *bufio.Scanner) map[int][]string {
	index := 1
	diagram := make(map[int][]string)
	lineLength := 35
	crateLength := 3
	gapLength := 1
	emptyCrate := "   "

	for fileScanner.Scan() && fileScanner.Text() != "" {
		index = 1
		pos := 0
		line := fileScanner.Text()
		for pos < lineLength {
			crate := line[pos : pos+crateLength]
			if crate != emptyCrate && crate[0] == '[' {
				diagram[index] = append([]string{string(crate[1])}, diagram[index]...)
			}
			pos += crateLength
			if pos == 35 {
				continue
			}
			pos += gapLength
			index++
		}
	}
	return diagram
}
