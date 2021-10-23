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
	groupAnswer := make(map[string]bool)
	for scanner.Scan() {
		replacedNewLine := strings.ReplaceAll(scanner.Text(), "\n", "")
		qs := strings.Split(replacedNewLine, "")
		for _, q := range qs {
			if _, ok := groupAnswer[q]; !ok {
				groupAnswer[q] = true
			}
		}
		fmt.Println(len(groupAnswer))
		sum += len(groupAnswer)
		groupAnswer = make(map[string]bool)
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
