package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calculateScorePart1(opponentHand string, ourHand string) int {
	// A, X = Rock
	// B, Y = Paper
	// C, Z = Scissors
	var score = 0
	if ourHand == "X" {
		score += 1
	} else if ourHand == "Y" {
		score += 2
	} else if ourHand == "Z" {
		score += 3
	}

	if opponentHand == "A" && ourHand == "X" ||
		opponentHand == "B" && ourHand == "Y" ||
		opponentHand == "C" && ourHand == "Z" {
		score += 3
	} else if opponentHand == "A" && ourHand == "Y" ||
		opponentHand == "B" && ourHand == "Z" ||
		opponentHand == "C" && ourHand == "X" {
		score += 6
	}
	return score
}

func calculateScorePart2(opponentHand string, outcome string) int {
	// X = Lose
	// Y = Draw
	// Z = Win
	var ourHand = ""
	if outcome == "X" {
		if opponentHand == "A" {
			ourHand = "Z"
		}
		if opponentHand == "B" {
			ourHand = "X"
		}
		if opponentHand == "C" {
			ourHand = "Y"
		}
	} else if outcome == "Y" {
		if opponentHand == "A" {
			ourHand = "X"
		}
		if opponentHand == "B" {
			ourHand = "Y"
		}
		if opponentHand == "C" {
			ourHand = "Z"
		}
	} else if outcome == "Z" {
		if opponentHand == "A" {
			ourHand = "Y"
		}
		if opponentHand == "B" {
			ourHand = "Z"
		}
		if opponentHand == "C" {
			ourHand = "X"
		}
	}
	return calculateScorePart1(opponentHand, ourHand)
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
