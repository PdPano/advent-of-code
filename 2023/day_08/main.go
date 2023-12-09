package main

import (
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	left  string
	right string
}


func main() {
	file, _ := os.ReadFile("input1.txt")
	data := strings.Split(string(file), "\n")
	instructions := strings.Split(data[0], "")
	var dirs = map[string]Pair{}

	for _, line := range data[2:] {
		if len(line) == 0 {
			continue
		}
		var source string
		dir := Pair{}
		fmt.Sscanf(line, "%3s = (%3s, %3s)", &source, &dir.left, &dir.right)
		dirs[source] = dir
	}
	//part1(instructions, dirs)
	part2(instructions, dirs)
}

func navigate(dir Pair, instruction string) string {
	if instruction == "L" {
		return dir.left
	} else {
		return dir.right
	}

}

func part1(instructions []string, dirs map[string]Pair) {
	steps := 0
	pos := 0
	curr := "AAA"
	for {
		steps++
		curr = navigate(dirs[curr], instructions[pos])
		if curr == "ZZZ" {
			break
		}
		pos = (pos + 1) % len(instructions)
	}
	fmt.Println(steps)
}

func part2(instructions []string, dirs map[string]Pair) {
	currs := make([]string, 0)
	for k := range dirs {
		if k[2] == 'A' {
			currs = append(currs, k)
		}
	}
    acc := 1
    offset := 0
    for _,k:=range currs{
        steps, off :=stepsUntilZ(instructions, dirs, k)
        acc = lcm(steps-off, acc)
        offset = off
    }
    acc += offset
    fmt.Print("Final: ", acc, "\n")
}

type InstructionLocation struct {
    pos int
    loc string
}

func stepsUntilZ(instructions []string, dirs map[string]Pair, curr string)(int,int){
	pos := 0
    steps := 0
    for {
        steps++
        curr = navigate(dirs[curr], instructions[pos])
		pos = (pos + 1) % len(instructions)

        if(curr[2]=='Z'){
            return steps,pos
        }
    }
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
