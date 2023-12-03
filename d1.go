package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var NumericLiterals = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// Note, go ensures O(n) for regex std lib,
// as a consequenes, not supported lookbehind/ahead, which we'd need for this
// for this overlapping pattern
func parse(line string) (int, error) {
	literals := make([]string, 0, len(NumericLiterals))
	for literal := range NumericLiterals {
		literals = append(literals, literal)
	}
	pattern := "(" + strings.Join(literals, "|") + "|[0-9]" + ")"
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(line, -1)

	fmt.Println(matches)

	if len(matches) == 0 {
		return 0, nil
	}
	parsedMatches := make([]string, len(matches))
	for i, match := range matches {
		if len(match) > 1 {
			parsedMatches[i] = NumericLiterals[match]
		} else {
			parsedMatches[i] = match
		}
	}
	value := parsedMatches[0] + parsedMatches[len(parsedMatches)-1]

	fmt.Println(line)
	fmt.Println(parsedMatches)
	fmt.Println(value)

	return strconv.Atoi(value)
}

func main() {
	args := os.Args[1:]
	b, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println(err)
	}
	text := string(b)
	sum := 0
	for _, line := range strings.Split(text, "\n") {
		parsedValue, err := parse(line)
		if err != nil {
			panic(err)
		}
		sum += parsedValue
	}
	result := fmt.Sprintf("result=%d", sum)
	fmt.Println(result)
}
