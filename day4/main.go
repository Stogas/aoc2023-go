// started        at: 2023-12-04 21:01
// finished part1 at: 2023-12-04 21:47
// pause start at:    2023-12-04 22:02
// resume at:         2023-12-04 22:06
// finished part2 at: 2023-12-04 22:11

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
//go:embed test.txt
var testInput string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
	testInput = strings.TrimRight(testInput, "\n")
	if len(testInput) == 0 {
		panic("empty test.txt file")
	}
}

func main() {
	var part int
	var test bool
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.BoolVar(&test, "test", false, "run with test.txt inputs?")
	flag.Parse()
	fmt.Println("Running part", part, ", test inputs = ", test)

  if test {
		input = testInput
	}

	var ans int
	switch part {
	case 1:
		ans = part1(input)
	case 2:
		ans = part2(input)
	}
	fmt.Println("Output:", ans)
}

type card struct {
	winningNumbers []int
	myNumbers []int
	multiplier int
}

// returns common elements between a and b
func innerJoin(a []int, b []int) (c []int) {
	for _, i := range a {
		for _, j := range b {
			if i == j {
				c = append(c, i)
			}
		}
	}
	return
}

func part1(input string) int {
	cards := parseInput(input)

	sum := 0

	for _, a := range cards {
		// finally, some good fucking spaghetti
		sum += int(math.Pow(2, float64(len(innerJoin(a.winningNumbers, a.myNumbers)) - 1)))
	}

	return sum
}

func part2(input string) int {
	cards := parseInput(input)

	sum := 0

	for i, a := range cards {
		j := i
		for j := j + len(innerJoin(a.winningNumbers, a.myNumbers)); j > i; j-- {
			cards[j].multiplier += cards[i].multiplier
		}
		sum += cards[i].multiplier
	}

	return sum
}

func parseInput(input string) (parsedInput []card) {
	for _, line := range strings.Split(input, "\n") {
		c := card{multiplier: 1}

		_, line, _ = strings.Cut(line, `: `)
		winningNumbers, myNumbers, _ := strings.Cut(line, ` | `)

		for _, nString := range strings.Fields(winningNumbers) {
			c.winningNumbers = append(c.winningNumbers, stringToInt(nString))
		}

		for _, nString := range strings.Fields(myNumbers) {
			c.myNumbers = append(c.myNumbers, stringToInt(nString))
		}

		parsedInput = append(parsedInput, c)
	}
	return
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}