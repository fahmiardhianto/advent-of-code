package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(SplitBlankLine)

	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	validCount := 0
	for scanner.Scan() {
		passport := scanner.Text()
		valid := true
		for _, r := range required {
			containsRequired := strings.Contains(passport, r)
			if !containsRequired {
				valid = false
				break
			}

		}
		if valid {
			validCount++
		}
	}

	fmt.Println(validCount)
}

func SplitBlankLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	blankLineBytes := []byte("\n\n")
	blankLineLen := len(blankLineBytes)
	dataLen := len(data)

	if atEOF && dataLen == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, blankLineBytes); i >= 0 {
		return i + blankLineLen, data[0:i], nil
	}

	if atEOF {
		return dataLen, data, nil
	}

	return 0, nil, nil
}
