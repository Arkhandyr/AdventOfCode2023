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
	times, records := []int{}, []int{}

	for i := range timesStr {
		time, _ := strconv.Atoi(timesStr[i])
		times = append(times, time)

		record, _ := strconv.Atoi(recordsStr[i])
		records = append(records, record)
	}

	marginOfError := 1
	for i, raceTime := range times {
		wins := 0
		for j := 1; j < raceTime; j++ {
			speed, time, record := j, raceTime-j, records[i]

			if speed*time > record {
				wins++
				fmt.Printf("Could win pressing %d milliseconds, as boat would travel %d millimeters (%d * %d), and record is %d\n", speed, speed*time, speed, time, record)
			}
		}
		marginOfError *= wins
		fmt.Printf("There are %d ways of winning race %d\n", wins, i+1)
	}
	fmt.Printf("Margin of error: %d\n", marginOfError)
}
