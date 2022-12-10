package main

import "fmt"

func day10(filename string) {
	data := getLines(filename)
	part1, part2 := cathodeRayTube(data)

	fmt.Println("Day 10")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:")
	drawScreen(part2)
}

func cathodeRayTube(data []string) (int, [][]bool) {
	var (
		signalStr   = 0
		register    = 1
		cycle       = 0
		targetCycle = 20

		sprite []bool
		screen = initScreen()
	)

	for _, line := range data {
		words := getWords(line)
		cyclesToComplete := 1

		if words[0] == "addx" {
			cyclesToComplete = 2
		}

		for i := 0; i < cyclesToComplete; i++ {
			cycle++
			updateSignalStr(cycle, register, &signalStr, &targetCycle)
			sprite = updateSprite(register)
			screen = updateScreen(screen, sprite, cycle)
		}

		if cyclesToComplete == 2 {
			register += readInt(words[1])
		}
	}

	return signalStr, screen
}

func updateSignalStr(cycle, register int, result, targetCycle *int) {
	if cycle > 220 {
		return
	}

	if cycle == *targetCycle {
		*targetCycle += 40

		*result += cycle * register
	}
}

func updateSprite(register int) []bool {
	sprite := make([]bool, 40)
	start := register - 1

	for i := start; i < start+3 && i < 40; i++ {
		if i < 0 {
			continue
		}
		sprite[i] = true
	}

	return sprite
}

func initScreen() [][]bool {
	screen := make([][]bool, 6)
	for i := 0; i < 6; i++ {
		row := make([]bool, 40)
		screen[i] = row
	}
	return screen
}

func updateScreen(screen [][]bool, sprite []bool, cycle int) [][]bool {
	row, col := screenRow(cycle), screenCol(cycle)
	if sprite[col] {
		screen[row][col] = true
	}
	return screen
}

func drawScreen(screen [][]bool) {
	for _, line := range screen {
		for _, b := range line {
			if b {
				fmt.Printf("##")
				continue
			}
			fmt.Printf("  ")
		}
		fmt.Printf("\n")
	}
}

func screenRow(cycle int) int {
	index := 0
	targetCycle := 41

	for cycle >= targetCycle {
		targetCycle += 40
		index++
	}

	return index
}

func screenCol(cycle int) int {
	index := cycle - (screenRow(cycle) * 40) - 1
	return index
}
