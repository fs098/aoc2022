package main

import (
	"fmt"
	"sort"
	"strings"
)

type monkeyList []*monkey

type monkey struct {
	items     []int
	inspected int

	opMult bool
	multBy int

	opAdd bool
	addBy int

	divBy   int
	ifTrue  int
	ifFalse int
}

func day11(filename string) {
	data := getLines(filename)
	part1, part2 := monkeyBusiness(data)

	fmt.Println("Day 11")
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}

func monkeyBusiness(data []string) (int, int) {
	var (
		monkeysPart1 = getMonkeyList(data)
		monkeysPart2 = getMonkeyList(data)
		modInt       = 1
	)

	for _, m := range *monkeysPart2 {
		modInt *= m.divBy
	}

	// part 1:
	for round := 0; round < 20; round++ {
		for i := 0; i < len(*monkeysPart1); i++ {
			inspectItems(monkeysPart1, i, modInt, "part1")
		}
	}
	// part 2:
	for round := 0; round < 10000; round++ {
		for i := 0; i < len(*monkeysPart2); i++ {
			inspectItems(monkeysPart2, i, modInt, "part2")
		}
	}

	inspectedList1 := []int{}
	inspectedList2 := []int{}
	for i := 0; i < len(*monkeysPart1); i++ {
		inspectedList1 = append(inspectedList1, (*monkeysPart1)[i].inspected)
		inspectedList2 = append(inspectedList2, (*monkeysPart2)[i].inspected)
	}

	sort.Slice(inspectedList1, func(i, j int) bool {
		return inspectedList1[i] > inspectedList1[j]
	})
	sort.Slice(inspectedList2, func(i, j int) bool {
		return inspectedList2[i] > inspectedList2[j]
	})

	part1 := inspectedList1[0] * inspectedList1[1]
	part2 := inspectedList2[0] * inspectedList2[1]
	return part1, part2
}

func inspectItems(list *monkeyList, index int, modInt int, part string) {
	m := (*list)[index]
	items := m.items
	m.inspected += len(items)

	if m.opAdd {
		for i := 0; i < len(items); i++ {
			items[i] += m.addBy
		}
	}

	if m.opMult {
		for i := 0; i < len(items); i++ {
			if m.multBy == -1 {
				items[i] *= items[i]
			} else {
				items[i] *= m.multBy
			}
		}
	}

	for i := 0; i < len(items); i++ {
		if part == "part1" {
			items[i] /= 3
		}
		if part == "part2" {
			items[i] %= modInt
		}
	}

	for i := 0; i < len(items); i++ {
		if items[i]%m.divBy == 0 {
			(*list)[m.ifTrue].items = append((*list)[m.ifTrue].items, items[i])
			continue
		}
		(*list)[m.ifFalse].items = append((*list)[m.ifFalse].items, items[i])
	}

	m.items = []int{}
}

func parseMonkey(index int, data []string) *monkey {
	var (
		m             monkey
		startingItems = strings.TrimSpace(data[index])
		operation     = strings.TrimSpace(data[index+1])
		testDivisible = strings.TrimSpace(data[index+2])
		ifTrue        = strings.TrimSpace(data[index+3])
		ifFalse       = strings.TrimSpace(data[index+4])
	)

	startingItems = strings.ReplaceAll(startingItems, "Starting items: ", "")
	m.items = mapReadInts(strings.Split(startingItems, ", "))

	m.inspected = 0

	operation = strings.ReplaceAll(operation, "Operation: new = old ", "")
	if operation[0] == '*' {
		m.opMult = true
		if strings.Contains(operation, "old") {
			m.multBy = -1
		} else {
			m.multBy = readInt(operation[2:])
		}
	} else {
		m.opAdd = true
		m.addBy = readInt(operation[2:])
	}

	testDivisible = strings.ReplaceAll(testDivisible, "Test: divisible by ", "")
	m.divBy = readInt(testDivisible)

	ifTrue = strings.ReplaceAll(ifTrue, "If true: throw to monkey ", "")
	m.ifTrue = readInt(ifTrue)

	ifFalse = strings.ReplaceAll(ifFalse, "If false: throw to monkey ", "")
	m.ifFalse = readInt(ifFalse)

	return &m
}

func getMonkeyList(data []string) *monkeyList {
	result := monkeyList{}
	for i := 0; i < len(data); i++ {
		if strings.HasPrefix(data[i], "Monkey") {
			m := parseMonkey(i+1, data)
			result = append(result, m)
		}
	}
	return &result
}
