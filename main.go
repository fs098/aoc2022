package main

import "fmt"

func main() {
	day := 3
	switch day {
	case 0:
		day3("./input/day3_test.txt")

	case 1:
		day1("./input/day1.txt")

	case 2:
		day2("./input/day2.txt")

	case 3:
		day3("./input/day3.txt")

	default:
		fmt.Println("Invalid input:", day)
	}
}
