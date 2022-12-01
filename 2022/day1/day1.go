package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

// Running total
func part1() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var groupCalories int = 0
	var largestCalorieGroup int = 0

	scanner := bufio.NewScanner(file)
	var calorie string

	for scanner.Scan() {
		calorie = scanner.Text()

		if calorie == "" {
			groupCalories = 0
		}

		if i, err := strconv.Atoi(calorie); err == nil {
			groupCalories += i
		}

		if largestCalorieGroup < groupCalories {
			largestCalorieGroup = groupCalories
		}
	}

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	return largestCalorieGroup
}

// Slice of slices
func part2() int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var groupCalories int = 0
	var largestCalorieGroup int = 0
	var secondLargestCalorieGroup int = 0
	var thirdLargestCalorieGroup int = 0
	var sumOfThreeLargestCalorieGroups int = 0

	scanner := bufio.NewScanner(file)
	var calorie string

	for scanner.Scan() {
		calorie = scanner.Text()

		if calorie == "" {
			groupCalories = 0
		}

		if i, err := strconv.Atoi(calorie); err == nil {
			groupCalories += i
		}

		if secondLargestCalorieGroup > groupCalories && thirdLargestCalorieGroup < groupCalories {
			thirdLargestCalorieGroup = groupCalories
		}

		if largestCalorieGroup > groupCalories && secondLargestCalorieGroup < groupCalories {
			thirdLargestCalorieGroup = secondLargestCalorieGroup
			secondLargestCalorieGroup = groupCalories
		}

		if largestCalorieGroup < groupCalories {
			thirdLargestCalorieGroup = secondLargestCalorieGroup
			secondLargestCalorieGroup = largestCalorieGroup
			largestCalorieGroup = groupCalories
		}
	}

	if err != nil {
		log.Fatalf("error: %s", err)
	}

	sumOfThreeLargestCalorieGroups = largestCalorieGroup + secondLargestCalorieGroup + thirdLargestCalorieGroup
	return sumOfThreeLargestCalorieGroups
}
