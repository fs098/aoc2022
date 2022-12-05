package main

import (
	"fmt"
	"log"
	"strings"
)

type stack struct {
	crates []string
}

func newStack() *stack {
	return &stack{
		crates: []string{},
	}
}

func (s *stack) push(c string) {
	(*s).crates = append((*s).crates, c)
}

func (s *stack) pop() string {
	if len((*s).crates) == 0 {
		log.Fatal("[Error] invalid pop operation")
	}

	crate := (*s).crates[len((*s).crates)-1]
	(*s).crates = (*s).crates[:len((*s).crates)-1]

	return crate
}

func (s *stack) popMultiple(n int) []string {
	if len((*s).crates) < n {
		log.Fatal("[Error] invalid popMultiple operation")
	}

	crates := []string{}
	for i := 0; i < n; i++ {
		crate := (*s).crates[len((*s).crates)-1]
		(*s).crates = (*s).crates[:len((*s).crates)-1]
		crates = append(crates, crate)
	}

	return crates
}

// day5 main function
func day5(filename string) {
	data := getLines(filename)
	part1, part2 := rearrangeCrates(data)

	fmt.Println("Day 5")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func rearrangeCrates(data []string) (string, string) {
	var (
		cargoPart1 = getCargo(data)
		cargoPart2 = getCargo(data)
		start      = getStartLine(data)

		part1 = ""
		part2 = ""
	)

	for i := start; i < len(data); i++ {
		applyMovement(data[i], cargoPart1, "9000")
		applyMovement(data[i], cargoPart2, "9001")
	}

	for i := 0; i < len(cargoPart1); i++ {
		topP1 := cargoPart1[i].pop()
		part1 += string(topP1[1])

		topP2 := cargoPart2[i].pop()
		part2 += string(topP2[1])
	}

	return part1, part2
}

func applyMovement(line string, cargo []*stack, craneType string) {
	var quantity, src, dst int

	_, err := fmt.Sscanf(line, "move %d from %d to %d", &quantity, &src, &dst)
	if err != nil {
		log.Fatal(err)
	}

	// correct index
	src--
	dst--

	// part 1
	if craneType == "9000" {
		for i := 0; i < quantity; i++ {
			crate := cargo[src].pop()
			cargo[dst].push(crate)
		}
		return
	}

	// part 2
	if craneType == "9001" {
		if quantity == 1 {
			crate := cargo[src].pop()
			cargo[dst].push(crate)
			return
		} else {
			crates := cargo[src].popMultiple(quantity)
			for i := len(crates) - 1; i >= 0; i-- {
				cargo[dst].push(crates[i])
			}
		}
		return
	}
	log.Fatal("Invalid Crane Type")
}

func getCargo(data []string) []*stack {
	var (
		size  = getCargoSize(data)
		cargo = make([]*stack, size)
		start = getStartLine(data) - 3
	)

	for i := 0; i < size; i++ {
		stack := newStack()
		cargo[i] = stack
	}

	for i := start; i >= 0; i-- {
		crates := stringToCrates(data[i])
		for j, crate := range crates {
			if crate == "   " {
				continue
			}
			cargo[j].push(crate)
		}
	}
	return cargo
}

func stringToCrates(s string) []string {
	result := []string{}
	for i := 0; i < len(s); i += 4 {
		crate := []byte{s[i], s[i+1], s[i+2]}
		result = append(result, string(crate))
	}
	return result
}

func getStartLine(data []string) int {
	for i := 0; i < len(data); i++ {
		if strings.HasPrefix(data[i], "move") {
			return i
		}
	}
	return 0
}

func getCargoSize(data []string) int {
	result := 0
	for i := 0; i < len(data[0]); i += 4 {
		result++
	}
	return result
}
