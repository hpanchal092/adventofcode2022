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

	want := 157

	data := bufio.NewScanner(raw)
	ans := Part1(data)

	if ans != want {
		t.Fatalf("Value returned is %d, want %d", ans, want)
	}
}

func TestPriorityOf(t *testing.T) {
	want := 39
	ans := priorityOf('M')

	if ans != want {
		t.Fatalf("Priority returned for m is %d, want %d", want, ans)
	}
}

func TestPart2(t *testing.T) {
	raw, err := os.Open("./test.txt")

	if err != nil {
		t.Fatalf("Error opening file, %v", err)
	}

	want := 70

	data := bufio.NewScanner(raw)
	ans := Part2(data)

	if ans != want {
		t.Fatalf("Value returned is %d, want %d", ans, want)
	}
}
