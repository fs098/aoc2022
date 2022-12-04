package main

import "fmt"

type shape int

const (
	rock shape = iota
	paper
	scissor
)

func day2(filename string) {
	data := getLines(filename)
	part1, part2 := totalScores(data)

	fmt.Println("Day 2")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func totalScores(data []string) (int, int) {
	part1 := 0
	part2 := 0

	for _, v := range data {
		var (
			words    = getWords(v)
			opponent = strToShape(words[0])
			mePart1  = strToShape(words[1])
			mePart2  = getWantedShape(opponent, words[1])
		)

		part1 += getScore(opponent, mePart1)
		part2 += getScore(opponent, mePart2)
	}
	return part1, part2
}

func getScore(opponent shape, me shape) int {
	result := 0
	if me == opponent {
		result += 3
	}
	if me == getWinningShape(opponent) {
		result += 6
	}

	result += getShapeScore(me)

	return result
}

func getShapeScore(s shape) int {
	if s == rock {
		return 1
	}
	if s == paper {
		return 2
	}
	if s == scissor {
		return 3
	}
	return -1
}

func getWantedShape(opponent shape, s string) shape {
	if s == "X" {
		return getLoosingShape(opponent)
	}
	if s == "Y" {
		return opponent
	}
	if s == "Z" {
		return getWinningShape(opponent)
	}
	return -1
}

func getWinningShape(s shape) shape {
	if s == rock {
		return paper
	}
	if s == paper {
		return scissor
	}
	if s == scissor {
		return rock
	}
	return -1
}

func getLoosingShape(s shape) shape {
	if s == rock {
		return scissor
	}
	if s == paper {
		return rock
	}
	if s == scissor {
		return paper
	}
	return -1
}

func strToShape(s string) shape {
	if s == "A" || s == "X" {
		return rock
	}
	if s == "B" || s == "Y" {
		return paper
	}
	if s == "C" || s == "Z" {
		return scissor
	}
	return -1
}
