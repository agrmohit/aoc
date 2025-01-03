package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/agrmohit/aoc/internal/inputs"
)

//go:embed input.txt
var input string

func concat(a int, b int) int {
	result, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))

	return result
}

// doesSolutionExist checks whether a solution exists recursively
func doesSolutionExist(result int, current int, rest []int) bool {
	// Exit condition
	if len(rest) == 0 {
		return result == current
	}

	current1 := rest[0]
	rest1 := rest[1:]

	return doesSolutionExist(result, current+current1, rest1) || doesSolutionExist(result, current*current1, rest1)
}

// doesSolutionExist checks whether a solution exists recursively, including concat operator
func doesSolutionExist2(result int, current int, rest []int) bool {
	// Exit condition
	if len(rest) == 0 {
		return result == current
	}

	current1 := rest[0]
	rest1 := rest[1:]

	return doesSolutionExist2(result, current+current1, rest1) || doesSolutionExist2(result, current*current1, rest1) || doesSolutionExist2(result, concat(current, current1), rest1)
}

func solvePart1(input string) int {
	input = strings.ReplaceAll(input, ":", "")
	rows, err := inputs.ExtractIntRows(input)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	solutionExistsSum := 0

	for _, line := range rows {
		if doesSolutionExist(line[0], line[1], line[2:]) {
			solutionExistsSum += line[0]
		}
	}

	return solutionExistsSum
}

func solvePart2(input string) int {
	input = strings.ReplaceAll(input, ":", "")
	rows, err := inputs.ExtractIntRows(input)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	solutionExistsSum := 0

	for _, line := range rows {
		if doesSolutionExist2(line[0], line[1], line[2:]) {
			solutionExistsSum += line[0]
		}
	}

	return solutionExistsSum
}

func main() {
	part1Solution := solvePart1(input)
	part2Solution := solvePart2(input)

	fmt.Println("Day 07 Part 1 solution:", part1Solution)
	fmt.Println("Day 07 Part 2 solution:", part2Solution)
}
