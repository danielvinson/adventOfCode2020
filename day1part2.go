package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func day1part2() {
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
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			for k := 0; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					fmt.Println(numbers[i], numbers[j], numbers[k], numbers[i]*numbers[j]*numbers[k])
				}
			}
		}
	}
}
