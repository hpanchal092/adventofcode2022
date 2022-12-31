package main

import (
	"bufio"
	"fmt"
	"os"
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

type Position struct {
	i    int
	j    int
	dist int
}

type BetterSlice []Position

func (s BetterSlice) contains(p Position) bool {
	for _, val := range s {
		if p.i == val.i && p.j == val.j {
			return true
		}
	}

	return false
}

func Part1(data *bufio.Scanner) int {
	mountain := make([][]byte, 0, 50)
	start := new(Position)
	end := new(Position)

	for i := 0; data.Scan(); i++ {
		line := data.Text()
		mountain = append(mountain, make([]byte, 0, 50))
		for j, char := range line {
			mountain[i] = append(mountain[i], byte(char))
			if char == 'S' {
				start.i, start.j = i, j
                mountain[i][j] = 'a'
			}
			if char == 'E' {
				end.i, end.j = i, j
				mountain[i][j] = 'z'
			}
		}
	}

	return pathfind(mountain, *start, *end)
}

func pathfind(grid [][]byte, s Position, e Position) int {
	var visited BetterSlice
	var needToVisit BetterSlice

	visited = append(visited, s)
	needToVisit = append(needToVisit, s)

	for len(needToVisit) > 0 {
		currPos := needToVisit[0]
		needToVisit = needToVisit[1:]

		if currPos.i == e.i && currPos.j == e.j {
			return currPos.dist
		}
        currElevation := grid[currPos.i][currPos.j]

		neighbors := make([]Position, 0, 4)
		if currPos.i-1 >= 0 && grid[currPos.i-1][currPos.j] <= currElevation+1 {
			neighbors = append(neighbors, Position{currPos.i - 1, currPos.j, currPos.dist + 1})
		}
		if currPos.i+1 < len(grid) && grid[currPos.i+1][currPos.j] <= currElevation+1 {
			neighbors = append(neighbors, Position{currPos.i + 1, currPos.j, currPos.dist + 1})
		}
		if currPos.j-1 >= 0 && grid[currPos.i][currPos.j-1] <= currElevation+1 {
			neighbors = append(neighbors, Position{currPos.i, currPos.j - 1, currPos.dist + 1})
		}
		if currPos.j+1 < len(grid[currPos.i]) && grid[currPos.i][currPos.j+1] <= currElevation+1 {
			neighbors = append(neighbors, Position{currPos.i, currPos.j + 1, currPos.dist + 1})
		}
		for _, neighbor := range neighbors {
			if !visited.contains(neighbor) {
				visited = append(visited, neighbor)
				needToVisit = append(needToVisit, neighbor)
			}
		}
	}

	fmt.Println("OMG PANIC AHHH")
	return -1
}

func Part2(data *bufio.Scanner) int {
	mountain := make([][]byte, 0, 50)
	end := new(Position)

	for i := 0; data.Scan(); i++ {
		line := data.Text()
		mountain = append(mountain, make([]byte, 0, 50))
		for j, char := range line {
			mountain[i] = append(mountain[i], byte(char))
			if char == 'S' {
                mountain[i][j] = 'a'
			}
			if char == 'E' {
				end.i, end.j = i, j
				mountain[i][j] = 'z'
			}
		}
	}
    
    return pathfind2(mountain, *end)
}

func pathfind2(grid [][]byte, s Position) int {
	var visited BetterSlice
	var needToVisit BetterSlice

	visited = append(visited, s)
	needToVisit = append(needToVisit, s)

	for len(needToVisit) > 0 {
		currPos := needToVisit[0]
		needToVisit = needToVisit[1:]
        currElevation := grid[currPos.i][currPos.j]
        if currElevation == 'a' {
            return currPos.dist
        }

		neighbors := make([]Position, 0, 4)
		if currPos.i-1 >= 0 && grid[currPos.i-1][currPos.j] >= currElevation-1 {
			neighbors = append(neighbors, Position{currPos.i - 1, currPos.j, currPos.dist + 1})
		}
		if currPos.i+1 < len(grid) && grid[currPos.i+1][currPos.j] >= currElevation-1 {
			neighbors = append(neighbors, Position{currPos.i + 1, currPos.j, currPos.dist + 1})
		}
		if currPos.j-1 >= 0 && grid[currPos.i][currPos.j-1] >= currElevation-1 {
			neighbors = append(neighbors, Position{currPos.i, currPos.j - 1, currPos.dist + 1})
		}
		if currPos.j+1 < len(grid[currPos.i]) && grid[currPos.i][currPos.j+1] >= currElevation-1 {
			neighbors = append(neighbors, Position{currPos.i, currPos.j + 1, currPos.dist + 1})
		}
		for _, neighbor := range neighbors {
			if !visited.contains(neighbor) {
				visited = append(visited, neighbor)
				needToVisit = append(needToVisit, neighbor)
			}
		}
	}

	fmt.Println("OMG PANIC AHHH")
	return -1
}
