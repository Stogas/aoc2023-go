// started        2023-12-02 10:45;
// finished part1 2023-12-02 11:35, 'go run' time s, run time after 'go build' s
// finished part2 2023-12-02 11:39, 'go run' time s, run time after 'go build' s

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

type gameRound struct {
	red int
	green int
	blue int
}

type game struct {
	id int
	rounds []gameRound
}

func gamePossible(gameData game, r int, g int, b int) bool {
		for _, set := range gameData.rounds {
			if set.red > r || set.green > g || set.blue > b {
				return false
			}
		}
		return true
}

func gameFewest(gameData game) (r int, g int, b int) {
	for _, set := range gameData.rounds {
		if set.red > r {
			r = set.red
		}
		if set.green > g {
			g = set.green
		}
		if set.blue > b {
			b = set.blue
		}
	}
	return
}

func part1(input string) int {
	games := parseGames(input)

	var sum int

	for _, game := range games {
		if gamePossible(game, 12, 13, 14) {
			sum += game.id
		}
	}

	return sum
}

func part2(input string) int {
	games := parseGames(input)

	var sum int

	for _, game := range games {
		r, g, b := gameFewest(game)
		sum += r * g * b
	}

	return sum
}

func parseGames(input string) (games []game) {
	for _, line := range strings.Split(input, "\n") {
		var aGame game

		gameSuffix, gameData, _ := strings.Cut(line, ": ")

		aGame.id = stringToInt(strings.Fields(gameSuffix)[1])
		for _, set := range strings.Split(gameData, "; ") {
			var roundData gameRound

			for _, hand := range strings.Split(set, ", ") {
				number := stringToInt(strings.Fields(hand)[0])
				color := strings.Fields(hand)[1]
				switch color {
				case "red":
					if number > roundData.red {
						roundData.red = number
					}
				case "green":
					if number > roundData.green {
						roundData.green = number
					}
				case "blue":
					if number > roundData.blue {
						roundData.blue = number
					}
				}
			}
			aGame.rounds = append(aGame.rounds, roundData)
		}

		games = append(games, aGame)
	}
	return
}

func stringToInt(input string) int {
	output, _ := strconv.Atoi(input)
	return output
}