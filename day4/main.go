package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

type Assignment struct {
	min int
	max int
}

func makeAssignment(assignment string) *Assignment {
	nums := strings.Split(assignment, "-")
	min, _ := strconv.Atoi(nums[0])
	max, _ := strconv.Atoi(nums[1])
	return &(Assignment{min, max})
}

func Part1(data *bufio.Scanner) int {
	score := 0

	for data.Scan() {
		line := data.Text()

		assignments := strings.Split(line, ",")
		ass1 := makeAssignment(assignments[0])
		ass2 := makeAssignment(assignments[1])

		fitsInAss1 := ass2.min <= ass1.min && ass2.max >= ass1.max
		fitsInAss2 := ass1.min <= ass2.min && ass1.max >= ass2.max

		if fitsInAss1 || fitsInAss2 {
			score++
		}
	}

	return score
}

func Part2(data *bufio.Scanner) int {
	score := 0

	for data.Scan() {
		line := data.Text()

		assignments := strings.Split(line, ",")
		ass1 := makeAssignment(assignments[0])
		ass2 := makeAssignment(assignments[1])

		if ass1.max >= ass2.min && ass1.min <= ass2.max {
			score++
		}
	}

	return score
}
