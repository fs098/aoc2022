package main

import "fmt"

type point struct {
	r int // row
	c int // col
}

func (p *point) add(p2 point) {
	p.r += p2.r
	p.c += p2.c
}

func day9(filename string) {
	data := getLines(filename)
	part1, part2 := ropeBridge(data)

	fmt.Println("Day 9")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func ropeBridge(data []string) (int, int) {
	var (
		visitedPart1 = make(map[point]bool)
		visitedPart2 = make(map[point]bool)
		head, tail   = &point{0, 0}, &point{0, 0}
		knots        = knotsInit()
	)

	for _, line := range data {
		words := getWords(line)
		count := readInt(words[1])

		for i := 0; i < count; i++ {
			moveHead(words[0], head, tail)
			moveKnots(words[0], knots)
			visitedPart1[*tail] = true
			visitedPart2[*knots[9]] = true
		}
	}

	return len(visitedPart1), len(visitedPart2)
}

func moveKnots(direction string, k []*point) {
	switch direction {
	case "U":
		k[0].add(point{-1, 0})
	case "D":
		k[0].add(point{1, 0})
	case "L":
		k[0].add(point{0, -1})
	case "R":
		k[0].add(point{0, 1})
	}

	for i := 1; i < len(k); i++ {
		moveTail(k[i-1], k[i])
	}
}

func moveHead(direction string, h, t *point) {
	switch direction {
	case "U":
		h.add(point{-1, 0})
	case "D":
		h.add(point{1, 0})
	case "L":
		h.add(point{0, -1})
	case "R":
		h.add(point{0, 1})
	}

	moveTail(h, t)
}

func moveTail(h, t *point) {
	if h.r-1 <= t.r && t.r <= h.r+1 && h.c-1 <= t.c && t.c <= h.c+1 {
		return
	}
	t.add(getMovement(h, t))
}

func getMovement(head, tail *point) point {
	// Head: {1, 2}
	// Tail {3, 1} -> {2, 2}
	// Head - Tail: {-2, 1} ==> moved {-1, 1}

	// Head: {0, 2}
	// Tail {1, 4} -> {0, 3}
	// Head - Tail: {-1, -2} ==> moved {-1, -1}

	m := pointSubtraction(head, tail)
	if m.r != 0 {
		if m.r > 0 {
			m.r = 1
		} else {
			m.r = -1
		}
	}
	if m.c != 0 {
		if m.c > 0 {
			m.c = 1
		} else {
			m.c = -1
		}
	}
	return m
}

func pointSubtraction(head, tail *point) point {
	return point{head.r - tail.r, head.c - tail.c}
}

func knotsInit() []*point {
	k := make([]*point, 10)
	for i := 0; i < 10; i++ {
		k[i] = &point{0, 0}
	}
	return k
}
