// skeleton template copied and modified from https://github.com/alexchao26/advent-of-code-go/blob/main/scripts/skeleton/tmpls/main.go

// started        at: ---
// finished part1 at: ---
// finished part2 at: ---

package main

import (
	_ "embed"
	"flag"
	"fmt"
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

func part1(input string) int {
	parsed := parseInput(input)
	fmt.Println(parsed)

	return 0
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (parsedInput []int) {
	for _, line := range strings.Split(input, "\n") {
		parsedInput = append(parsedInput, stringToInt(line))
	}
	return parsedInput
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}