package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = `Time:        51     92     68     90
Distance:   222   2031   1126   1225`

func main() {
	lines := strings.Split(input, "\n")
	timesStr, recordsStr := strings.Fields(lines[0])[1:], strings.Fields(lines[1])[1:]
	raceTimeStr, recordStr := "", ""

	for i := range timesStr {
		raceTimeStr += timesStr[i]
		recordStr += recordsStr[i]
	}

	raceTime, _ := strconv.Atoi(raceTimeStr)
	record, _ := strconv.Atoi(recordStr)

	wins := 0
	firstWin := 0
	lastWin := 0
	for i := 1; i < raceTime; i++ {
		speed, time := i, raceTime-i

		if speed*time > record {
			firstWin = i
			fmt.Printf("Could win pressing %d milliseconds, as boat would travel %d millimeters (%d * %d), %d more than record\n", speed, speed*time, speed, time, speed*time-record)
			break
		}
	}

	for i := raceTime; i >= 1; i-- {
		speed, time := i, raceTime-i

		if speed*time > record {
			lastWin = i
			fmt.Printf("Could win pressing %d milliseconds, as boat would travel %d millimeters (%d * %d), %d more than record\n", speed, speed*time, speed, time, speed*time-record)
			break
		}
	}

	wins = lastWin - firstWin + 1
	fmt.Printf("There are %d ways of winning the race\n", wins)
}
