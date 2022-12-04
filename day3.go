package main

import "fmt"

func day3(filename string) {
	data := getLines(filename)
	part1, part2 := sumOfPriorities(data)

	fmt.Println("Day 3")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func sumOfPriorities(data []string) (int, int) {
	part1 := 0
	part2 := 0

	for i := 0; i < len(data); i += 3 {
		groupData := []string{data[i], data[i+1], data[i+2]}
		foundSet := make(map[byte]int) // foundSet[x] must == 3

		for _, line := range groupData {
			part1 += priorityFromLine(line)
			foundSetFromLine(line, &foundSet)
		}

		for key, val := range foundSet {
			if val == 3 {
				part2 += priority(key)
			}
		}
	}
	return part1, part2
}

// part 1
func priorityFromLine(line string) int {
	var (
		middle     = len(line) / 2
		partition1 = line[:middle]
		partition2 = line[middle:]

		set = make(map[byte]bool)
	)

	for i := 0; i < len(partition1); i++ {
		set[partition1[i]] = true
	}

	for i := 0; i < len(partition2); i++ {
		if set[partition2[i]] {
			return priority(partition2[i])
		}
	}
	return 0
}

// part 2
func foundSetFromLine(line string, set *map[byte]int) {
	foundSet := *set
	alreadyFound := make(map[byte]bool)

	for j := 0; j < len(line); j++ {
		item := line[j]

		if alreadyFound[item] {
			continue
		}
		alreadyFound[item] = true
		foundSet[item]++
	}
}

func priority(b byte) int {
	if 'a' <= b && b <= 'z' {
		return int(b) - 'a' + 1
	}
	if 'A' <= b && b <= 'Z' {
		return int(b) - 'A' + 27
	}
	return 0
}
