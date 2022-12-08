package main

import "fmt"

type tree struct {
	r int // row
	c int // col
}

type Forest map[tree]int

func (t *tree) add(t2 tree) {
	t.r += t2.r
	t.c += t2.c
}

func day8(filename string) {
	data := getLines(filename)
	part1, part2 := treeHouse(data)

	fmt.Println("Day 8")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func treeHouse(data []string) (int, int) {
	var (
		forest = loadForest(data)
		part1  = 0
		part2  = 0
	)

	for tree, height := range forest {
		isVisible, scenicScore := visAndScore(tree, height, forest)

		if isVisible {
			part1++
		}

		if part2 < scenicScore {
			part2 = scenicScore
		}
	}

	return part1, part2
}

func visAndScore(t tree, height int, forest Forest) (bool, int) {
	var (
		directions  = []tree{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		isVisible   = false
		scenicScore = 1
	)

	for _, direction := range directions {
		dist := 1
		tree := t

		for {
			tree.add(direction)

			treeHeight, ok := forest[tree]
			if !ok {
				// reached edge
				isVisible = true
				scenicScore *= dist - 1
				break
			}

			if height <= treeHeight {
				scenicScore *= dist
				break
			}

			dist++
		}
	}

	return isVisible, scenicScore
}

func loadForest(data []string) Forest {
	f := make(map[tree]int)

	for r, line := range data {
		for c, char := range line {
			f[tree{r, c}] = readInt(string(char))
		}
	}
	return f
}
