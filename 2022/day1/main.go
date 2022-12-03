package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// https://stackoverflow.com/questions/6141604/go-readline-string
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	// https://go.dev/tour/flowcontrol/12
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// https://go.dev/tour/moretypes/15
	elvesCalories := make([]int, 1)
	max := 0
	for scanner.Scan() {
		if line := scanner.Text(); len(line) > 0 {
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			elvesCalories[len(elvesCalories)-1] += calories
			if elvesCalories[len(elvesCalories)-1] > max {
				max = elvesCalories[len(elvesCalories)-1]
			}
		} else {
			elvesCalories = append(elvesCalories, 0)
		}
	}

	fmt.Println("Part 1:", max)

	// Sort the array and take the last 3 for part 2.
	sort.Sort(sort.Reverse(sort.IntSlice(elvesCalories)))

	// Does Go not have a `sum()` function?
	sumTop3Elves := 0
	for _, elfCalories := range elvesCalories[:3] {
		sumTop3Elves += elfCalories
	}
	fmt.Println("Part 2:", sumTop3Elves)
}
