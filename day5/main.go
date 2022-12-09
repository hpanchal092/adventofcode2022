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
	fmt.Printf("Part 1: %s\n", ans)

	raw.Seek(0, 0)
	data = bufio.NewScanner(raw)

	ans2 := Part2(data)
	fmt.Printf("Part 2: %s\n", ans2)
}

func Part1(data *bufio.Scanner) string {
	output := ""

	lineBuff := make([]string, 0, 10)
	for data.Scan() && data.Text() != "" {
		lineBuff = append(lineBuff, data.Text())
	}
	lineBuff[len(lineBuff)-1] = ""

	stacks := make([]string, 0, 10)
	for i := len(lineBuff) - 1; i >= 0; i-- {
		line := lineBuff[i]

		for currStack := 0; currStack*4+1 < len(line); currStack++ {
			if len(stacks) <= currStack {
				stacks = append(stacks, "")
			}
			crate := string(line[currStack*4+1])
			if crate == " " {
				continue
			}
			stacks[currStack] += crate
		}
	}

	for data.Scan() {
		line := data.Text()
		instructions := strings.Fields(line)
		numCrates, _ := strconv.Atoi(instructions[1])
		fromStack, _ := strconv.Atoi(instructions[3])
		toStack, _ := strconv.Atoi(instructions[5])

		for i := 0; i < numCrates; i++ {
			stacks[toStack-1] += pop(&stacks[fromStack-1])
		}
	}

	for i := 0; i < len(stacks); i++ {
		output += pop(&stacks[i])
	}

	return output
}

func pop(s *string) string {
	output := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return string(output)
}

func popMulti(s *string, n int) string {
	output := (*s)[len(*s)-n:]
	(*s) = (*s)[:len(*s)-n]
	return output
}

func Part2(data *bufio.Scanner) string {
	output := ""

	lineBuff := make([]string, 0, 10)
	for data.Scan() && data.Text() != "" {
		lineBuff = append(lineBuff, data.Text())
	}
	lineBuff[len(lineBuff)-1] = ""

	stacks := make([]string, 0, 10)
	for i := len(lineBuff) - 1; i >= 0; i-- {
		line := lineBuff[i]

		for currStack := 0; currStack*4+1 < len(line); currStack++ {
			if len(stacks) <= currStack {
				stacks = append(stacks, "")
			}
			crate := string(line[currStack*4+1])
			if crate == " " {
				continue
			}
			stacks[currStack] += crate
		}
	}

	for data.Scan() {
		line := data.Text()
		instructions := strings.Fields(line)
		numCrates, _ := strconv.Atoi(instructions[1])
		fromStack, _ := strconv.Atoi(instructions[3])
		toStack, _ := strconv.Atoi(instructions[5])

		stacks[toStack-1] += popMulti(&stacks[fromStack-1], numCrates)
	}

	for i := 0; i < len(stacks); i++ {
		output += pop(&stacks[i])
	}

	return output
}
