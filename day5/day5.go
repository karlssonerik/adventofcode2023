package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/schollz/progressbar/v3"
)

func main() {
	fmt.Println("part 1")
	part1()
	fmt.Println("part 2")
	part2()
}

func part2() {
	maps := []conversionMap{}
	in := input
	seedsMatches := regexp.MustCompile(seedsRe).FindStringSubmatch(in)
	seedRanges := toIntSlice(seedsMatches[1])

	mapMatches := regexp.MustCompile(mpasRe).FindAllStringSubmatch(in, -1)

	seedToSoil := getConversionMap(mapMatches[0])
	maps = append(maps, seedToSoil)
	soilToFert := getConversionMap(mapMatches[1])
	maps = append(maps, soilToFert)
	fertToWat := getConversionMap(mapMatches[2])
	maps = append(maps, fertToWat)
	watToLig := getConversionMap(mapMatches[3])
	maps = append(maps, watToLig)
	ligToTemp := getConversionMap(mapMatches[4])
	maps = append(maps, ligToTemp)
	tempToHum := getConversionMap(mapMatches[5])
	maps = append(maps, tempToHum)
	humToLoc := getConversionMap(mapMatches[6])
	maps = append(maps, humToLoc)

	MaxUint := ^uint(0)
	minSeed := int(MaxUint >> 1)

	type seedRange struct {
		start int
		end   int
	}
	srs := []seedRange{}

	sumz := int64(0)

	for i := 0; i < len(seedRanges); i += 2 {
		srs = append(srs, seedRange{
			start: seedRanges[i],
			end:   seedRanges[i] + seedRanges[i+1],
		})

		sumz += int64(seedRanges[i+1])
	}

	bar := progressbar.Default(sumz, "converting seed to location")

	for _, sr := range srs {
		for i := sr.start; i < sr.end; i++ {
			bar.Add(1)
			convValue := i
			for _, conv := range maps {
				for _, c := range conv.conversions {
					if convValue >= c.start && convValue <= c.end {
						convValue = convValue - c.diff
						break
					}
				}
			}
			if convValue < minSeed {
				minSeed = convValue
			}
		}

	}

	fmt.Println(minSeed)
}

func part1() {
	maps := []conversionMap{}
	in := input
	seedsMatches := regexp.MustCompile(seedsRe).FindStringSubmatch(in)

	seeds := toIntSlice(seedsMatches[1])

	mapMatches := regexp.MustCompile(mpasRe).FindAllStringSubmatch(in, -1)

	seedToSoil := getConversionMap(mapMatches[0])
	maps = append(maps, seedToSoil)
	soilToFert := getConversionMap(mapMatches[1])
	maps = append(maps, soilToFert)
	fertToWat := getConversionMap(mapMatches[2])
	maps = append(maps, fertToWat)
	watToLig := getConversionMap(mapMatches[3])
	maps = append(maps, watToLig)
	ligToTemp := getConversionMap(mapMatches[4])
	maps = append(maps, ligToTemp)
	tempToHum := getConversionMap(mapMatches[5])
	maps = append(maps, tempToHum)
	humToLoc := getConversionMap(mapMatches[6])
	maps = append(maps, humToLoc)

	MaxUint := ^uint(0)
	minSeed := int(MaxUint >> 1)

	for _, seed := range seeds {
		convValue := seed
		for _, conv := range maps {
			for _, c := range conv.conversions {
				if convValue >= c.start && convValue <= c.end {
					convValue = convValue - c.diff
					break
				}
			}
		}
		if convValue < minSeed {
			minSeed = convValue
		}

	}

	fmt.Println(minSeed)
}

type conversionMap struct {
	conversions []conversion
}

type conversion struct {
	start      int
	startValue int
	end        int
	diff       int
}

func getConversionMap(mapMatch []string) conversionMap {
	mapIntMatches := regexp.MustCompile(`([^\n]+)`).FindAllStringSubmatch(mapMatch[1], -1)

	out := conversionMap{
		conversions: []conversion{},
	}

	for _, intStringMatch := range mapIntMatches {
		mapRow := toIntSlice(intStringMatch[1])
		conv := conversion{
			start:      mapRow[1],
			end:        mapRow[1] + mapRow[2] - 1,
			diff:       mapRow[1] - mapRow[0],
			startValue: mapRow[0],
		}
		out.conversions = append(out.conversions, conv)
	}

	return out
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

var seedsRe = `\nseeds:([^\n]+)\n`
var mpasRe = `map:\n([^A-Za-z]+)`

var example = `
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`
var input = `
`
