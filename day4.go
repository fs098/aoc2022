package main

import (
	"fmt"
	"log"
)

func day4(filename string) {
	data := getLines(filename)
	part1, part2 := totalOverlap(data)

	fmt.Println("Day 4")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func totalOverlap(data []string) (int, int) {
	part1 := 0
	part2 := 0

	for _, v := range data {
		contains, overlaps := existsOverlap(v)
		if contains {
			part1++
		}
		if overlaps {
			part2++
		}
	}
	return part1, part2
}

func existsOverlap(s string) (bool, bool) {
	var (
		start1, end1 int
		start2, end2 int

		contains = false
		overlaps = false
	)

	_, err := fmt.Sscanf(s, "%d-%d,%d-%d", &start1, &end1, &start2, &end2)
	if err != nil {
		log.Fatal(err)
	}

	if start1 >= start2 && end1 <= end2 {
		contains = true
	}
	if start2 >= start1 && end2 <= end1 {
		contains = true
	}

	if (start1 >= start2 && start1 <= end2) || (end1 <= end2 && end1 >= start2) {
		overlaps = true
	}
	if (start2 >= start1 && start2 <= end1) || (end2 <= end1 && end2 >= start1) {
		overlaps = true
	}

	return contains, overlaps
}
