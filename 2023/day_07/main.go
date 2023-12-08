package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input1.txt")
	if err != nil {
		panic(err)
	}
	asString := string(file)
	data := strings.Split(asString, "\n")
	part1(data)
	part2(data)
}

func part1(hands []string) {
	parsedHands := make([]Hand, 0)
	for _, hand := range hands {
		if len(hand) == 0 {
			continue
		}
		parsedHands = append(parsedHands, parseHand(hand, computeKind))
	}
	sort.Slice(parsedHands, func(i, j int) bool {
		return parsedHands[i].isLess(parsedHands[j])
	})
	acc := 0
	for i, v := range parsedHands {
		acc += (i + 1) * v.bid
	}
	fmt.Println(acc)
}

func part2(hands []string) {
	cardMap["J"] = 0 // Downgrade the score
	parsedHands := make([]Hand, 0)
	for _, hand := range hands {
		if len(hand) == 0 {
			continue
		}
		parsedHands = append(parsedHands, parseHand(hand, computeKindWithJokers))
	}
	sort.Slice(parsedHands, func(i, j int) bool {
		return parsedHands[i].isLess(parsedHands[j])
	})
	acc := 0
	for i, v := range parsedHands {
		acc += (i + 1) * v.bid
	}
	fmt.Println(acc)
}

type handKind int

const (
	cardHigh handKind = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

var cardMap = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

type Hand struct {
	id   int
	kind handKind
	bid  int
}

func parseHand(data string, kindComputer func(string) handKind) Hand {
	d := strings.Split(data, " ")
	hand := d[0]
	bid, _ := strconv.Atoi(d[1])
	return Hand{id: computeId(hand), kind: kindComputer(hand), bid: bid}
}

func computeId(hand string) int {
	acc := 0
	for i, c := range hand {
		acc += intPow(14, len(hand)-1-i) * cardMap[string(c)]
	}
	return acc
}

func intPow(i int, n int) int {
	acc := 1
	for j := 0; j < n; j++ {
		acc *= i
	}
	return acc
}

func computeKindWithJokers(hand string) handKind {
	if !strings.Contains(hand, "J") {
		return computeKind(hand)
	}

	maxKind := cardHigh
	for k := range cardMap {
		score := computeKind(strings.ReplaceAll(hand, "J", k))
		if score > maxKind {
			maxKind = score
		}
	}
	return maxKind
}

func computeKind(hand string) handKind {
	var grouped = map[string]int{}
	for _, c := range hand {
		grouped[string(c)] = grouped[string(c)] + 1
	}
	if len(grouped) == 5 {
		return cardHigh
	}
	if len(grouped) == 4 {
		return onePair
	}
	if len(grouped) == 3 {
		for _, count := range grouped {
			if count == 2 {
				return twoPair
			}
			if count == 3 {
				return threeOfAKind
			}
		}
	}
	if len(grouped) == 2 {
		for _, count := range grouped {
			if count == 4 || count == 1 {
				return fourOfAKind
			}
			if count == 3 || count == 2 {
				return fullHouse
			}
		}
	}
	if len(grouped) == 1 {
		return fiveOfAKind
	}
	return cardHigh
}

func (re Hand) isLess(other Hand) bool {
	return (re.kind < other.kind) || (re.kind == other.kind && re.id < other.id)
}
