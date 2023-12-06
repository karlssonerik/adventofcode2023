package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var gameNoRe = "Game (\\d+):.*"
var set = ".*: (.+)$"
var colorNo = "(\\d+) %s"

var colors = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	part1()
	pow := 0
	for _, in := range input {
		sets := regexp.MustCompile(set).FindStringSubmatch(in)

		setsz := strings.Split(sets[1], ";")

		localMin := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, set := range setsz {

			for color := range colors {
				subs := regexp.MustCompile(fmt.Sprintf(colorNo, color)).FindStringSubmatch(set)
				if len(subs) > 1 {
					s, err := strconv.Atoi(subs[1])
					if err != nil {
						panic(err)
					}

					if s > localMin[color] {
						localMin[color] = s
					}
				}
			}

		}

		pow += localMin["red"] * localMin["blue"] * localMin["green"]

	}
	fmt.Println("pow", pow)

}

func part1() {
	sum := 0
	for _, in := range input {
		gameId := regexp.MustCompile(gameNoRe).FindStringSubmatch(in)

		sets := regexp.MustCompile(set).FindStringSubmatch(in)

		setsz := strings.Split(sets[1], ";")

		possibleGame := true

		for _, set := range setsz {
			if !possibleGame {
				break
			}

			for color, max := range colors {
				subs := regexp.MustCompile(fmt.Sprintf(colorNo, color)).FindStringSubmatch(set)
				if len(subs) > 1 {
					s, err := strconv.Atoi(subs[1])
					if err != nil {
						panic(err)
					}

					if s > max {
						possibleGame = false
						break
					}
				}
			}

		}

		if possibleGame {
			s, _ := strconv.Atoi(gameId[1])
			sum += s
		}

	}
	fmt.Println("   ", sum)
}

var example = []string{
	"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
	"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
	"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
	"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
	"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
}

var input = []string{}
