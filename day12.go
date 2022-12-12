package main

import "fmt"

type position struct {
	r int
	c int
}

func day12(filename string) {
	data := getLines(filename)
	part1, part2 := hillClimbing(data)

	fmt.Println("Day 12")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func hillClimbing(data []string) (int, int) {
	var (
		start, end, grid = loadGrid(data)
		part1            = shortestPath(start, end, grid)
		paths            = []int{}
	)

	for r, line := range grid {
		for c, n := range line {
			if n == 1 {
				newStart := position{r, c}
				path := shortestPath(newStart, end, grid)
				if path > 0 {
					paths = append(paths, path)
				}
			}
		}
	}
	part2 := min(paths...)

	return part1, part2
}

func shortestPath(start, end position, grid [][]int) int {
	var (
		rows, cols = len(grid), len(grid[0])

		visit  = make(map[position]bool)
		queue  = []position{}
		result = 0
	)
	queue = append(queue, start)
	visit[start] = true

	for len(queue) > 0 {
		qLen := len(queue)

		for i := 0; i < qLen; i++ {
			pos := queue[0]
			queue = queue[1:]

			if pos == end {
				return result
			}

			neighbors := getNeighbors(pos)
			for _, n := range neighbors {
				var (
					r, c        = n.r, n.c
					outOfBounds = min(r, c) < 0 || r >= rows || c >= cols
				)

				if outOfBounds || visit[n] {
					continue
				}

				var (
					curHeight  = grid[pos.r][pos.c]
					destHeight = grid[r][c]
					invalid    = curHeight+1 < destHeight
				)

				if invalid {
					continue
				}

				queue = append(queue, n)
				visit[n] = true
			}
		}
		result++
	}
	return -1
}

func getNeighbors(p position) []position {
	var (
		r, c  = p.r, p.c
		up    = position{r - 1, c}
		down  = position{r + 1, c}
		left  = position{r, c - 1}
		right = position{r, c + 1}
	)

	return []position{down, up, right, left}
}

func loadGrid(data []string) (position, position, [][]int) {
	var (
		start position
		end   position
		grid  [][]int
	)

	for r, line := range data {
		var row []int
		for c, char := range line {
			if char == 'S' {
				start = position{r, c}
				row = append(row, positionHeight('a'))
				continue
			}
			if char == 'E' {
				end = position{r, c}
				row = append(row, positionHeight('z'))
				continue
			}
			row = append(row, positionHeight(char))
		}
		grid = append(grid, row)
	}

	return start, end, grid
}

func positionHeight(b rune) int {
	if 'a' <= b && b <= 'z' {
		return int(b) - 'a' + 1
	}
	return 0
}
