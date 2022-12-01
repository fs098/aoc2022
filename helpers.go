package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// getLines returns a list of strings, each one representing a line of a file
func getLines(fileName string) []string {

	// os.Open() opens specific file in
	// read-only mode and this return
	// a pointer of type os.
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalf("failed to open")

	}

	// The bufio.NewScanner() function is called in which the
	// object os.File passed as its parameter and this returns a
	// object bufio.Scanner which is further used on the
	// bufio.Scanner.Split() method.
	scanner := bufio.NewScanner(file)

	// The bufio.ScanLines is used as an
	// input to the method bufio.Scanner.Split()
	// and then the scanning forwards to each
	// new line using the bufio.Scanner.Scan()
	// method.
	scanner.Split(bufio.ScanLines)

	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	// The method os.File.Close() is called
	// on the os.File object to close the file
	file.Close()

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
		fmt.Println(err)
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
