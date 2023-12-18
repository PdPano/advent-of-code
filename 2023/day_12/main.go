package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input1.txt")
	rows := strings.Split(strings.TrimSpace(string(file)), "\n")
	part1(rows)
	part2(rows)
}

func part1(rows []string) {
	acc := 0
	for _, r := range rows {
		s := strings.Split(r, " ")
		tmp := strings.Split(s[1], ",")
		consec := make([]int, len(tmp))
		for i := 0; i < len(consec); i++ {
			consec[i], _ = strconv.Atoi(tmp[i])
		}
		acc += computeForRow(s[0][0], []byte(s[0][1:]), consec[0], consec[1:], false, 0)
	}
	fmt.Println("Part1:", acc)
}

func part2(rows []string) {
	acc := 0
	for _, r := range rows {
		s := strings.Split(r, " ")
		tmp := strings.Split(s[1], ",")
		consec := make([]int, 5*len(tmp))
		for j := 0; j < 5; j++ {
			for i := 0; i < len(tmp); i++ {
				consec[i+j*(len(tmp))], _ = strconv.Atoi(tmp[i])
			}
		}
        x:=s[0]+"?"+s[0]+"?"+s[0]+"?"+s[0]+"?"+s[0]
        r:= computeForRow(x[0], []byte(x[1:]), consec[0], consec[1:], false, 0)
        fmt.Println(r)
        acc += r
	}
	fmt.Println("Part2:", acc)
}

func computeForRow(head byte, info []byte, currentRun int, consecutiveCounts []int, inRun bool, acc int) int {
	//fmt.Println(string(head), string(info), currentRun, consecutiveCounts, inRun, acc)
	if len(info) == 0 && (head == '.' || head == '?') && len(consecutiveCounts) == 0 && currentRun == 0 {
		return 1
	}

	if len(info) == 0 && (head == '#' || head == '?') && len(consecutiveCounts) == 0 && currentRun == 1 {
		return 1
	}

	if len(info) == 0 {
		return 0
	}

	if head == '?' {
		acc += computeForRow('#', info, currentRun, consecutiveCounts, inRun, 0)
		acc += computeForRow('.', info, currentRun, consecutiveCounts, inRun, 0)
		return acc
	}

	if head == '#' {
		if inRun {
			if currentRun == 0 {
				return 0
			}
			currentRun = currentRun - 1
		} else {
			if currentRun > 0 {
				inRun = true
				currentRun = currentRun - 1
			} else {
				return 0
			}
		}
	}

	if head == '.' {
		if inRun {
			if currentRun == 0 {
				inRun = false
				if len(consecutiveCounts) > 0 {
					currentRun = consecutiveCounts[0]
					consecutiveCounts = consecutiveCounts[1:]
				}
			} else {
				return 0
			}
		}
	}

	return acc + computeForRow(info[0], info[1:], currentRun, consecutiveCounts, inRun, acc)
}
