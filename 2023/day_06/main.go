package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	{
		file, err := os.ReadFile("input1.txt")
		if err != nil {
			panic(err)
		}
		asString := string(file)
		data := strings.Split(asString, "\n")
		times := parseArray(data[0])
		distances := parseArray(data[1])
		part1(times, distances)
	}
	{
		file, err := os.ReadFile("input2.txt")
		if err != nil {
			panic(err)
		}
		asString := string(file)
		data := strings.Split(asString, "\n")
		times := parseArray(data[0])
		distances := parseArray(data[1])
		part2(times[0], distances[0])
	}
}

func parseArray(data string) []int {
	acc := make([]int, 0)
	numbers := strings.Split(strings.Split(data, ":")[1], " ")
	for _, v := range numbers {
		if v == "" {
			continue
		}
		num, _ := strconv.Atoi(v)
		acc = append(acc, num)
	}
	return acc
}

func part2(time int, distance int) {
	target := func(t int) bool {
		return targetFunc(t, time, distance)
	}

	lr := bisect(target, 0, time/2)
	rr := bisect(target, time/2+1, time)
    fmt.Println(rr-lr)
}

func bisect(f func(int) bool, l int, r int) int {
	fl := f(l)

	for {
		if r-l <= 1 {
			break
		}
		mid := l + (r-l)/2
		fm := f(mid)
		if fl != fm {
            r = mid
		} else { // fm != fr
            l = mid
		}
	}
	return l
}

func part1(times []int, distances []int) {
	acc := 1
	for i := 0; i < len(times); i++ {
		acc = acc * winningRacesCount(times[i], distances[i])
	}
	fmt.Println(acc)
}

func winningRacesCount(time int, distance int) int {
	acc := 0
	for t := 0; t < time; t++ {
		if targetFunc(t, time, distance) {
			acc++
		}
	}
	return acc
}

func targetFunc(t int, time int, distance int) bool {
	return (time-t)*t-distance > 0
}
