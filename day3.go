package main

import "fmt"

func day3(filename string) {
	var (
		data  = getLines(filename)
		part1 = sumOfPriorities1(data)
		part2 = sumOfPriorities2(data)
	)

	fmt.Println("Day 3")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func sumOfPriorities1(data []string) int {
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

func sumOfPriorities2(data []string) int {
	result := 0

	for i := 0; i < len(data); i += 3 {
		groupData := []string{data[i], data[i+1], data[i+2]}
		foundSet := make(map[byte]int) // foundSet[x] must == 3

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
