package main

import (
	"fmt"
	"strings"
)

var input = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

var cardToValue = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var types = []string{}

func main() {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		hand := strings.Fields(line)[0]
		GetHandType(hand)
	}

}

func GetHandType(hand string) {
	labels := []rune{}
	for _, card := range hand {
		if !tools.contains(labels, card) {
			labels = append(labels, card)
		}
	}

	fmt.Println(labels)
}
