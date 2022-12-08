package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	raw, _ := os.Open("./input.txt")

	data := bufio.NewScanner(raw)

	part1, part2 := Solution(data)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func Solution(data *bufio.Scanner) (int, int) {
	// probably could have made the 3 variables into a slice or arr for more
	// scalability
	max := 0
	max2 := 0
	max3 := 0

	curr := 0

	for data.Scan() {
		line := data.Text()

		if line == "" {
			if curr > max {
				max3 = max2
				max2 = max
				max = curr
			} else if curr > max2 {
				max3 = max2
				max2 = curr
			} else if curr > max3 {
				max3 = curr
			}
			curr = 0
			continue
		}

		num, _ := strconv.Atoi(line)
		curr += num
	}
	// run one more time at EOF, yes its jank ik
	if curr > max {
		max3 = max2
		max2 = max
		max = curr
	} else if curr > max2 {
		max3 = max2
		max2 = curr
	} else if curr > max3 {
		max3 = curr
	}
	curr = 0
	return max, max + max2 + max3
}
