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

	ans := Part1(data)
	fmt.Printf("Part 1: %d\n", ans)

	raw.Seek(0, 0)
	data = bufio.NewScanner(raw)

	ans2 := Part2(data)
	fmt.Printf("Part 2: %d\n", ans2)
}

type RopeNode struct {
	x int
	y int
}

func moveHead(head *RopeNode, direction byte) {
	switch direction {
	case 'U':
		head.y++
	case 'D':
		head.y--
	case 'L':
		head.x--
	case 'R':
		head.x++
	}
}

func moveTail(head *RopeNode, tail *RopeNode) {
	tailNextToHead := tail.x >= head.x-1 && tail.x <= head.x+1
	tailNextToHead = tailNextToHead && tail.y >= head.y-1 && tail.y <= head.y+1

	if !tailNextToHead {
		if head.x > tail.x {
			tail.x++
		}
		if head.x < tail.x {
			tail.x--
		}
		if head.y > tail.y {
			tail.y++
		}
		if head.y < tail.y {
			tail.y--
		}
	}
}

func Part1(data *bufio.Scanner) int {
	head := new(RopeNode)
	tail := new(RopeNode)
	instructions := make([]byte, 0, 50)

	// makeshift set lmao
	set := make(map[RopeNode]struct{})

	for data.Scan() {
		line := data.Text()
		distance, _ := strconv.Atoi(line[2:])
		direction := line[0]

		for i := 0; i < distance; i++ {
			instructions = append(instructions, direction)
		}
	}

	for _, direction := range instructions {
		moveHead(head, direction)
        moveTail(head, tail)
		set[*tail] = struct{}{}
	}

	return len(set)
}

func moveRope(rope []*RopeNode, direction byte) {
    moveHead(rope[0], direction)
	for i := 0; i < len(rope)-1; i++ {
        moveTail(rope[i], rope[i+1])
	}
}

func Part2(data *bufio.Scanner) int {
	rope := make([]*RopeNode, 10)

	for i := 0; i < 10; i++ {
		rope[i] = new(RopeNode)
	}

	instructions := make([]byte, 0, 50)

	// makeshift set lmao
	set := make(map[RopeNode]struct{})

	for data.Scan() {
		line := data.Text()
		distance, _ := strconv.Atoi(line[2:])
		direction := line[0]

		for i := 0; i < distance; i++ {
			instructions = append(instructions, direction)
		}
	}

	for _, direction := range instructions {
		// moveHeadTail(head, tail, direction)
		moveRope(rope, direction)
		set[*rope[9]] = struct{}{}
	}

	return len(set)
}
