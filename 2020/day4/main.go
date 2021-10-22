package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
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

	validCount := 0
	for scanner.Scan() {
		replacedNewLine := strings.ReplaceAll(scanner.Text(), "\n", " ")
		passportData := strings.TrimRight(replacedNewLine, " ")
		valid := isValidPassport(strings.Split(passportData, " "))
		if valid {
			validCount++
		}
	}

	fmt.Println(validCount)
}

func isValidPassport(data []string) bool {
	checkBit := 0
	for _, pair := range data {
		kv := strings.Split(pair, ":")
		key := kv[0]
		val := kv[1]

		switch key {
		case "byr":
			year, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			if year >= 1920 && year <= 2002 {
				checkBit = checkBit | 1
			}
		case "iyr":
			year, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			if year >= 2010 && year <= 2020 {
				checkBit = checkBit | 1<<1
			}
		case "eyr":
			year, err := strconv.Atoi(val)
			if err != nil {
				continue
			}
			if year >= 2020 && year <= 2030 {
				checkBit = checkBit | 1<<2
			}
		case "hgt":
			if strings.HasSuffix(val, "cm") {
				h, err := strconv.Atoi(strings.TrimSuffix(val, "cm"))
				if err != nil {
					continue
				}
				if h >= 150 && h <= 193 {
					checkBit = checkBit | 1<<3
				}
			}
			if strings.HasSuffix(val, "in") {
				h, err := strconv.Atoi(strings.TrimSuffix(val, "in"))
				if err != nil {
					continue
				}
				if h >= 59 && h <= 76 {
					checkBit = checkBit | 1<<3
				}
			}
		case "hcl":
			_, err := hex.DecodeString(strings.TrimPrefix(val, "#"))
			if err != nil || !strings.HasPrefix(val, "#") {
				continue
			}
			checkBit = checkBit | 1<<4
		case "ecl":
			for _, v := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
				if val == v {
					checkBit = checkBit | 1<<5
					break
				}
			}
		case "pid":
			_, err := strconv.Atoi(val)
			if err == nil && len(val) == 9 {
				checkBit = checkBit | 1<<6
			}
		}
	}

	if checkBit == 1<<7-1 {
		return true
	}

	return false
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
