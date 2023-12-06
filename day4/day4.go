package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	part1()
	rrr := `Card\s+(\d+):([^|]*)\|(.*)$`
	cards := make(map[int]int)
	for _, lottoCard := range input {
		matches := regexp.MustCompile(rrr).FindStringSubmatch(lottoCard)

		cardNumber, _ := strconv.Atoi(matches[1])
		cards[cardNumber] += 1

		winNumbers := toIntSlice(matches[2])

		cardNumbers := toIntSlice(matches[3])

		wins := contains(winNumbers, cardNumbers)

		cardToAdd := cardNumber
		for i := 0; i < wins; i++ {
			cardToAdd++
			if cardToAdd > len(input) {
				break
			}
			cards[cardToAdd] += cards[cardNumber]
		}

	}
	sumCards := 0
	for _, v := range cards {
		sumCards += v
	}
	fmt.Println(sumCards)
}

func part1() {
	rrr := `Card\s+(\d+):([^|]*)\|(.*)$`
	earnings := 0
	for _, lottoCard := range input {
		matches := regexp.MustCompile(rrr).FindStringSubmatch(lottoCard)

		winNumbers := toIntSlice(matches[2])

		cardNumbers := toIntSlice(matches[3])

		wins := contains(winNumbers, cardNumbers)
		if wins > 0 {
			earnings += 1 << (contains(winNumbers, cardNumbers) - 1)
		}

	}
	fmt.Println(earnings)
}

func contains(a, b []int) int {
	count := 0

	for _, win := range a {
		for _, lotto := range b {
			if win == lotto {
				count++
				break
			}
		}
	}

	return count
}

func toIntSlice(s string) []int {
	out := []int{}
	re := `(\d+)`
	matches := regexp.MustCompile(re).FindAllStringSubmatch(s, -1)
	for _, subMatch := range matches {
		i, _ := strconv.Atoi(subMatch[0])
		out = append(out, i)
	}

	return out
}

var example = []string{
	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}

var input = []string{}
