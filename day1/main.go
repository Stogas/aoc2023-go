// started        2023-12-01 12:55;
// finished part1 2023-12-01 13:23, 'go run' time s, run time after 'go build' s
// finished part2 , 'go run' time s, run time after 'go build' s

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"regexp"
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

// I know this is complex, but I really didn't want to make a loop for every character of a line,
// and I'm quite comfortable with regex'es
func part1(input string) int {
	parsed := parseInput(input)

	firstNumRE := regexp.MustCompile(`\d`)
	lastNumRE := regexp.MustCompile(`(\d)\D*$`)
	var sum int

	for _, line := range parsed {
		first := firstNumRE.FindString(line)
		last := lastNumRE.FindStringSubmatch(line)[1]
		realNum := stringToInt(first + last)
		sum += realNum
	}

	return sum
}

func part2(input string) int {
	return 0
}

func parseInput(input string) (parsedInput []string) {
	for _, line := range strings.Split(input, "\n") {
		parsedInput = append(parsedInput, line)
	}
	return parsedInput
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}