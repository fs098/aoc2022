package main

import (
	"fmt"
	"strings"
)

func day4(filename string) {
	fmt.Println("Day 4")

	part1, part2 := totalOverlap(filename)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func totalOverlap(filename string) (int, int) {
	var (
		data       = getLines(filename)
		contained  = 0
		overlapped = 0
	)

	for _, v := range data {
		contains, overlaps := existsOverlap(v)
		if contains {
			contained++
		}
		if overlaps {
			overlapped++
		}
	}
	return contained, overlapped
}

func existsOverlap(s string) (bool, bool) {
	var (
		values = strings.Split(s, ",")
		range1 = strings.Split(values[0], "-")
		range2 = strings.Split(values[1], "-")

		start1, end1 = readInt(range1[0]), readInt(range1[1])
		start2, end2 = readInt(range2[0]), readInt(range2[1])

		contains = false
		overlaps = false
	)

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
