package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.Open("./input.txt")
	check(err)

	scan := bufio.NewScanner(data)

    // probably could have made the 3 variables into a slice or arr for more
    // scalability
	max := 0
	max2 := 0
	max3 := 0

	curr := 0

	for scan.Scan() {
		line := scan.Text()

		if line == "" {
			if curr > max {
				max3 = max2
				max2 = max
				max = curr
			} else if curr > max2 {
				max3 = max2
				max2 = curr
			} else if curr > max3 {
				max3 = curr
			}
			curr = 0
			continue
		}

		num, err := strconv.Atoi(line)
		check(err)
		curr += num
	}
    // run one more time at EOF, yes its jank ik
	if curr > max {
		max3 = max2
		max2 = max
		max = curr
	} else if curr > max2 {
		max3 = max2
		max2 = curr
	} else if curr > max3 {
		max3 = curr
	}
	curr = 0

	fmt.Println(max + max2 + max3)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
