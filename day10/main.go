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
	fmt.Printf("Part 2:\n%s\n", ans2)
}

func Part1(data *bufio.Scanner) int {
	output := 0
	check := 20
	cycle := 0
	xReg := 1

	for data.Scan() {
		line := data.Text()
		instruction := strings.Fields(line)

		cycle++
		output += xReg * checkCycle(cycle, &check)

		if instruction[0] == "addx" {
			num, _ := strconv.Atoi(instruction[1])
			cycle++
			output += xReg * checkCycle(cycle, &check)

			xReg += num
		}
	}

	return output
}

func checkCycle(cycle int, check *int) int {
	if cycle == *check {
		if *check != 220 {
			*check += 40
		}
		return cycle
	}
	return 0
}

func Part2(data *bufio.Scanner) string {
	output := ""
	cycle := 0
	xReg := 1
	pos := 0

	for data.Scan() {
		line := data.Text()
		instruction := strings.Fields(line)

		cycle++
		pos = (cycle - 1) % 40
		output += getPixel(pos, xReg)
		if cycle%40 == 0 {
			output += "\n"
		}

		if instruction[0] == "addx" {
			num, _ := strconv.Atoi(instruction[1])

			cycle++
            pos = (cycle - 1) % 40
            output += getPixel(pos, xReg)
			if cycle%40 == 0 {
				output += "\n"
			}

			xReg += num
		}
	}

	return output
}

func getPixel(pos int, xReg int) string {
	if xReg-1 <= pos && pos <= xReg+1 {
		return "#"
	}
	return "."
}
