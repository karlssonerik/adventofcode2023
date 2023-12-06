package main

import "fmt"

func main() {
	part1()
	part2()
}

func part2() {
	for _, race := range inputRaces2 {
		waysToWin := calculateDistnace(race.time, race.distance)
		fmt.Println(waysToWin)
	}
}

func part1() {
	pow := 0
	for _, race := range inputRaces {
		waysToWin := calculateDistnace(race.time, race.distance)
		if pow == 0 {
			pow = waysToWin
		} else {
			pow *= waysToWin
		}
	}

	fmt.Println(pow)
}

func calculateDistnace(time int, winningDistance int) int {
	waysToWin := 0
	for hold := 0; hold < time; hold++ {
		distance := hold * (time - hold)
		if distance > winningDistance {
			waysToWin++
		}
	}

	return waysToWin
}

var example = `
Time:      7  15   30
Distance:  9  40  200
`
var input = `
`

type raceInfo struct {
	time     int
	distance int
}

var exampleRaces = []raceInfo{{7, 9}, {15, 40}, {30, 200}}
var exampleRaces2 = []raceInfo{{71530, 940200}}
var inputRaces = []raceInfo{{48, 296}, {93, 1928}, {85, 1236}, {95, 1391}}
var inputRaces2 = []raceInfo{{48938595, 296192812361391}}
