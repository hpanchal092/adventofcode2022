package main

import (
	"bufio"
	"os"
	"testing"
)

func TestNoDupFalse(t *testing.T) {
    want := false
    ans := noDup("harsh")
    if ans != want {
        t.Fatalf("want %v, returned %v", want, ans)
    }
}

func TestNoDupTrue(t *testing.T) {
    want := true
    ans := noDup("kyle")
    if ans != want {
        t.Fatalf("want %v, returned %v", want, ans)
    }
}

func TestPart1(t *testing.T) {
	raw, err := os.Open("./test.txt")

	if err != nil {
		t.Fatalf("Error opening file, %v", err)
	}

	want := 7

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

	want := 19

	data := bufio.NewScanner(raw)
	ans := Part2(data)

	if ans != want {
		t.Fatalf("Value returned is %d, want %d", ans, want)
	}
}
