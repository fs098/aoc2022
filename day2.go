package main

import "fmt"

type shape int

const (
	rock shape = iota
	paper
	scissor
)

func day2(filename string) {
	fmt.Println("Day 2")

	part1 := totalScorePart1(filename)
	fmt.Println("Part 1:", part1)

	part2 := totalScorePart2(filename)
	fmt.Println("Part 2:", part2)
}

func totalScorePart1(filename string) int {
	result := 0
	data := getLines(filename)
	for _, v := range data {
		if v == "" {
			continue
		}
		words := getWords(v)
		opponent, me := strToShape(words[0]), strToShape(words[1])

		result += getScore(opponent, me)
	}
	return result
}

func totalScorePart2(filename string) int {
	result := 0
	data := getLines(filename)
	for _, v := range data {
		if v == "" {
			continue
		}

		var (
			words    = getWords(v)
			opponent = strToShape(words[0])
			me       = getWantedShape(opponent, words[1])
		)
		result += getScore(opponent, me)
	}
	return result
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
