package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input1.txt")
	if err != nil {
		panic(err)
	}
	asString := string(file)
	cards := strings.Split(asString, "\n")
	cards = cards[:len(cards)-1]
	part1(cards)
	part2(cards)
}

func part2(cards []string) {
	cardCounts := make([]int, len(cards))
    for i:=0;i<len(cardCounts);i++ {
        cardCounts[i] = 1
    }
	for i, card := range cards {
		numMatches := computeNumberOfMatches(card)
		for j := 0; j < numMatches; j++ {
			if i+j+1 < len(cards) {
				cardCounts[i+j+1] += cardCounts[i]
			}
		}
	}
	acc := 0
	for _, v := range cardCounts {
		acc += v
	}
	fmt.Println(acc)
}

func part1(cards []string) {
	totalScore := 0

	for _, card := range cards {
		totalScore += computeNumberOfMatches(card)
	}
	fmt.Println(totalScore)
}

func computeCardScore(card string) int {
	numMatches := computeNumberOfMatches(card)
	if numMatches == 0 {
		return 0
	} else {
		acc := 1
		for i := 0; i < numMatches-1; i++ {
			acc *= 2
		}
		return acc
	}
}

func computeNumberOfMatches(card string) int {
	parts := strings.Split(strings.Split(card, ":")[1], "|")
	numbersIHave := parseNumbers(parts[0])
	winningNumbers := parseNumbers(parts[1])

	return computeNumberOfMatchingElementsInLists(numbersIHave, winningNumbers)
}

func computeNumberOfMatchingElementsInLists(a []int, b []int) int {
	acc := 0
	for _, ai := range a {
		for _, bi := range b {
			if ai == bi {
				acc += 1
				break
			}
		}
	}
	return acc
}

func parseNumbers(numbersList string) []int {
	var acc []int
	for {
		if len(numbersList) < 3 {
			break
		}
		n, _ := strconv.Atoi(strings.TrimSpace(numbersList[0:3]))
		numbersList = numbersList[3:]
		acc = append(acc, n)
	}
	return acc
}
