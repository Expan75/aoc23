package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func parse(line string) (int, error) {
	re := regexp.MustCompile("^[0-9]|[0-9]$")
	matches := re.FindAllStringSubmatch(line, -1)
	if matches != nil {
		number := strings.Join(matches[0], "")
		return strconv.Atoi(number)
	}
	return 0, nil
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
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
