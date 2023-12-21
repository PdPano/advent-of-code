package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input1.txt")
	ss := strings.Split(strings.TrimSpace(string(file)), ",")
	part1(ss)
	part2(ss)
}

func part1(ss []string) {
	acc := 0
	for _, s := range ss {
		acc += hash(s)
	}
	fmt.Println("Part1:", acc)
}

type Lens struct {
	label       string
	boxId       int
	focalLength int
}

func part2(ss []string) {
	boxes := make([][]Lens, 256)

	for _, s := range ss {
		t := instructionType(s)
		switch t {
		case minus:
			m := parseMinus(s)
			deleteLabel(&boxes[m.boxId], m.label)
			break
		case equals:
			e := parseEquals(s)
			box := boxes[e.boxId]
			i, f := findLabel(&box, e.label)
			if f {
				box[i] = e
			} else {
				boxes[e.boxId] = append(box, e)
			}
			break
		}
	}
	acc := 0
	for i, b := range boxes {
		for j, l := range b {
			acc += (i + 1) * (j + 1) * l.focalLength
		}
	}
	fmt.Println("Part2:", acc)
}

func deleteLabel(a *[]Lens, lab string) {
	i, found := findLabel(a, lab)
	if !found {
		return
	}
	*a = append((*a)[:i], (*a)[i+1:]...)
}

func findLabel(a *[]Lens, lab string) (int, bool) {
	for i, l := range *a {
		if l.label == lab {
			return i, true
		}
	}
	return -1, false
}

type instType int

const (
	minus instType = iota
	equals
)

func instructionType(s string) instType {
	if strings.HasSuffix(s, "-") {
		return minus
	}
	return equals
}

func parseMinus(s string) Lens {
	label := s[:len(s)-1]
	return Lens{label: label, boxId: hash(label)}
}

func parseEquals(s string) Lens {
	parts := strings.Split(s, "=")
	label := parts[0]
	fl, _ := strconv.Atoi(parts[1])
	return Lens{label: label, boxId: hash(label), focalLength: fl}
}

func hash(s string) int {
	b := []byte(s)
	h := 0
	for _, c := range b {
		h += int(c)
		h *= 17
		h = h % 256
	}
	return h
}
