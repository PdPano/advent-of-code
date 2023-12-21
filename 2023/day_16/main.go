package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input1.txt")
	ss := strings.Split(strings.TrimSpace(string(file)), "\n")
	part1(ss)
	part2(ss)
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

type Flow struct {
	dir direction
	i   int
	j   int
}

func walk(f Flow, d direction) Flow {
	switch d {
	case up:
		return Flow{i: f.i - 1, j: f.j, dir: d}
	case down:
		return Flow{i: f.i + 1, j: f.j, dir: d}
	case left:
		return Flow{i: f.i, j: f.j - 1, dir: d}
	case right:
		return Flow{i: f.i, j: f.j + 1, dir: d}
	}
	panic("Got unexpected dir")
}

func inBounds(f Flow, maxI, maxJ int) bool {
	return f.i >= 0 && f.i < maxI && f.j >= 0 && f.j < maxJ
}

func part1(mirrors []string) {
	s := scoreFrom(mirrors, Flow{i: 0, j: 0, dir: right})
	fmt.Println("Part1", s)
}

func part2(mirrors []string) {
	mScore := -1
	for i := 0; i < len(mirrors); i++ {
		fl := scoreFrom(mirrors, Flow{i: i, j: 0, dir: right})
		fr := scoreFrom(mirrors, Flow{i: i, j: len(mirrors[0]) - 1, dir: left})
		if fl > mScore {
			mScore = fl
		}
		if fr > mScore {
			mScore = fr
		}
	}
	for j := 0; j < len(mirrors[0]); j++ {
		fu := scoreFrom(mirrors, Flow{i: 0, j: j, dir: down})
		fd := scoreFrom(mirrors, Flow{i: len(mirrors) - 1, j: j, dir: up})
		if fu > mScore {
			mScore = fu
		}
		if fd > mScore {
			mScore = fd
		}
	}
	fmt.Println("Part2:", mScore)
}

func scoreFrom(mirrors []string, initialFlow Flow) int {
	visited := map[Flow]struct{}{}
	toVisit := make([]Flow, 0)
	toVisit = append(toVisit, initialFlow)

	for {
		head := toVisit[len(toVisit)-1]
		toVisit = toVisit[:len(toVisit)-1]

		_, seen := visited[head]

		if !seen {
			visited[head] = struct{}{}
			toVisit = append(toVisit, getNext(mirrors, head)...)
		}

		if len(toVisit) == 0 {
			break
		}
	}
	acc := map[Flow]struct{}{}
	for k := range visited {
		acc[Flow{i: k.i, j: k.j}] = struct{}{}
	}
	return len(acc)
}

type Dir struct {
	m byte
	d direction
}

var nextDirs = map[Dir]([]direction){
	Dir{m: '.', d: up}:    []direction{up},
	Dir{m: '.', d: down}:  []direction{down},
	Dir{m: '.', d: left}:  []direction{left},
	Dir{m: '.', d: right}: []direction{right},

	Dir{m: '/', d: up}:    []direction{right},
	Dir{m: '/', d: down}:  []direction{left},
	Dir{m: '/', d: left}:  []direction{down},
	Dir{m: '/', d: right}: []direction{up},

	Dir{m: '\\', d: up}:    []direction{left},
	Dir{m: '\\', d: down}:  []direction{right},
	Dir{m: '\\', d: left}:  []direction{up},
	Dir{m: '\\', d: right}: []direction{down},

	Dir{m: '|', d: up}:    []direction{up},
	Dir{m: '|', d: down}:  []direction{down},
	Dir{m: '|', d: left}:  []direction{up, down},
	Dir{m: '|', d: right}: []direction{up, down},

	Dir{m: '-', d: up}:    []direction{left, right},
	Dir{m: '-', d: down}:  []direction{left, right},
	Dir{m: '-', d: left}:  []direction{left},
	Dir{m: '-', d: right}: []direction{right},
}

func getNext(mirrors []string, currFlow Flow) []Flow {
	m := mirrors[currFlow.i][currFlow.j]
	n := make([]Flow, 0)

	for _, dir := range nextDirs[Dir{m: m, d: currFlow.dir}] {
		maybeNext := walk(currFlow, dir)
		if inBounds(maybeNext, len(mirrors), len(mirrors[0])) {
			n = append(n, maybeNext)
		}
	}

	return n
}
