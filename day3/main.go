package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	raw, _ := os.Open("./input.txt")

	data := bufio.NewScanner(raw)

	ans := Part1(data)
	fmt.Printf("Part 1: %d\n", ans)

	raw.Seek(0, 0)
	data = bufio.NewScanner(raw)

	ans2 := Part2(data)
	fmt.Printf("Part 2: %d\n", ans2)
}

func Part1(data *bufio.Scanner) int {
	score := 0

	for data.Scan() {
		rucksack := data.Text()
		comp1 := rucksack[0 : len(rucksack)/2]
		comp2 := rucksack[len(rucksack)/2:]

		for i := 0; i < len(comp1); i++ {
			if strings.ContainsRune(comp1, rune(comp2[i])) {
				score += priorityOf(comp2[i])
				break
			}
		}
	}

	return score
}

func priorityOf(char byte) int {
	output := 0

	// if capital
	if char > 64 && char < 91 {
		output = int(char) - 38
	}

	// if lowercase
	if char > 96 && char < 123 {
		output = int(char) - 96
	}

	return output
}

func Part2(data *bufio.Scanner) int {
	score := 0

	for data.Scan() {
		rucksack1 := data.Text()
		data.Scan()
		rucksack2 := data.Text()
		data.Scan()
		rucksack3 := data.Text()

		for i := 0; i < len(rucksack1); i++ {
			inRucksack2 := strings.ContainsRune(rucksack2, rune(rucksack1[i]))
			inRucksack3 := strings.ContainsRune(rucksack3, rune(rucksack1[i]))
			if inRucksack2 && inRucksack3 {
				score += priorityOf(rucksack1[i])
				break
			}
		}
	}

	return score
}
