// started        2023-12-03 10:44;
// finished part1 2023-12-03 12:03, 'go run' time s, run time after 'go build' s
// finished part2 2023-12-03 12:58, 'go run' time s, run time after 'go build' s

package main

import (
	_ "embed"
	"errors"
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

type numberLocation struct {
	lineNumber int
	startPos   int
	endPos     int
	value      int
}

func findAllNumbers (s []string) (l []numberLocation) {
	n := -1

	for i, line := range s {
		firstFound := false
		for j, char := range line {
			_, err := strconv.Atoi(string(char))

			// This isn't an integer. Move on
			if err != nil {
				firstFound = false

				continue
			}

			// We found a character. Is it a new startPos or endPos?
			if !firstFound {
				l = append(l, numberLocation{})
				n++
				l[n].lineNumber = i
				l[n].startPos = j
				l[n].endPos = j
				firstFound = true
			} else {
				l[n].endPos = j
			}
		}
	}

	for i, number := range l {
		var value int
		var err error
		// fmt.Printf("Finding value for number. doing s[%v][%v:%v]\n", number.lineNumber, number.startPos, number.endPos)
		if number.endPos + 1 > len(s[number.lineNumber]) {
			value, err = strconv.Atoi(s[number.lineNumber][number.startPos:])
		} else {
			value, err = strconv.Atoi(s[number.lineNumber][number.startPos:number.endPos+1])
		}

		if err != nil {panic("IDK man this shouldnt happen")}

		l[i].value = value
	}

	return
}

func isSymbol(s rune) bool {
	_, err := strconv.Atoi(string(s))
	if err == nil {
		return false
	}

	if s == '.' {
		return false
	}

	return true
}

func isEnginePart(s []string, l numberLocation) bool {
	// this could be optimized, because we always scan the number itself too,
	// and all of these multiple checks are wonky.
	// Fuck it, I'm lazy

	var lineNumberUpper, lineNumberLower int
	
	if l.lineNumber > 0 {
		lineNumberUpper = l.lineNumber - 1
	} else {
		lineNumberUpper = l.lineNumber
	}

	if l.lineNumber < len(s) {
		lineNumberLower = l.lineNumber + 2
	} else {
		lineNumberLower = l.lineNumber + 1
	}

	if lineNumberLower > len(s) {
		for _, line := range s[lineNumberUpper:] {
			var startPosFirst, endPostLast int
	
			if l.startPos > 0 {
				startPosFirst = l.startPos - 1
			} else {
				startPosFirst = l.startPos
			}
		
			if l.endPos < len(line) {
				endPostLast = l.endPos + 2
			} else {
				endPostLast = l.endPos + 1
			}
	
			if endPostLast > len(line) {
				for _, char := range line[startPosFirst:] {
					if isSymbol(char) {
						return true
					}
				}
			} else {
				for _, char := range line[startPosFirst:endPostLast] {
					if isSymbol(char) {
						return true
					}
				}
			}
		}
	} else {
		for _, line := range s[lineNumberUpper:lineNumberLower] {
			var startPosFirst, endPostLast int
	
			if l.startPos > 0 {
				startPosFirst = l.startPos - 1
			} else {
				startPosFirst = l.startPos
			}
		
			if l.endPos < len(line) {
				endPostLast = l.endPos + 2
			} else {
				endPostLast = l.endPos + 1
			}
	
			if endPostLast > len(line) {
				for _, char := range line[startPosFirst:] {
					if isSymbol(char) {
						return true
					}
				}
			} else {
				for _, char := range line[startPosFirst:endPostLast] {
					if isSymbol(char) {
						return true
					}
				}
			}
		}
	}

	return false
}

func part1(input string) int {
	parsed := parseInput(input)
	possibleNumbers := findAllNumbers(parsed)
	sum := 0

	for _, p := range possibleNumbers {
		// fmt.Printf("Checking number %v, isEnginePart = %v\n", p.value, isEnginePart(parsed, p))
		if isEnginePart(parsed, p) {
			// fmt.Println("Adding number ", p.value)
			sum += p.value
		}
	}

	return sum
}

type gearLocation struct {
	lineNum int
	pos     int
}

func findAllStars(s []string) (g []gearLocation) {
	for i, line := range s {
		for j, char := range line {
			if char == '*' {
				g = append(g, gearLocation{lineNum: i, pos: j})
			}
		}
	}

	return
}

type numberForGears struct {
	firstDigitLine int
	firstDigitPos  int
}

func figureOutNumber(s []string, line int, pos int) (l numberForGears, value int) {
	l.firstDigitLine = line
	l.firstDigitPos = pos
	for l.firstDigitPos > 0 {
		_, err := strconv.Atoi(string(s[line][l.firstDigitPos - 1]))
		if err != nil {
			break
		}
		l.firstDigitPos--
	}

	var numberString string
	fmt.Println(len(s[line]))
	for x := l.firstDigitPos; x < len(s[line]); x++ {
		_, err := strconv.Atoi(string(s[line][x]))
		if err != nil {
			x--
			break
		}
		numberString += string(s[line][x])
	}

	value, _ = strconv.Atoi(numberString)

	return
}

func gearRatio(s []string, g gearLocation) (r int, e error) {
	adjacentNumbers := make(map[numberForGears]int)

	lineStart := max(g.lineNum - 1, 0)
	posStart := max(g.pos - 1, 0)
	lineEnd := min(g.lineNum + 1, len(s))
	posEnd := min(g.pos + 1, len(s[lineEnd]))

	for x := lineStart; x <= lineEnd; x++ {
		for y := posStart; y <= posEnd; y++ {
			_, err := strconv.Atoi(string(s[x][y]))
			if err != nil {
				continue
			}
			l, value := figureOutNumber(s, x, y)
			adjacentNumbers[l] = value
		}
	}

	if len(adjacentNumbers) != 2 {
		e = errors.New("This isn't a gear")
		return
	}

	r = 1
	for _, v := range adjacentNumbers {
		r *= v
	}

	return
}

func part2(input string) int {
	parsed := parseInput(input)
	possibleGears := findAllStars(parsed)
	sum := 0

	for _, p := range possibleGears {
		r, err := gearRatio(parsed, p)
		if err != nil {
			continue
		}

		sum += r
	}

	return sum
}

func parseInput(input string) (parsedInput []string) {
	for _, line := range strings.Split(input, "\n") {
		parsedInput = append(parsedInput, line)
	}
	return parsedInput
}

// func stringToInt(input string) int {
// 	output, _ := strconv.Atoi(input)
// 	return output
// }