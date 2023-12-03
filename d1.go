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

// simple and inefficient but only non-index based go version. nothing included in std lib :(
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

// Note, go ensures O(n) for regex std lib,
// as a consequenes, not supported lookbehind/ahead, which we'd need for this
// for this overlapping pattern.
//
// One way of solving this, is to match in reverse. Regex doesn't really support that
// but we can reverse every pattern, and match the reversed string with those patterns,
// achieving a similar thing.
func parse(line string) (int, error) {
	literals := make([]string, 0, len(NumericLiterals))
	for literal := range NumericLiterals {
		literals = append(literals, literal)
	}

	pattern := strings.Join(literals, "|")
	reversedPattern := Reverse(pattern)

	reLeftToRight := regexp.MustCompile("(" + pattern + "|[1-9])")
	reRightToLeft := regexp.MustCompile("(" + reversedPattern + "|[1-9])")

	left := reLeftToRight.FindString(line)
	right := reRightToLeft.FindString(Reverse(line))

	// same for both
	if left == "" {
		return 0, nil
	}

	translatedLeftLiteral, leftLiteral := NumericLiterals[left]
	translatedRightLiteral, rightLiteral := NumericLiterals[Reverse(right)]

	if leftLiteral && rightLiteral {
		return strconv.Atoi(translatedLeftLiteral + translatedRightLiteral)
	} else if rightLiteral {
		return strconv.Atoi(left + translatedRightLiteral)
	} else {
		return strconv.Atoi(translatedLeftLiteral + right)
	}
}

// strategy 2: rolling hash for each pattern, return upon first match

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
