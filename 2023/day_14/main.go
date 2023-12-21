package main

import (
	"fmt"
	"os"
	"strings"
)

func getInput() [][]byte {
	file, _ := os.ReadFile("sample.txt")
	srows := strings.Split(strings.TrimSpace(string(file)), "\n")
	rows := make([][]byte, len(srows))
	for i := 0; i < len(srows); i++ {
		rows[i] = []byte(srows[i])
	}
    return rows
}

func main() {
    rows:=getInput()
	part1(rows)
    rows=getInput()
    part2(rows)
}

func part1(rows [][]byte) {
    bubbleNorth(&rows)
    fmt.Println("Part1: ",scoreMap(rows))
}

func mapStr(rows [][]byte) string {
    s:=make([]string, len(rows))
    for i, r:= range rows {
        s[i]=string(r)
    }
    return strings.Join(s, "")
}

func scoreMap(rows [][]byte) int{
	acc := 0
	for i := 0; i < len(rows); i++ {
		weight := len(rows) - i
		for col := 0; col < len(rows[0]); col++ {
			if rows[i][col] == 'O' {
				acc += weight
			}
		}
	}
    return acc
}

type Step struct {
    step int
}

func part2(rows [][]byte){
    seen:=map[string]Step{}
    cycleLength:=1
    currentStep:=-1

    for i:=0;i<1000000000;i++ {
        ms:=mapStr(rows)
        step, ok := seen[ms]
        if(ok){
            fmt.Println("Saw configuration on step", step.step)
            fmt.Println("Cycle length is", i-step.step)
            cycleLength=i-step.step
            currentStep=i
            break
        }else{
            seen[ms]=Step{step: i}
        }
        cycle(&rows)
    }
    fmt.Println(currentStep)
    fmt.Println(cycleLength)
    remaining := (1000000000-currentStep)%cycleLength
    for i:=0;i<remaining;i++{cycle(&rows)}
    fmt.Println(scoreMap(rows))
}

func cycle(rows *[][]byte){
    bubbleNorth(rows)
    bubbleWest(rows)
    bubbleSouth(rows)
    bubbleEast(rows)
}

func bubbleNorth(rows *[][]byte){
	for col := 0; col < len((*rows)[0]); col++ {
		for i := 0; i < len(*rows); i++ {
			for j := i; j > 0; j-- {
				if (*rows)[j][col] == 'O' && (*rows)[j-1][col] == '.' {
					(*rows)[j][col], (*rows)[j-1][col] = (*rows)[j-1][col], (*rows)[j][col]
				}
			}
		}
	}
}

func bubbleSouth(rows *[][]byte){
	for col := 0; col < len((*rows)[0]); col++ {
		for i := len(*rows)-1; i >=0; i-- {
			for j := i+1; j < len(*rows); j++ {
				if (*rows)[j-1][col] == 'O' && (*rows)[j][col] == '.' {
					(*rows)[j][col], (*rows)[j-1][col] = (*rows)[j-1][col], (*rows)[j][col]
				}
			}
		}
	}
}

func bubbleWest(rows *[][]byte){
	for r := 0; r < len(*rows); r++ {
		for i := 0; i < len((*rows)[0]); i++ {
			for j := i; j > 0; j-- {
				if (*rows)[r][j] == 'O' && (*rows)[r][j-1] == '.' {
					(*rows)[r][j], (*rows)[r][j-1] = (*rows)[r][j-1], (*rows)[r][j]
				}
			}
		}
	}
}

func bubbleEast(rows *[][]byte){
	for r := 0; r < len(*rows); r++ {
		for i := len((*rows)[0]); i >=0; i-- {
			for j := i+1; j < len((*rows)[0]); j++ {
				if (*rows)[r][j-1] == 'O' && (*rows)[r][j] == '.' {
					(*rows)[r][j], (*rows)[r][j-1] = (*rows)[r][j-1], (*rows)[r][j]
				}
			}
		}
	}
}

func printMap(rows [][]byte) {
	for _, r := range rows {
		fmt.Println(string(r))
	}
}
