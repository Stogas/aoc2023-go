// started        2023-12-01 12:55;
// finished part1 2023-12-01 13:23, 'go run' time s, run time after 'go build' s
// finished part2 2023-12-01 14:20, 'go run' time s, run time after 'go build' s

package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string
//go:embed test.txt
var testInput string
//go:embed test2.txt
var testInput2 string

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

	var ans int
	switch part {
	case 1:
		if test {
			input = testInput
		}
		ans = part1(input)
	case 2:
		if test {
			input = testInput2
		}
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
	parsed := parseInput(input)
	parsedReverse := ReverseStringSlice(parsed)

	stringDigits := []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}
	stringDigitsReverse := ReverseStringSlice(stringDigits)

	// fmt.Println("zero: ", stringDigits[0], ", nine: ", stringDigits[9])
	// fmt.Println("zero: ", stringDigitsReverse[0], ", nine: ", stringDigitsReverse[9])

	firstNumREQuery := `\d`
	for _, digit := range stringDigits {
		firstNumREQuery += `|` + digit
	}
	lastNumREQuery := `\d`
	for _, digit := range stringDigitsReverse {
		lastNumREQuery += `|` + digit
	}

	firstNumRE := regexp.MustCompile(firstNumREQuery)
	lastNumRE := regexp.MustCompile(lastNumREQuery)
	var sum int

	for i, line := range parsed {
		lineReverse := parsedReverse[i]

		first := firstNumRE.FindString(line)
		last := ReverseString(lastNumRE.FindString(lineReverse))

		firstString, err := DigitToString(first, stringDigits)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Converting digits to strings failed with error: %v\n", err)
			os.Exit(2)
		}
		lastString, err := DigitToString(last, stringDigits)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Converting digits to strings failed with error: %v\n", err)
			os.Exit(2)
		}

		realNum := stringToInt(firstString + lastString)
		sum += realNum
	}

	return sum
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

func ReverseStringSlice(s []string) (result []string) {
	result = make([]string, len(s))
  for i, v := range s {
		result[i] = ReverseString(v)
  }
  return 
}

func ReverseString(s string) string {
	var result string
	for _, b := range s {
		result = string(b) + result
	}
	return result
}

func DigitToString(s string, m []string) (string, error) {
	i, err := strconv.Atoi(s)
	if err == nil {
		return strconv.Itoa(i), nil
	}
	for i, v := range m {
		if v == s {
			return strconv.Itoa(i), nil
		}
	}
	return "-1", errors.New("Couldn't convert digit `" + s + "` to integer!")
}