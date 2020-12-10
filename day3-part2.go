package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func openFileString(fileName string) []string {
	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func checkSlope(input []string, down int, right int) int {
	treeCount := 0
	lineLength := len(input[0])
	for i := 0; i < len(input); i += down {
		line := input[i]
		horizontalLocation := (i * right) % lineLength
		if string(line[horizontalLocation]) == "#" {
			treeCount++
		}
	}
	return treeCount
}

func day3part2() {
	lines := openFileString("input-day3.txt")
	tests := [][]int{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}
	var results []int

	for _, test := range tests {
		fmt.Println(test[0], test[1])
		results = append(results, checkSlope(lines, test[0], test[1]))
	}

	fmt.Println(results)

	res := results[0] * results[1] * results[2] * results[3] * results[4]

	fmt.Println(res)
}
