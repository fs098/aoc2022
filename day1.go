package main

import (
	"fmt"
	"sort"
)

func day1(filename string) {
	fmt.Println("Day 1")

	part1 := maxCals(filename)
	fmt.Printf("Part 1: %d\n", part1)

	part2 := maxCalsPart2(filename)
	fmt.Printf("Part 2: %d\n", part2)
}

func maxCals(filename string) int {
	var (
		data           = getLines(filename)
		maxCals        = 0
		currentElfCals = 0
	)

	if data[len(data)-1] != "" {
		data = append(data, "")
	}

	for _, v := range data {
		if v == "" {
			if currentElfCals > maxCals {
				maxCals = currentElfCals
			}
			currentElfCals = 0
			continue
		}
		currentElfCals += readInt(v)
	}

	return maxCals
}

func maxCalsPart2(filename string) int {
	var (
		data           = getLines(filename)
		totalCalsList  = []int{}
		currentElfCals = 0
	)

	if data[len(data)-1] != "" {
		data = append(data, "")
	}

	for _, v := range data {
		if v == "" {
			totalCalsList = append(totalCalsList, currentElfCals)
			currentElfCals = 0
			continue
		}
		currentElfCals += readInt(v)
	}

	sort.Slice(totalCalsList, func(i, j int) bool {
		return totalCalsList[i] > totalCalsList[j]
	})

	return totalCalsList[0] + totalCalsList[1] + totalCalsList[2]
}
