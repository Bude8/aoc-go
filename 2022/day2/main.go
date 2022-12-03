package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculateScorePart1(opponentHand string, ourHand string) int {
	var score = 0
	switch true {
	case (opponentHand == "A" && ourHand == "X"):
		score = 4
	case (opponentHand == "A" && ourHand == "Y"):
		score = 8
	case (opponentHand == "A" && ourHand == "Z"):
		score = 3
	case (opponentHand == "B" && ourHand == "X"):
		score = 1
	case (opponentHand == "B" && ourHand == "Y"):
		score = 5
	case (opponentHand == "B" && ourHand == "Z"):
		score = 9
	case (opponentHand == "C" && ourHand == "X"):
		score = 7
	case (opponentHand == "C" && ourHand == "Y"):
		score = 2
	case (opponentHand == "C" && ourHand == "Z"):
		score = 6
	}
	return score
}

func calculateScorePart2(opponentHand string, outcome string) int {
	var score = 0
	switch true {
	case (opponentHand == "A" && outcome == "X"):
		score = 3
	case (opponentHand == "A" && outcome == "Y"):
		score = 4
	case (opponentHand == "A" && outcome == "Z"):
		score = 8
	case (opponentHand == "B" && outcome == "X"):
		score = 1
	case (opponentHand == "B" && outcome == "Y"):
		score = 5
	case (opponentHand == "B" && outcome == "Z"):
		score = 9
	case (opponentHand == "C" && outcome == "X"):
		score = 2
	case (opponentHand == "C" && outcome == "Y"):
		score = 6
	case (opponentHand == "C" && outcome == "Z"):
		score = 7
	}
	return score
}

func main() {
	// https://stackoverflow.com/questions/6141604/go-readline-string
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	// https://go.dev/tour/flowcontrol/12
	defer file.Close()

	var totalScorePart1 = 0
	var totalScorePart2 = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hands := strings.Split(scanner.Text(), " ")
		opponentHand, secondColumn := hands[0], hands[1]
		totalScorePart1 += calculateScorePart1(opponentHand, secondColumn)
		totalScorePart2 += calculateScorePart2(opponentHand, secondColumn)
	}
	fmt.Println("Part 1: ", totalScorePart1)
	fmt.Println("Part 2: ", totalScorePart2)
}
