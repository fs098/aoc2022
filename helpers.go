package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// getLines returns a list of strings, each one representing a line of a file
func getLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open %s\n", filename)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

// getWords splits a string into words, separated by spaces
func getWords(line string) []string {
	return strings.Split(line, " ")
}

// readInt converts a string to a int
func readInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

// mapToInts maps Atoi to a list of strings
func mapReadInts(list []string) []int {
	ints := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		ints[i] = readInt(list[i])
	}
	return ints
}

// removeEmpty removes "" from a list of words
func removeEmpty(words []string) []string {
	result := []string{}
	for i := 0; i < len(words); i++ {
		if words[i] != "" {
			result = append(result, words[i])
		}
	}
	return result
}

// sum returns the sum of a slice of ints
func sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// min returns the min of given values
func min(nums ...int) int {
	min := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}

// max returns the max of given values
func max(nums ...int) int {
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
