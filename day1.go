package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func testNum(num int64, rest []int64) int64 {
	for i := 0; i < len(rest); i++ {
		if num+rest[i] == 2020 {
			return rest[i]
		}
	}

	return 0
}

func day1() {
	// Read file
	file, err := os.Open("input-day1.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse file lines to int[]

	var numbers []int64
	var line int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// In other languages, I'd just convert to a Set data structure here
	// but since that seems to be pretty annoying in go, I'm going to use
	// a int[] that is a bit slower.

	// Solve
	var buf int64

	for i := 0; i < len(numbers); i++ {
		// This is just a bad n^2 algorithm
		buf = testNum(numbers[i], numbers[i + 1:])
		if buf != 0 {
			fmt.Println(buf, numbers[i], buf * numbers[i])
			break
		}
	}
}
