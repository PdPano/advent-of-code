package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
    acc:=0
    totalPower:=0

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

        game := parseGame(string(line))
        if (isPossible(game)) {
            acc += game.gameId
        }
        totalPower += game.result.red * game.result.blue * game.result.green
    }
    fmt.Println(acc)
    fmt.Println(totalPower)
    
}

func isPossible(game ParsedGame) bool {
    return game.result.red <=12 && game.result.green <=13 && game.result.blue <= 14
}

type GameResult struct {
	red   int
	blue  int
	green int
}

type ParsedGame struct {
	gameId int
	result GameResult
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func composeGames(game1 GameResult, game2 GameResult) GameResult {
	var composed GameResult

	composed.red = max(game1.red, game2.red)
	composed.blue = max(game1.blue, game2.blue)
	composed.green = max(game1.green, game2.green)

	return composed
}

func parseGame(game string) ParsedGame {
	parts := strings.Split(game, ":")
	gameId, err := strconv.Atoi(parts[0][5:])
	if err != nil {
		panic(err)
	}
	gameDraws := strings.Split(parts[1], ";")

	var gameResult GameResult
	for _, draw := range gameDraws {
		gameResult = composeGames(gameResult, parseGameDraw(draw))
	}
	return ParsedGame{gameId: gameId, result: gameResult}
}

func parseGameDraw(draw string) GameResult {
	var result GameResult
	cubes := strings.Split(draw, ",")

	for _, c := range cubes {
		ci := strings.Split(strings.TrimSpace(c), " ")

		count, _ := strconv.Atoi(ci[0])
		color := ci[1]

		if color == "red" {
			result.red = count
		}
		if color == "green" {
			result.green = count
		}
		if color == "blue" {
			result.blue = count
		}
	}

	return result
}
