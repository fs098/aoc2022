package main

import "fmt"

type grid [][]int

func day8(filename string) {
	data := getLines(filename)
	part1, part2 := treeHouse(data)

	fmt.Println("Day 8")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func treeHouse(data []string) (int, int) {
	var (
		grid = loadGrid(data)

		part1 = 0
		part2 = 0
	)

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			isVisible, scenicScore := visibilityAndDistance(r, c, grid)

			if isVisible {
				part1++
			}

			if part2 < scenicScore {
				part2 = scenicScore
			}
		}
	}
	return part1, part2
}

func visibilityAndDistance(r int, c int, g grid) (bool, int) {
	var (
		rows, cols = len(g), len(g[0])
		val        = g[r][c]

		// part 1
		up    = true
		down  = true
		left  = true
		right = true
		// part 2
		distanceUp    = 0
		distanceDown  = 0
		distanceLeft  = 0
		distanceRight = 0
	)

	// up
	for i := r - 1; i >= 0; i-- {
		if val <= g[i][c] {
			up = false
			distanceUp++
			break
		}
		distanceUp++
	}
	// down
	for i := r + 1; i < rows; i++ {
		if val <= g[i][c] {
			down = false
			distanceDown++
			break
		}
		distanceDown++
	}
	// left
	for i := c - 1; i >= 0; i-- {
		if val <= g[r][i] {
			left = false
			distanceLeft++
			break
		}
		distanceLeft++
	}
	// right
	for i := c + 1; i < cols; i++ {
		if val <= g[r][i] {
			right = false
			distanceRight++
			break
		}
		distanceRight++
	}

	isVisible := up || down || left || right
	scenicScore := distanceUp * distanceDown * distanceLeft * distanceRight

	return isVisible, scenicScore
}

func loadGrid(data []string) grid {
	rows := len(data)
	cols := len(data[0])
	g := make([][]int, rows)
	for i, val := range data {
		r := make([]int, cols)
		for j, char := range val {
			r[j] = readInt(string(char))
		}
		g[i] = r
	}
	return g
}
