package main

import "fmt"

func main() {
	day := 1

	switch day {
	case 0:
		day1("./input/day1_test.txt")

	case 1:
		day1("./input/day1.txt")

	default:
		fmt.Println("Invalid input:", day)
	}
}
