package main

import (
	"bufio"
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

	valid1 := 0
	valid2 := 0
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		low, high := parseRange(s[0])
		char := strings.TrimSuffix(s[1], ":")
		password := s[2]

		if validPolicy1(password, char, low, high) {
			valid1++
		}
		if validPolicy2(password, char, low, high) {
			valid2++
		}
	}

	fmt.Println(valid1)
	fmt.Println(valid2)
}

func parseRange(s string) (int, int) {
	d := strings.Split(s, "-")
	var minMax []int
	for _, v := range d {
		x, _ := strconv.Atoi(v)
		minMax = append(minMax, x)
	}

	return minMax[0], minMax[1]
}

func validPolicy1(password, char string, min, max int) bool {
	contains := strings.Count(password, char)
	if contains >= min && contains <= max {
		return true
	}
	return false
}

func validPolicy2(password, char string, pos1, pos2 int) bool {
	pos1 -= 1
	pos2 -= 1
	splitPass := strings.Split(password, "")
	a := splitPass[pos1] == char
	b := splitPass[pos2] == char

	if (a || b) && !(a && b) {
		return true
	}
	return false
}
