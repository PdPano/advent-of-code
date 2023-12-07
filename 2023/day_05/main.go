package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("input1.txt")
	if err != nil {
		panic(err)
	}
	asString := string(file)
	almanac := strings.Split(asString, "\n")
	part1(almanac)
	part2(almanac)
}

type RangeMap struct {
	begin int
	end   int
	shift int
}

func isInRange(v int, rangeMap RangeMap) bool {
	return rangeMap.begin <= v && v <= rangeMap.end
}

func maybeConvertUsingRange(v int, rangeMap RangeMap) (int, bool) {
	if isInRange(v, rangeMap) {
		return v + rangeMap.shift, true
	}
	return v, false
}

func part2(almanac []string) {
	fmt.Println("Part 2")
	seedsRaw := numberArrayFromString(strings.Split(almanac[0], ":")[1])
	seeds := make([]RangeMap, len(seedsRaw)/2)
	for i := 0; i < len(seedsRaw)/2; i++ {
		seeds[i] = RangeMap{begin: seedsRaw[2*i], end: seedsRaw[2*i] + seedsRaw[2*i+1] - 1}
	}
	intervals := buildMap(almanac[1:])
	for _, layer := range intervals {
		seeds = intersectAll(seeds, layer)
	}
    sort.Slice(seeds, func(i, j int) bool {
        return seeds[i].begin < seeds[j].begin
    })
	fmt.Println(seeds[0].begin)
}

func countSeeds(seeds []RangeMap) int {
	acc := 0
	for _, s := range seeds {
		acc += s.end - s.begin + 1
	}
	return acc
}

func intersectAll(seedsRange []RangeMap, layerMap []RangeMap) []RangeMap {
	res := make([]RangeMap, 0)
	for _, seedRange := range seedsRange {
		res = append(res, intersectOne(seedRange, layerMap)...)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].begin < res[j].begin
	})
	return res
}

func intersectOne(seedRange RangeMap, layerMap []RangeMap) []RangeMap {
	res := make([]RangeMap, 0)
	for _, interval := range layerMap {
        // No intersection yet
		if seedRange.begin > interval.end {
			continue
		}

		// All remaining intervals are to the right
		if seedRange.end < interval.begin {
            res = append(res, seedRange)
            seedRange.begin=seedRange.end+1;
            break
		}


		// seed range is completely contained in interval
		if isInRange(seedRange.begin, interval) && isInRange(seedRange.end, interval) {
			res = append(res, RangeMap{begin: seedRange.begin + interval.shift, end: seedRange.end + interval.shift})
            seedRange.begin=seedRange.end+1;
			break // we are done with this seed range
		}

		// chomping from the left
		if isInRange(seedRange.begin, interval) && !isInRange(seedRange.end, interval) {
			res = append(res, RangeMap{begin: seedRange.begin + interval.shift, end: interval.end + interval.shift}) // Right part is mapped, the rest continues
			seedRange.begin = interval.end + 1
			continue
		}

		// chomping from the center (and we checked all other cases above)
		if !isInRange(seedRange.begin, interval) && !isInRange(seedRange.end, interval) {
			res = append(res, RangeMap{begin: seedRange.begin, end: interval.begin - 1})                            // Append left part
			res = append(res, RangeMap{begin: interval.begin + interval.shift, end: interval.end + interval.shift}) // Chomp
			seedRange.begin = interval.end + 1
			continue
		}

		// chomping from the right
		if !isInRange(seedRange.begin, interval) && isInRange(seedRange.end, interval) {
			res = append(res, RangeMap{begin: seedRange.begin, end: interval.begin - 1})                             // left part, not mapped
			res = append(res, RangeMap{begin: interval.begin + interval.shift, end: seedRange.end + interval.shift}) // all the rest, mapped
            seedRange.begin=seedRange.end+1;
			break                                                                                                    // we are done with this seed range
		}
	}
    if(seedRange.begin<=seedRange.end){
        res = append(res, seedRange)
    }
	return res
}

func part1(almanac []string) {
	fmt.Println("Part 1")
	seeds := numberArrayFromString(strings.Split(almanac[0], ":")[1])
	intervals := buildMap(almanac[1:])
	minLocation := math.MaxInt
	for _, seed := range seeds {
		currValue := seed
		for _, intervalList := range intervals {
			currValue = useMap(currValue, intervalList)
		}
		if currValue < minLocation {
			minLocation = currValue
		}
	}
	fmt.Println(minLocation)
}

func useMap(value int, intervals []RangeMap) int {
	for _, m := range intervals {
		v, converted := maybeConvertUsingRange(value, m)
		if converted {
			return v
		}
	}
	return value
}

func buildMap(almanac []string) [][]RangeMap {
	intervals := make([][]RangeMap, 0)
	for _, line := range almanac {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, ":") {
			intervals = append(intervals, make([]RangeMap, 0))
			continue
		}
		rangeData := numberArrayFromString(line)
		ref := intervals[len(intervals)-1]
		ref = append(ref, RangeMap{begin: rangeData[1], end: rangeData[1] + rangeData[2] - 1, shift: rangeData[0] - rangeData[1]})
		sort.Slice(ref, func(i, j int) bool {
			return ref[i].begin < ref[j].begin
		})
		intervals[len(intervals)-1] = ref
	}
	return intervals
}

func numberArrayFromString(row string) []int {
	arr := strings.Split(strings.TrimSpace(row), " ")
	res := make([]int, len(arr))
	for i, v := range arr {
		c, _ := strconv.Atoi(v)
		res[i] = c
	}
	return res
}
