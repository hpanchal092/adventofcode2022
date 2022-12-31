package main

import (
	"bufio"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	raw, err := os.Open("./test.txt")

	if err != nil {
		t.Fatalf("Error opening file, %v", err)
	}

	want := 31

	data := bufio.NewScanner(raw)
	ans := Part1(data)

	if ans != want {
		t.Fatalf("Value returned is %d, want %d", ans, want)
	}
}

func TestPart2(t *testing.T) {
	raw, err := os.Open("./test.txt")

	if err != nil {
		t.Fatalf("Error opening file, %v", err)
	}

	want := 29

	data := bufio.NewScanner(raw)
	ans := Part2(data)

	if ans != want {
		t.Fatalf("Value returned is %d, want %d", ans, want)
	}
}
