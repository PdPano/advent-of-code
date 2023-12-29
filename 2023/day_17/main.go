package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input1.txt")
	ss := strings.Split(strings.TrimSpace(string(file)), "\n")
	heatMap := make([][]int, len(ss))

	for i, r := range ss {
		heatMap[i] = make([]int, len(r))
		for j, n := range r {
			num, _ := strconv.Atoi(string(n))
			heatMap[i][j] = num
		}
	}
	part1(heatMap)
	part2(heatMap)
}

func part1(heatMap [][]int) {
	start := image.Point{X: 0, Y: 0}
	goal := image.Point{X: len(heatMap[0]) - 1, Y: len(heatMap) - 1}
	minPath, minLoss := astar(start, goal, heatMap, 0, 3)
	overlayPath(minPath, heatMap)
	fmt.Println("Part1:", minLoss)
}

func part2(heatMap [][]int) {
	start := image.Point{X: 0, Y: 0}
	goal := image.Point{X: len(heatMap[0]) - 1, Y: len(heatMap) - 1}
	minPath, minLoss := astar(start, goal, heatMap, 4, 10)
	overlayPath(minPath, heatMap)
	fmt.Println("Part2:", minLoss)
}

type direction int

const (
	d direction = iota //0
	u                  //1
	l                  //2
	r                  //3
)

type pointMetadata struct {
	dirOnPosition    direction
	consecutiveSteps int
}

type pointWithMetadata struct {
	point    image.Point
	metadata pointMetadata
}

func astar(start, goal image.Point, heatMap [][]int, minConsecutive, maxConsecutive int) ([]pointWithMetadata, int) {
	openSet := map[pointWithMetadata]struct{}{}
	openSet[pointWithMetadata{point: start, metadata: pointMetadata{dirOnPosition: r}}] = struct{}{}

	cameFrom := map[pointWithMetadata](pointWithMetadata){}

	gScore := map[pointWithMetadata]int{}
	gScore[pointWithMetadata{point: start}] = 0

	heuristic := func(n image.Point) int {
		return ((goal.X - n.X) + (goal.Y - n.Y))
	}

	fScore := map[image.Point]int{}
	fScore[start] = heuristic(start)

	minOpenKey := func() pointWithMetadata {
		minVal := math.MaxInt
		minK := pointWithMetadata{}
		found := false
		for k := range openSet {
			s, ok := fScore[k.point]
			if ok && s < minVal {
				minVal = s
				minK = k
				found = true
			}
		}
		if !found {
			panic("Should not happen")
		}
		delete(openSet, minK)
		return minK
	}

	shifts := map[direction](image.Point){
		u: image.Point{X: 0, Y: -1},
		d: image.Point{X: 0, Y: 1},
		l: image.Point{X: -1, Y: 0},
		r: image.Point{X: 1, Y: 0},
	}
	minHeatLoss := math.MaxInt
	minPath := make([](pointWithMetadata), 0)

	for len(openSet) > 0 {
		currentPointWithMetadata := minOpenKey()
		if currentPointWithMetadata.point == goal {
			currentHeatLoss := gScore[currentPointWithMetadata]
			if currentHeatLoss < minHeatLoss {
				minPath = computePath(cameFrom, currentPointWithMetadata)
				minHeatLoss = currentHeatLoss
			}
		}

		getNeighbourMetadata := func(n image.Point) pointMetadata {
			dir := computeDir(currentPointWithMetadata.point, n)
			if dir == currentPointWithMetadata.metadata.dirOnPosition {
				return pointMetadata{dirOnPosition: dir, consecutiveSteps: currentPointWithMetadata.metadata.consecutiveSteps + 1}
			}
			return pointMetadata{dirOnPosition: dir, consecutiveSteps: 1}
		}

		getNeighboursForCurrent := func() [](image.Point) {
			neighbours := make([](image.Point), 0)
			for _, nDir := range validNext[currentPointWithMetadata.metadata.dirOnPosition] {
				if nDir == currentPointWithMetadata.metadata.dirOnPosition && currentPointWithMetadata.metadata.consecutiveSteps >= maxConsecutive {
					continue
				}
				if nDir != currentPointWithMetadata.metadata.dirOnPosition && currentPointWithMetadata.metadata.consecutiveSteps < minConsecutive {
					continue
				}
				n := currentPointWithMetadata.point.Add(shifts[nDir])
				if n.X < 0 || n.X >= len(heatMap[0]) || n.Y < 0 || n.Y >= len(heatMap) {
					continue
				}
				neighbours = append(neighbours, n)
			}
			return neighbours
		}

		neighbours := getNeighboursForCurrent()
		for _, neighbour := range neighbours {
			tentative_gscore := gScore[currentPointWithMetadata] + heatMap[neighbour.Y][neighbour.X]
			neighbourWithMetadata := pointWithMetadata{point: neighbour, metadata: getNeighbourMetadata(neighbour)}

			currNeighbourScore, ok := gScore[neighbourWithMetadata]
			if !ok {
				currNeighbourScore = math.MaxInt
			}

			if tentative_gscore < currNeighbourScore {
				cameFrom[neighbourWithMetadata] = currentPointWithMetadata
				gScore[neighbourWithMetadata] = tentative_gscore
				fScore[neighbour] = tentative_gscore + heuristic(neighbour)
				openSet[neighbourWithMetadata] = struct{}{}
			}
		}
	}
	return minPath, minHeatLoss
}

func computeDir(orig, dest image.Point) direction {
	if dest.X > orig.X {
		return r
	}
	if dest.X < orig.X {
		return l
	}
	if dest.Y < orig.Y {
		return u
	}
	return d
}

func computePath(cameFrom map[pointWithMetadata](pointWithMetadata), current pointWithMetadata) [](pointWithMetadata) {
	path := make([]pointWithMetadata, 0)
	for {
		path = append(path, current)
		currentWithMeta, ok := cameFrom[current]
		if !ok {
			break
		}
		current = currentWithMeta
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

var dirToString = map[direction]string{
	u: "^",
	d: "v",
	l: "<",
	r: ">",
}

var validNext = map[direction]([]direction){
	u: []direction{u, l, r},
	d: []direction{d, l, r},
	r: []direction{u, r, d},
	l: []direction{u, d, r},
}

func overlayPath(path []pointWithMetadata, heatMap [][]int) {
	for i, r := range heatMap {
		row := make([]string, len(heatMap[0]))
		for j, v := range r {
			row[j] = strconv.Itoa(v)
			for _, n := range path {
				if n.point.X == j && n.point.Y == i {
					row[j] = dirToString[n.metadata.dirOnPosition]
				}
			}
		}
		fmt.Println(strings.Join(row, ""))
	}
}
