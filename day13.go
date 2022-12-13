package main

import (
	"fmt"
	"sort"
)

type elem struct {
	elemType string
	number   int
	list     []elem
	div      string
}

type elemsPair struct {
	first  []elem
	second []elem
}

// compare returns 0 if elems are NOT in the right order, 1 if they are, 2 if pairs are equal
func (e *elemsPair) compare() int {
	for i := 0; i < min(len(e.first), len(e.second)); i++ {
		var (
			first  = e.first[i]
			second = e.second[i]
		)

		if first.elemType != second.elemType {
			if first.elemType == "number" {
				first = newElem("list", 0, first)
			}
			if second.elemType == "number" {
				second = newElem("list", 0, second)
			}
		}

		if first.elemType == "number" {
			if first.number < second.number {
				return 1
			}

			if first.number > second.number {
				return 0
			}
		}

		if first.elemType == "list" {
			pair := elemsPair{first.list, second.list}
			if pair.compare() == 1 {
				return 1
			}
			if pair.compare() == 0 {
				return 0
			}
		}
	}

	if len(e.first) == len(e.second) {
		return 2
	}

	if len(e.first) < len(e.second) {
		return 1
	}

	return 0
}

func newElem(etype string, val int, elems ...elem) elem {
	var result elem

	if etype == "number" {
		result = elem{
			elemType: etype,
			number:   val,
			list:     nil,
		}
	}

	if etype == "list" {
		result = elem{
			elemType: etype,
			number:   0,
			list:     elems,
		}
	}

	return result
}

func day13(filename string) {
	data := getLines(filename)
	part1, part2 := distressSignal(data)

	fmt.Println("Day 12")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func distressSignal(data []string) (int, int) {
	var (
		pairs    = parsePairs(data)
		elems    = pairsToElems(pairs)
		dividers = []string{"[[2]]", "[[6]]"}
		l1, l2   = getList(dividers[0], 0), getList(dividers[1], 0)
		div1     = newElem("list", 0, l1...)
		div2     = newElem("list", 0, l2...)

		part1 = 0
		part2 = 1
	)
	div1.div = "div1"
	div2.div = "div2"

	for i, p := range pairs {
		if p.compare() == 1 {
			part1 += i + 1
		}
	}

	elems = append(elems, div1, div2)

	sort.Slice(elems, func(i, j int) bool {
		pair := elemsPair{elems[i].list, elems[j].list}
		cmp := pair.compare()
		if cmp == 1 {
			return true
		}
		return false
	})

	for i, elem := range elems {
		if elem.div == "div1" {
			part2 *= (i + 1)
		}
		if elem.div == "div2" {
			part2 *= (i + 1)
		}
	}

	return part1, part2
}

func parsePairs(data []string) []elemsPair {
	var (
		result []elemsPair
		first  []elem
		second []elem
		count  = 1
	)

	data = removeEmpty(data)

	for _, line := range data {
		if count == 1 {
			first = getList(line, 0)
		}

		if count == 2 {
			second = getList(line, 0)
			result = append(result, elemsPair{first, second})
			count = 0
		}

		count++
	}

	return result
}

func getList(line string, index int) []elem {
	var (
		result []elem
		N      = getListN(line, index)
	)

	for i := index + 1; i < N; i++ {
		if line[i] == '[' {
			elem := newElem("list", 0, getList(line, i)...)
			result = append(result, elem)

			i = getListN(line, i)
			continue
		}

		if isInt(line[i]) {
			str := []byte{line[i]}
			for j := i + 1; j < len(line) && isInt(line[j]); j++ {
				str = append(str, line[j])
			}

			elem := newElem("number", readInt(string(str)))
			result = append(result, elem)
		}
	}
	return result
}

func getListN(line string, index int) int {
	parentheses := 0
	for i := index + 1; i < len(line); i++ {
		if line[i] == ']' && parentheses == 0 {
			return i
		}
		if line[i] == '[' {
			parentheses++
		}
		if line[i] == ']' {
			parentheses--
		}

	}
	return -1
}

func isInt(b byte) bool {
	return '0' <= b && b <= '9'
}

func pairsToElems(pairs []elemsPair) []elem {
	var result []elem
	for _, p := range pairs {
		result = append(result, newElem("list", 0, p.first...), newElem("list", 0, p.second...))
	}
	return result
}
