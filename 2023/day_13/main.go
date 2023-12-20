package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input1.txt")
	rows := strings.Split(strings.TrimSpace(string(file)), "\n")
	patterns := make([][]string, 1)
	patternInd := 0

	for _, row := range rows {
		if len(row) > 0 {
			patterns[patternInd] = append(patterns[patternInd], row)
		} else {
			patternInd += 1
			patterns = append(patterns, make([]string, 0))
		}
	}
	part1(patterns)
	part2(patterns)
}

func part1(patterns [][]string) {
	acc := 0
	for _, pattern := range patterns {
		acc += scorePattern(pattern)
	}
	fmt.Println("Part1:", acc)
}

func part2(patterns [][]string) {
	acc := 0
	for _, pattern := range patterns {
		origScore := scorePattern(pattern)
		acc += scoreSmudge(pattern, origScore)
	}
	fmt.Println("Part2:", acc)
}

func scoreSmudge(pattern []string, origScore int) int {
	for j := 0; j < len(pattern[0]); j++ {
		for i := 0; i < len(pattern); i++ {
			patternCopy := make([]string, len(pattern))
			copy(patternCopy, pattern)
			s := []byte(patternCopy[i])
			if s[j] == '.' {
				s[j] = '#'
			} else {
				s[j] = '.'
			}
			patternCopy[i] = string(s)
			newScore := scorePattern(patternCopy)
			if origScore != newScore && newScore > 0 {
				res := 0
				if (origScore % 100) != (newScore % 100) {
					res += newScore % 100
				}
				if (origScore / 100) != (newScore / 100) {
					res += (newScore / 100) * 100
				}
				return res
			}
		}
	}
	return origScore
}

func scorePattern(pattern []string) int {
	return 100*scorePatternForRow(pattern) + scorePatternForRow(transpose(pattern))
}

func transpose(rows []string) []string {
	nRows := len(rows[0])
	nCols := len(rows)
	transposedRows := make([]string, len(rows[0]))
	for i := 0; i < nRows; i++ {
		tmp := make([]string, 0)
		for j := 0; j < nCols; j++ {
			tmp = append(tmp, string(rows[j][i]))
		}
		transposedRows[i] = strings.Join(tmp, "")
	}
	return transposedRows
}

func scorePatternForRow(pattern []string) int {
	for i := 0; i < len(pattern)-1; i++ {
		head := pattern[:i+1]
		tail := make([]string, 0)
		for j := i + 1; j < len(pattern); j++ {
			tail = append(tail, pattern[j])
		}

		for i2, j := 0, len(tail)-1; i2 < j; i2, j = i2+1, j-1 {
			tail[i2], tail[j] = tail[j], tail[i2]
		}

		var a string
		var b string

		if len(head) > len(tail) {
			a = strings.Join(head, "")
			b = strings.Join(tail, "")
		} else {
			b = strings.Join(head, "")
			a = strings.Join(tail, "")
		}
		if strings.HasSuffix(a, b) {
			return i + 1
		}
	}
	return 0
}
