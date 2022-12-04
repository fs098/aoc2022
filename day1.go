package main

import (
	"fmt"
	"sort"
)

func day1(filename string) {
	data := getLines(filename)
	part1, part2 := maxCals(data)

	fmt.Println("Day 1")
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}

func maxCals(data []string) (int, int) {
	totalCalsList := []int{}
	currentElfCals := 0

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

	part1 := totalCalsList[0]
	part2 := totalCalsList[0] + totalCalsList[1] + totalCalsList[2]

	return part1, part2
}
