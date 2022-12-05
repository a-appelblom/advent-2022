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

	fileReader, err := os.Open("inputDay4.txt")
	if err != nil {
		fmt.Println("Something went wrong", err)
	}

	fileScanner := bufio.NewScanner(fileReader)

	fileScanner.Split(bufio.ScanLines)
	if part1 {
		counter := 0
		for fileScanner.Scan() {
			line := fileScanner.Text()
			fmt.Println(line)
			assingments := strings.Split(line, ",")
			elf1 := strings.Split(assingments[0], "-")
			elf2 := strings.Split(assingments[1], "-")
			elf1start, _ := strconv.ParseInt(elf1[0], 10, 0)
			elf1end, _ := strconv.ParseInt(elf1[1], 10, 0)
			elf2start, _ := strconv.ParseInt(elf2[0], 10, 0)
			elf2end, _ := strconv.ParseInt(elf2[1], 10, 0)
			if elf1start <= elf2start && elf1end >= elf2end {
				counter++
			} else if elf2start <= elf1start && elf2end >= elf1end {
				counter++
			}
		}
		fmt.Println(counter)
	}

	if !part1 {
		counter := 0
		for fileScanner.Scan() {
			line := fileScanner.Text()
			fmt.Println(line)
			assingments := strings.Split(line, ",")
			elf1 := strings.Split(assingments[0], "-")
			elf2 := strings.Split(assingments[1], "-")
			elf1start, _ := strconv.ParseInt(elf1[0], 10, 0)
			elf1end, _ := strconv.ParseInt(elf1[1], 10, 0)
			elf2start, _ := strconv.ParseInt(elf2[0], 10, 0)
			elf2end, _ := strconv.ParseInt(elf2[1], 10, 0)
			if elf1start >= elf2start && elf1start <= elf2end {
				counter++
			} else if elf1end >= elf2start && elf1end <= elf2end {
				counter++
			} else if elf2start >= elf1start && elf2start <= elf1end {
				counter++
			} else if elf2end >= elf1start && elf2end <= elf1end {
				counter++
			}
		}
		fmt.Println(counter)
	}

	fileReader.Close()
}
