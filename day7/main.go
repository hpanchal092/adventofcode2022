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

type TreeNode struct {
	name      string
	size      int
	contents  []*TreeNode
	parentDir *TreeNode
}

func cd(dir *TreeNode, name string) *TreeNode {
	if name == ".." {
		return dir.parentDir
	}

	for i := 0; i < len(dir.contents); i++ {
		if dir.contents[i].name == name {
			return dir.contents[i]
		}
	}

	fmt.Printf("\nERROR DIDN'T FIND %s, RETURNING nil\n", name)
	return nil
}

func updateSizes(dir *TreeNode) int {
	for i := 0; i < len(dir.contents); i++ {
		if len(dir.contents[i].contents) == 0 {
			dir.size += dir.contents[i].size
		} else {
			dir.size += updateSizes(dir.contents[i])
		}
	}

	return dir.size
}

func findLessThan100000(dir *TreeNode, output *int) {
	for i := 0; i < len(dir.contents); i++ {
		if len(dir.contents[i].contents) != 0 {
			findLessThan100000(dir.contents[i], output)
		}
	}
	if dir.size <= 100000 {
		*output += dir.size
	}
}

func findDeletableSize(dir *TreeNode, space int, output *int) {
	for i := 0; i < len(dir.contents); i++ {
		if len(dir.contents[i].contents) != 0 {
			findDeletableSize(dir.contents[i], space, output)
		}
	}
	if dir.size >= space && dir.size < *output {
		*output = dir.size
	}
}

func Part1(data *bufio.Scanner) int {
	output := 0
	headNode := new(TreeNode)

	currNode := new(TreeNode)
	currNode.name = "/"

	headNode.contents = append(headNode.contents, currNode)
	currNode = headNode

	for data.Scan() {
		line := data.Text()
		words := strings.Fields(line)
		if words[0] == "$" && words[1] == "cd" {
			currNode = cd(currNode, words[2])
		} else if words[0] == "dir" {
			childDir := new(TreeNode)
			childDir.name = words[1]
			childDir.parentDir = currNode
			currNode.contents = append(currNode.contents, childDir)
		} else {
			file := new(TreeNode)
			file.size, _ = strconv.Atoi(words[0])
			file.name = words[1]
			file.parentDir = currNode
			currNode.contents = append(currNode.contents, file)
		}
	}
	currNode = headNode
	updateSizes(currNode)
	findLessThan100000(currNode, &output)

	return output
}

func Part2(data *bufio.Scanner) int {
	output := 0
	headNode := new(TreeNode)

	currNode := new(TreeNode)
	currNode.name = "/"

	headNode.contents = append(headNode.contents, currNode)
	currNode = headNode

	for data.Scan() {
		line := data.Text()
		words := strings.Fields(line)
		if words[0] == "$" && words[1] == "cd" {
			currNode = cd(currNode, words[2])
		} else if words[0] == "dir" {
			childDir := new(TreeNode)
			childDir.name = words[1]
			childDir.parentDir = currNode
			currNode.contents = append(currNode.contents, childDir)
		} else {
			file := new(TreeNode)
			file.size, _ = strconv.Atoi(words[0])
			file.name = words[1]
			file.parentDir = currNode
			currNode.contents = append(currNode.contents, file)
		}
	}
	currNode = headNode
	updateSizes(currNode)

	output = headNode.size
	freeSpace := 70000000 - headNode.size
	spaceRequired := 30000000 - freeSpace
	findDeletableSize(currNode, spaceRequired, &output)

	return output
}
