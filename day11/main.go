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

// my enum lmao
type Operation uint64

const (
	Add      Operation = 0
	Multiply Operation = 1
)

type Monkey struct {
	items        []uint64
	operand1     int
	operand2     int
	operation    Operation
	check        int
	trueThrow    int
	falseThrow   int
	inspectCount uint64
}

func Part1(data *bufio.Scanner) uint64 {
	monkeys := make([]*Monkey, 0, 10)

	// fill out the monkeys 🐵
	currMonkey := 0
	for data.Scan() {
		line := data.Text()
		words := strings.Fields(line)

		if line == "" {
			currMonkey++
			continue
		}

		if words[0] == "Monkey" {
			monkey := new(Monkey)
			monkeys = append(monkeys, monkey)
			currMonkey, _ = strconv.Atoi(string(words[1][0]))
		}
		if words[0] == "Starting" {
			fillMonkeyItems(monkeys[currMonkey], words)
		}
		if words[0] == "Operation:" {
			fillMonkeyOperation(monkeys[currMonkey], words)
		}
		if words[0] == "Test:" {
			monkeys[currMonkey].check, _ = strconv.Atoi(words[3])
		}
		if words[0] == "If" && words[1] == "true:" {
			monkeys[currMonkey].trueThrow, _ = strconv.Atoi(words[5])
		}
		if words[0] == "If" && words[1] == "false:" {
			monkeys[currMonkey].falseThrow, _ = strconv.Atoi(words[5])
		}
	}

	// perform rounds
	for i := 0; i < 20; i++ {
		performRound1(monkeys)
	}

	var max uint64
	var max2 uint64
	for _, m := range monkeys {
		if m.inspectCount > max {
			max2 = max
			max = m.inspectCount
		} else if m.inspectCount > max2 {
			max2 = m.inspectCount
		}
	}

	return max * max2
}

func performRound1(monkeys []*Monkey) {
	for _, m := range monkeys {
		for len(m.items) > 0 {
			// increment inspectCount
			m.inspectCount++

			// perform operation
			performOperation(m, &m.items[0])

			// get bored
			m.items[0] /= 3

			// perform check and throw
			var throwMonkey *Monkey
			if m.items[0]%uint64(m.check) == 0 {
				throwMonkey = monkeys[m.trueThrow]
			} else {
				throwMonkey = monkeys[m.falseThrow]
			}
			throwMonkey.items = append(throwMonkey.items, m.items[0])
			m.items = m.items[1:]
		}
	}
}

func performOperation(m *Monkey, item *uint64) {
	var op1 uint64
	var op2 uint64

	if m.operand1 == -1 {
		op1 = *item
	} else {
		op1 = uint64(m.operand1)
	}

	if m.operand2 == -1 {
		op2 = *item
	} else {
		op2 = uint64(m.operand2)
	}

	if m.operation == Multiply {
		*item = op1 * op2
	} else if m.operation == Add {
		*item = op1 + op2
	}
}

func fillMonkeyItems(m *Monkey, words []string) {
	for i := 2; i < len(words)-1; i++ {
		item, _ := strconv.ParseUint(words[i][:len(words[i])-1], 10, 64)
		m.items = append(m.items, item)
	}
	item, _ := strconv.ParseUint(words[len(words)-1], 10, 64)
	m.items = append(m.items, item)
}

func fillMonkeyOperation(m *Monkey, words []string) {
	if words[3] == "old" {
		m.operand1 = -1
	} else {
		m.operand1, _ = strconv.Atoi(words[3])
	}

	if words[5] == "old" {
		m.operand2 = -1
	} else {
		m.operand2, _ = strconv.Atoi(words[5])
	}

	if words[4] == "*" {
		m.operation = Multiply
	} else if words[4] == "+" {
		m.operation = Add
	}
}

func Part2(data *bufio.Scanner) uint64 {
	monkeys := make([]*Monkey, 0, 10)

	// fill out the monkeys 🐵
	currMonkey := 0
	for data.Scan() {
		line := data.Text()
		words := strings.Fields(line)

		if line == "" {
			currMonkey++
			continue
		}

		if words[0] == "Monkey" {
			monkey := new(Monkey)
			monkeys = append(monkeys, monkey)
			currMonkey, _ = strconv.Atoi(string(words[1][0]))
		}
		if words[0] == "Starting" {
			fillMonkeyItems(monkeys[currMonkey], words)
		}
		if words[0] == "Operation:" {
			fillMonkeyOperation(monkeys[currMonkey], words)
		}
		if words[0] == "Test:" {
			monkeys[currMonkey].check, _ = strconv.Atoi(words[3])
		}
		if words[0] == "If" && words[1] == "true:" {
			monkeys[currMonkey].trueThrow, _ = strconv.Atoi(words[5])
		}
		if words[0] == "If" && words[1] == "false:" {
			monkeys[currMonkey].falseThrow, _ = strconv.Atoi(words[5])
		}
	}

	// perform rounds
	for i := 0; i < 10000; i++ {
		performRound2(monkeys)
	}

	var max uint64
	var max2 uint64
	for _, m := range monkeys {
		if m.inspectCount > max {
			max2 = max
			max = m.inspectCount
		} else if m.inspectCount > max2 {
			max2 = m.inspectCount
		}
	}

	return max * max2
}

func performRound2(monkeys []*Monkey) {
	for _, m := range monkeys {
		for len(m.items) > 0 {
			// increment inspectCount
			m.inspectCount++

			// perform operation
			performOperation(m, &m.items[0])

			// perform check and throw
			var throwMonkey *Monkey
			if m.items[0]%uint64(m.check) == 0 {
				throwMonkey = monkeys[m.trueThrow]
			} else {
				throwMonkey = monkeys[m.falseThrow]
			}
			throwMonkey.items = append(throwMonkey.items, m.items[0])
			m.items = m.items[1:]
		}
	}
}
