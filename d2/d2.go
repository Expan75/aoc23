package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	id            int
	colorCount    map[string]int
	colorMaxCount map[string]int
}

func (g Game) Valid(maxColorCount map[string]int) bool {
	for k := range g.colorMaxCount {
		allowedMax, allowed := maxColorCount[k]
		observedMax, observed := g.colorMaxCount[k]
		if !(allowed && observed && observedMax <= allowedMax) {
			return false
		}
	}
	return true
}

func (g Game) minSetSize(maxColorCount map[string]int) int {
	result := 1
	for k := range g.colorCount {
		_, colorIncluded := maxColorCount[k]
		if colorIncluded {
			result *= g.colorCount[k]
		}
	}
	return result
}

func parseGame(line string) (Game, error) {
	idPattern := regexp.MustCompile("Game.[0-9]+")
	colorsPattern := regexp.MustCompile("[0-9]\\w.[a-z]+")
	id := idPattern.FindString(line)
	colors := colorsPattern.FindAllString(line, -1)

	if id == "" {
		return Game{}, errors.New("invalid game")
	}
	parsedId, _ := strconv.Atoi(strings.Split(id, " ")[1])
	colorCount := make(map[string]int, 0)
	maxColorCount := make(map[string]int, 0)

	for _, c := range colors {
		brokenColorCount := strings.Split(c, " ")
		count, _ := strconv.Atoi(brokenColorCount[0])
		color := brokenColorCount[1]

		existingCount, seenColorBefore := colorCount[color]
		existingMax, _ := maxColorCount[color]

		// increment counted observations
		if seenColorBefore {
			colorCount[color] = existingCount + count
			if count > existingMax {
				maxColorCount[color] = count
			}
		} else {
			maxColorCount[color] = count
			colorCount[color] = count
		}

	}
	return Game{parsedId, colorCount, maxColorCount}, nil
}

func main() {
	args := os.Args[1:]
	b, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println(err)
	}
	text := string(b)

	// set in example
	allowedCount := map[string]int{"red": 12, "green": 13, "blue": 14}
	validGameIdSum := 0
	totalMinSetSize := 0

	for _, line := range strings.Split(text, "\n") {
		g, err := parseGame(line)
		if err == nil && g.Valid(allowedCount) {
			validGameIdSum += g.id
		} else if err == nil {
			totalMinSetSize += g.minSetSize(allowedCount)
		}
	}
	fmt.Println("sum of valid gameIds: ", validGameIdSum)
	fmt.Println("sum of minSetSize: ", totalMinSetSize)
}
