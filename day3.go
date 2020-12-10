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

func day3() {
	lines := openFileString("input-day3.txt")
	lineLength := len(lines[0])
	treeCount := 0

	for i, line := range lines {
		horizontalLocation := (i * 3) % lineLength
		if string(line[horizontalLocation]) == "#" {
			treeCount++
		}
	}

	fmt.Println(treeCount)
}
