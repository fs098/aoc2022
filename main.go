package main

import "fmt"

func main() {
	day := 12
	switch day {
	case 0:
		day12("./input/day12_test.txt")
	case 1:
		day1("./input/day1.txt")
	case 2:
		day2("./input/day2.txt")
	case 3:
		day3("./input/day3.txt")
	case 4:
		day4("./input/day4.txt")
	case 5:
		day5("./input/day5.txt")
	case 6:
		day6("./input/day6.txt")
	case 7:
		day7("./input/day7.txt")
	case 8:
		day8("./input/day8.txt")
	case 9:
		day9("./input/day9.txt")
	case 10:
		day10("./input/day10.txt")
	case 11:
		day11("./input/day11.txt")
	case 12:
		day12("./input/day12.txt")

	default:
		fmt.Println("Invalid input:", day)
	}
}
