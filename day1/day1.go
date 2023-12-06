package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var regexs = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	`\d`,
}

func main() {
	sum := 0

	for _, s := range input {
		firstPos := -1
		lastPos := -1
		first := ""
		last := ""
		for _, reg := range regexs {
			re := regexp.MustCompile(reg)
			indexes := re.FindAllStringIndex(s, -1)
			for _, index := range indexes {
				if len(index) == 0 {
					continue
				}
				if index[0] < firstPos || firstPos == -1 {
					firstPos = index[0]
					first = s[index[0]:index[1]]
				}

				if index[0] > lastPos || lastPos == -1 {
					lastPos = index[0]
					last = s[index[0]:index[1]]
				}
			}
		}

		if len(first) > 1 {
			first = digit(first)
		}

		if len(last) > 1 {
			last = digit(last)
		}

		num, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}
		sum += num
	}
	fmt.Println("sum ", sum)
}

func mainPart1() {
	sum := 0
	re := regexp.MustCompile(`\d`)
	for _, s := range special {
		first := ""
		last := ""
		digits := re.FindAllString(s, -1)

		fmt.Println("meh", digits)

		first = digits[0]

		last = digits[len(digits)-1]

		num, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}
		sum += num
	}
	fmt.Println("sum ", sum)
}

func digit(s string) string {
	switch s {

	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}

	panic(s)
}

var example = []string{
	"1abc2",
	"pqr3stu8vwx",
	"a1b2c3d4e5f",
	"treb7uchet",
}

var example2 = []string{
	"two1nine",
	"eightwothree",
	"abcone2threexyz",
	"xtwone3four",
	"4nineeightseven2",
	"zoneight234",
	"7pqrstsixteen",
}

var special = []string{
	"eighthree",
	"sevenine",
	"7eighthreeeight",
	"7eighthreeeight44",
}

var input = []string{}
