package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	part1()
	numbers := regexp.MustCompile("(\\d+)")
	gears := regexp.MustCompile("[^\\d.#]+")
	partMap := make(map[int]map[int]bool)
	for rowIdx, row := range input {
		partMap[rowIdx] = make(map[int]bool)

		partLocs := gears.FindAllStringIndex(row, -1)
		for _, loc := range partLocs {
			partMap[rowIdx][loc[0]] = true
		}
	}

	gearToPart := make(map[string][]int)

	for rowIdx, row := range input {
		numLocs := numbers.FindAllStringIndex(row, -1)
		for _, loc := range numLocs {
			minX := loc[0] - 1
			maxX := loc[1]
			minY := rowIdx - 1
			maxY := rowIdx + 1
			x := minX
			y := minY

			num, err := strconv.Atoi(row[loc[0]:loc[1]])
			if err != nil {
				panic(err)
			}

			for y <= maxY {
				for x <= maxX {
					if partMap[y][x] {
						gearToPart[fmt.Sprintf("(%d,%d)", x, y)] = append(gearToPart[fmt.Sprintf("(%d,%d)", x, y)], num)
					}
					x += 1
				}
				x = minX
				y += 1
			}

		}
	}
	gearRatio := 0
	for _, v := range gearToPart {
		if len(v) == 2 {
			gearRatio += v[0] * v[1]
		}
	}
	fmt.Println(gearRatio)
}

func part1() {
	numbers := regexp.MustCompile("(\\d+)")
	parts := regexp.MustCompile("[^\\d.]+")
	partMap := make(map[int]map[int]bool)
	for rowIdx, row := range input {
		partMap[rowIdx] = make(map[int]bool)

		partLocs := parts.FindAllStringIndex(row, -1)
		for _, loc := range partLocs {
			partMap[rowIdx][loc[0]] = true
		}
	}
	partSum := 0
	for rowIdx, row := range input {
		numLocs := numbers.FindAllStringIndex(row, -1)
		for _, loc := range numLocs {
			minX := loc[0] - 1
			maxX := loc[1]
			minY := rowIdx - 1
			maxY := rowIdx + 1
			x := minX
			y := minY
			isPartNumber := false
			for y <= maxY {
				for x <= maxX {
					isPartNumber = isPartNumber || partMap[y][x]
					x += 1
				}
				x = minX
				y += 1
			}

			if isPartNumber {

				num, err := strconv.Atoi(row[loc[0]:loc[1]])
				if err != nil {
					panic(err)
				}
				partSum += num
			}
		}
	}
	fmt.Println("part1: ", partSum)
}

var example = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

var input = []string{}
