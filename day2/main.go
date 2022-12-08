package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	raw, err := os.Open("./input.txt")
	check(err)

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
		line := data.Text()

		switch line[2] {
		case 'X':
			score += 1 + rps(line[0], 'C', 'A')
		case 'Y':
			score += 2 + rps(line[0], 'A', 'B')
		case 'Z':
			score += 3 + rps(line[0], 'B', 'C')
		}
	}

	return score
}

func rps(input byte, win byte, draw byte) int {
	output := 0
	switch input {
	case win:
		output = 6
	case draw:
		output = 3
	}
	return output
}

func Part2(data *bufio.Scanner) int {
	score := 0

	for data.Scan() {
		line := data.Text()
		switch line[2] {
		case 'X':
			switch line[0] {
			case 'A':
				score += 3
			case 'B':
				score += 1
			case 'C':
				score += 2
			}
		case 'Y':
			score += 3
			switch line[0] {
			case 'A':
				score += 1
			case 'B':
				score += 2
			case 'C':
				score += 3
			}
		case 'Z':
			score += 6
			switch line[0] {
			case 'A':
				score += 2
			case 'B':
				score += 3
			case 'C':
				score += 1
			}
		}
	}

	return score
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
