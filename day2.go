package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day2() {
	// Read file
	file, err := os.Open("input-day2.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Parse file lines to string

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Solve
	var validPasswordsCount int = 0
	var sbuf []string
	var password string
	var passwordRule string
	var passwordRuleChar string
	var passwordRuleMin int64
	var passwordRuleMax int64
	var charCount int64

	for i := 0; i < len(lines); i++ {
		sbuf = strings.Split(lines[i], ":")
		passwordRule, password = sbuf[0], sbuf[1]
		sbuf = strings.Split(passwordRule, " ")
		passwordRuleChar = sbuf[1]
		sbuf = strings.Split(sbuf[0], "-")
		passwordRuleMin, err = strconv.ParseInt(sbuf[0], 10, 64)
		passwordRuleMax, err = strconv.ParseInt(sbuf[1], 10, 64)

		charCount = 0
		// Count instances of the character which is required
		for _, c := range password {
			if string(c) == passwordRuleChar {
				charCount++
			}
		}

		// Check if that is within expected range
		if charCount >= passwordRuleMin && charCount <= passwordRuleMax {
			validPasswordsCount++
		}
	}

	fmt.Println(validPasswordsCount)
}
