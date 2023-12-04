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
	schematic := strings.Split(asString, "\n")
	part1(schematic)
	part2(schematic)
}

type Interval struct {
	begin int
	end   int
	value int
}

func part2(schematic []string) {
    totalGearRatio := 0
	intervals := buildIntervals(schematic)
	for i, line := range schematic {
		for j, char := range line {
			if char == '*' {
				intervalsInRange := findIntervalsAround(i, j, intervals)
                if len(intervalsInRange) == 2 {
                    totalGearRatio += intervalsInRange[0].value * intervalsInRange[1].value
                }
			}
		}
	}
    fmt.Printf("Part 2: %d", totalGearRatio)
}

func findIntervalsAround(i int, j int, intervals [][]Interval) []Interval {
	matchedIntervals := map[Interval]bool{}
	for shiftI := -1; shiftI <= 1; shiftI++ {
		for shiftJ := -1; shiftJ <= 1; shiftJ++ {
			interval, found := peekInterval(i+shiftI, j+shiftJ, intervals)
			if found {
				matchedIntervals[interval] = true
			}
		}
	}
	var result []Interval
	for k := range matchedIntervals {
		result = append(result, k)
	}
	return result
}

func peekInterval(i int, j int, intervals [][]Interval) (Interval, bool) {
	if i >= 0 && i < len(intervals) {
		for _, interval := range intervals[i] {
			if interval.begin <= j && j <= interval.end {
				return interval, true
			}
		}
	}
	return Interval{}, false
}

func buildIntervals(schematic []string) [][]Interval {
	var intervals [][]Interval

	for i := 0; i < len(schematic)-1; i++ {
		intervals = append(intervals, make([]Interval, 0))
		currNumber := 0
		begin := -1
		for j := 0; j < len(schematic[i]); j++ {
			char := schematic[i][j]
			if isDigit(char) {
				v, _ := strconv.Atoi(string(char))
				currNumber = 10*currNumber + v
				if begin == -1 {
					begin = j
				}
			} else {
				if currNumber != 0 {
					intervals[i] = append(intervals[i], Interval{begin: begin, end: j - 1, value: currNumber})
				}
				currNumber = 0
				begin = -1
			}
		}
		if currNumber != 0 {
			intervals[i] = append(intervals[i], Interval{begin: begin, end: len(schematic[i]) - 1, value: currNumber})
		}
	}
	return intervals
}

func part1(schematic []string) {
	totalValue := 0

	for i := 0; i < len(schematic); i++ {
		currNumber := 0
		shouldUse := false
		for j := 0; j < len(schematic[i]); j++ {
			char := schematic[i][j]
			if isDigit(char) {
				v, _ := strconv.Atoi(string(char))
				currNumber = 10*currNumber + v
				shouldUse = shouldUse || hasSymbolAround(i, j, schematic)
			} else {
				if shouldUse {
					totalValue += currNumber
				}
				currNumber = 0
				shouldUse = false
			}
		}
		if shouldUse {
			totalValue += currNumber
		}
	}
	fmt.Printf("Part1: %d\n", totalValue)
}

func isSymbol(c byte) bool {
	return c != '.' && !isDigit(c)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func hasSymbolAround(i int, j int, schematic []string) bool {
	for shiftI := -1; shiftI <= 1; shiftI++ {
		for shiftJ := -1; shiftJ <= 1; shiftJ++ {
			if peek(i+shiftI, j+shiftJ, schematic) {
				return true
			}
		}
	}
	return false
}

func peek(i int, j int, schematic []string) bool {
	if i < 0 || j < 0 || i >= len(schematic) || j >= len(schematic[i]) {
		return false
	}
	return isSymbol(schematic[i][j])
}
