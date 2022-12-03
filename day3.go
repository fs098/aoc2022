package main

import "fmt"

func day3(filename string) {
	fmt.Println("Day 3")

	part1 := sumOfPriorities(filename)
	fmt.Println("Part 1:", part1)

	part2 := sumOfPrioritiesPart2(filename)
	fmt.Println("Part 2:", part2)
}

func sumOfPriorities(filename string) int {
	data := getLines(filename)
	result := 0

	for _, val := range data {
		var (
			middle     = len(val) / 2
			partition1 = val[:middle]
			partition2 = val[middle:]

			set = make(map[byte]bool)
		)

		for i := 0; i < len(partition1); i++ {
			set[partition1[i]] = true
		}

		for i := 0; i < len(partition2); i++ {
			if set[partition2[i]] {
				result += priority(partition2[i])
				break
			}
		}
	}
	return result
}

func sumOfPrioritiesPart2(filename string) int {
	data := getLines(filename)
	result := 0

	for i := 0; i < len(data); i += 3 {
		groupData := []string{data[i], data[i+1], data[i+2]}
		foundSet := make(map[byte]int) // must == 3

		for _, val := range groupData {
			alreadyFound := make(map[byte]bool)

			for j := 0; j < len(val); j++ {
				item := val[j]

				if alreadyFound[item] {
					continue
				}
				alreadyFound[item] = true
				foundSet[item]++
			}
		}

		for key, val := range foundSet {
			if val == 3 {
				result += priority(key)
			}
		}
	}
	return result
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
