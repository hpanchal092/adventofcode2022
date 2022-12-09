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

func noDup(s string) bool {
	output := true
	for i := 0; i < len(s) && output; i++ {
		output = strings.Index(s, string(s[i])) == strings.LastIndex(s, string(s[i]))
	}

	return output
}

func Part1(data *bufio.Scanner) int {
	data.Scan()
	stream := data.Text()
    buffer := ""
	pos := 4

	for pos < len(stream) {
		buffer = stream[pos-4 : pos]
        if noDup(buffer) {
            break
        }
        pos++
	}

	return pos
}

func Part2(data *bufio.Scanner) int {
	data.Scan()
	stream := data.Text()
    buffer := ""
	pos := 14

	for pos < len(stream) {
		buffer = stream[pos-14 : pos]
        if noDup(buffer) {
            break
        }
        pos++
	}

	return pos
}
