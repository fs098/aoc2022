package main

import (
	"fmt"
	"log"
)

func day6(filename string) {
	data := getLines(filename)
	part1, part2 := findMarkerStart(data[0])

	fmt.Println("Day 6")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func findMarkerStart(data string) (int, int) {
	var (
		packetFound  = false
		messageFound = false
		packetStart  = 0
		messageStart = 0
	)

	for i := 0; i < len(data); i++ {
		if packetFound && messageFound {
			break
		}
		if !packetFound {
			if isMarker(i, 4, data) {
				packetStart = i + 4
				packetFound = true
			}
		}
		if !messageFound {
			if isMarker(i, 14, data) {
				messageStart = i + 14
				messageFound = true
			}
		}
	}

	return packetStart, messageStart
}

func isMarker(index, end int, data string) bool {
	set := make(map[byte]bool)

	if len(data) < index+end {
		log.Fatalf("invalid index at %d + %d", index, end)
	}

	for i := index; i < index+end; i++ {
		if set[data[i]] {
			return false
		}
		set[data[i]] = true
	}
	return true
}
