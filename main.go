package main

import "fmt"

func main() {
	day := 2
	switch day {
	case 0:
		day2("./input/day2_test.txt")

	case 1:
		day1("./input/day1.txt")

	case 2:
		day2("./input/day2.txt")

	default:
		fmt.Println("Invalid input:", day)
	}
}
