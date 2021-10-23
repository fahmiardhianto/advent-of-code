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

	sum := 0
	groupAnswer := make(map[string]int)
	for scanner.Scan() {
		trimmed := strings.TrimSpace(scanner.Text())
		answers := strings.Split(trimmed, "\n")
		numOfPeople := len(answers)
		for _, a := range answers {
			questions := strings.Split(a, "")
			for _, q := range questions {
				_, ok := groupAnswer[q]
				if !ok {
					groupAnswer[q] = 1
				} else {
					groupAnswer[q] += 1
				}
				if groupAnswer[q] == numOfPeople {
					sum += 1
				}
			}
		}
		groupAnswer = make(map[string]int)
	}
	fmt.Println(sum)
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
