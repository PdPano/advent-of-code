package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input1.txt")
	sequences := strings.Split(strings.TrimSpace(string(file)), "\n")
	part1(sequences)
	part2(sequences)
}

func part1(sequences []string) {
	acc := 0
	for _, seq := range sequences {
		acc += extrapolate(seq, false)
	}
	fmt.Println(acc)
}

func part2(sequences []string) {
	acc := 0
	for _, seq := range sequences {
		acc += extrapolate(seq, true)
	}
	fmt.Println(acc)
}

func extrapolate(seq string, revert bool) int {
	numsAsStr := strings.Fields(seq)
	arr := make([]int, len(numsAsStr))
	for i, v := range numsAsStr {
			arr[i], _ = strconv.Atoi(v)
	}
    if(revert){
        slices.Reverse(arr)
    }

	for j := 0; j < len(arr); j++ {
		for i := 0; i < len(arr)-1-j; i++ {
			arr[i] = (arr[i+1] - arr[i])
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		arr[i+1] = arr[i+1] + arr[i]
	}
	return arr[len(arr)-1]
}
