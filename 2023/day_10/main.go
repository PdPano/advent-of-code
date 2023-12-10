package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input1.txt")
	rows := strings.Split(strings.TrimSpace(string(file)), "\n")
	conns, sCoord := buildDiagram(rows)
	visited := map[Coordinate]struct{}{}
	toVisit := make([]Coordinate, 0)
	toVisit = append(toVisit, sCoord)
	for {
		curr := pop(&toVisit)
		visited[curr] = struct{}{}
		for _, neigh := range conns[curr] {
			_, ok := visited[neigh]
			if !ok {
				toVisit = append(toVisit, neigh)
			}
		}
		if len(toVisit) == 0 {
			break
		}
	}
    sScore:=0
    sConns:=conns[sCoord]
    lInMap:=slices.Contains(sConns, buildCoordinate(sCoord.i,sCoord.j-1))
    rInMap:=slices.Contains(sConns, buildCoordinate(sCoord.i,sCoord.j+1))
    uInMap:=slices.Contains(sConns, buildCoordinate(sCoord.i-1,sCoord.j))
    dInMap:=slices.Contains(sConns, buildCoordinate(sCoord.i+1,sCoord.j))
    if(lInMap&&uInMap){sScore=1}
    if(lInMap&&dInMap){sScore=-1}
    if(rInMap&&uInMap){sScore=-1}
    if(rInMap&&dInMap){sScore=1}
    if(uInMap&&dInMap){sScore=2}

	part1(visited)
	part2(sScore, rows, conns, visited, len(rows), len(rows[0]))
}

type Coordinate struct {
	i int
	j int
}

func buildCoordinate(i, j int) Coordinate {
	return Coordinate{i: i, j: j}
}

func part2(sScore int, rows []string, conns map[Coordinate]([]Coordinate), visited map[Coordinate]struct{}, lenY int, lenX int) {
	count := 0
	for i := 0; i < lenY; i++ {
		currIn := false
        score:=0
		for j := 0; j < lenX; j++ {
            if(score==2 || score==-2){
                score=0
                currIn=!currIn
            }
			// Is on border?
			curr := buildCoordinate(i, j)
			_, onBorder := visited[curr]
			if onBorder {
                score+=pieceScore(string(rows[i][j]), sScore)
			}

			if currIn && !onBorder{
				count++
			}
		}
	}
	fmt.Println(count)
}

func pieceScore(p string, sScore int) int{
    if(p=="S"){return sScore}
    if(p=="J" || p=="F"){return 1}
    if(p=="7" || p=="L"){return -1}
    if(p=="|"){return 2}
    return 0
}

func part1(visited map[Coordinate]struct{}) {
	fmt.Println(len(visited) / 2)
}

func pop(alist *[]Coordinate) Coordinate {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}

func buildDiagram(rows []string) (map[Coordinate]([]Coordinate), Coordinate) {
	var sCoord Coordinate
	m := map[Coordinate]([]Coordinate){}

	for i, row := range rows {
		for j, c := range row {
			if c == 'S' {
				sCoord = buildCoordinate(i, j)
			}
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
	return m, sCoord
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
	return use && (curr == '|' || curr == 'L' || curr == 'J' || curr == 'S')
}

func canConnectDown(curr rune, use bool) bool {
	return use && (curr == '|' || curr == 'F' || curr == '7' || curr == 'S')
}

func canConnectLeft(curr rune, use bool) bool {
	return use && (curr == '-' || curr == '7' || curr == 'J' || curr == 'S')
}

func canConnectRight(curr rune, use bool) bool {
	return use && (curr == '-' || curr == 'L' || curr == 'F' || curr == 'S')
}
