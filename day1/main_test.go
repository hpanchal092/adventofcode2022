package main

import (
    "testing"
    "bufio"
    "os"
)

func TestPart1(t *testing.T) {
	raw, err := os.Open("./test.txt")

    if err != nil {
        t.Fatalf("Error opening file, %v", err)
    }
	data := bufio.NewScanner(raw)

    want := 24000
    ans, _ := Solution(data)

    if ans != want {
        t.Fatalf("Value is %d, want %d", ans, want)
    }
}

func TestPart2(t *testing.T) {
	raw, err := os.Open("./test.txt")

    if err != nil {
        t.Fatalf("Error opening file, %v", err)
    }
	data := bufio.NewScanner(raw)

    want := 45000
    _, ans := Solution(data)

    if ans != want {
        t.Fatalf("Value is %d, want %d", ans, want)
    }
}
