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

func Part1(data *bufio.Scanner) int {
	grid := make([][]int, 0, 100)
	i := 0
	set := make(map[*int]struct{})

	for data.Scan() {
		line := data.Text()
		row := make([]int, 0, 100)
		grid = append(grid, row)

		for j := 0; j < len(line); j++ {
			num, _ := strconv.Atoi(string(line[j]))
			grid[i] = append(grid[i], num)
		}
		i++
	}

	addUp(grid, set)
	addDown(grid, set)
	addLeft(grid, set)
	addRight(grid, set)

	return len(set)
}

func addLeft(grid [][]int, set map[*int]struct{}) {
	for i := 0; i < len(grid); i++ {
		tallestTree := -1
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > tallestTree {
				set[&grid[i][j]] = struct{}{}
				tallestTree = grid[i][j]
			}
		}
	}
}

func addDown(grid [][]int, set map[*int]struct{}) {
	for i := 0; i < len(grid); i++ {
		tallestTree := -1
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[j][i] > tallestTree {
				set[&grid[j][i]] = struct{}{}
				tallestTree = grid[j][i]
			}
		}
	}
}

func addRight(grid [][]int, set map[*int]struct{}) {
	for i := 0; i < len(grid); i++ {
		tallestTree := -1
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] > tallestTree {
				set[&grid[i][j]] = struct{}{}
				tallestTree = grid[i][j]
			}
		}
	}
}

func addUp(grid [][]int, set map[*int]struct{}) {
	for i := 0; i < len(grid); i++ {
		tallestTree := -1
		for j := 0; j < len(grid[i]); j++ {
			if grid[j][i] > tallestTree {
				set[&grid[j][i]] = struct{}{}
				tallestTree = grid[j][i]
			}
		}
	}
}

func Part2(data *bufio.Scanner) int {
	output := 0
	grid := make([][]int, 0, 100)
	i := 0

	for data.Scan() {
		line := data.Text()
		row := make([]int, 0, 100)
		grid = append(grid, row)

		for j := 0; j < len(line); j++ {
			num, _ := strconv.Atoi(string(line[j]))
			grid[i] = append(grid[i], num)
		}
		i++
	}

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			upDist := checkUp(grid, i, j)
			downDist := checkDown(grid, i, j)
			leftDist := checkLeft(grid, i, j)
			rightDist := checkRight(grid, i, j)

			viewScore := upDist * downDist * leftDist * rightDist
			if output < viewScore {
				output = viewScore
			}
		}
	}

	return output
}

func checkUp(grid [][]int, i int, j int) int {
	output := 0
	for viewing := i - 1; viewing >= 0; viewing-- {
		output++
		if grid[i][j] <= grid[viewing][j] {
			break
		}
	}
	return output
}

func checkDown(grid [][]int, i int, j int) int {
	output := 0
	for viewing := i + 1; viewing < len(grid); viewing++ {
		output++
		if grid[i][j] <= grid[viewing][j] {
			break
		}
	}
	return output
}

func checkRight(grid [][]int, i int, j int) int {
	output := 0
	for viewing := j + 1; viewing < len(grid[i]); viewing++ {
		output++
		if grid[i][j] <= grid[i][viewing] {
			break
		}
	}
	return output
}

func checkLeft(grid [][]int, i int, j int) int {
	output := 0
	for viewing := j - 1; viewing >= 0; viewing-- {
		output++
		if grid[i][j] <= grid[i][viewing] {
			break
		}
	}
	return output
}
