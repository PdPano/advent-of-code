package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("sample.txt")
	rows := strings.Split(strings.TrimSpace(string(file)), "\n")
	part1(rows)
}

type Coordinate struct {
	i int
	j int
}

func buildCoordinate(i, j int) Coordinate {
	return Coordinate{i: i, j: j}
}

func part1(rows []string) {
	fmt.Println(buildDiagram(rows))
}

func buildDiagram(rows []string) map[Coordinate]([]Coordinate) {
	m := map[Coordinate]([]Coordinate){}

	for i, row := range rows {
		for j, c := range row {
			if canConnectUp(c, true) && canConnectDown(peek(i-1, j, rows)) {
				appendToMap(&m, buildCoordinate(i, j), buildCoordinate(i-1, j))
			}
			if canConnectDown(c, true) && canConnectUp(peek(i+1, j, rows)) {
				appendToMap(&m, buildCoordinate(i, j), buildCoordinate(i+1, j))
			}
			if canConnectLeft(c, true) && canConnectRight(peek(i, j-1, rows)) {
				appendToMap(&m, buildCoordinate(i, j), buildCoordinate(i, j-1))
			}
			if canConnectRight(c, true) && canConnectLeft(peek(i, j+1, rows)) {
				appendToMap(&m, buildCoordinate(i, j), buildCoordinate(i, j+1))
			}
		}
	}
	return m
}

func appendToMap(m *map[Coordinate]([]Coordinate), from, to Coordinate) {
	v, ok := (*m)[from]
	if !ok {
		(*m)[from] = make([]Coordinate, 0)
	}
	(*m)[from] = append(v, to)
}

func peek(i, j int, rows []string) (rune, bool) {
	if i < 0 || i >= len(rows) || j < 0 || j >= len(rows[i]) {
		return '.', false
	}
	return rune(rows[i][j]), true
}

func canConnectUp(curr rune, use bool) bool {
	return use && (curr == '|' || curr == 'L' || curr == 'J')
}

func canConnectDown(curr rune, use bool) bool {
	return use && (curr == '|' || curr == 'F' || curr == '7')
}

func canConnectLeft(curr rune, use bool) bool {
	return use && (curr == '-' || curr == '7' || curr == 'J')
}

func canConnectRight(curr rune, use bool) bool {
	return use && (curr == '-' || curr == 'L' || curr == 'F')
}
