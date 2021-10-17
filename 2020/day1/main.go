package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	entries := make([]int, 0)
	for scanner.Scan() {
		e, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		entries = append(entries, e)
	}

	sort.Ints(entries)

	twoEntries := findTwoEntries(entries, 2020)
	fmt.Println(reduce(twoEntries))

	threeEntries := findThreeEntries(entries, 2020)
	fmt.Println(reduce(threeEntries))
}

func reduce(numbers []int) int {
	result := 1
	for _, v := range numbers {
		result *= v
	}
	return result
}

func findTwoEntries(entries []int, target int) []int {
	left := 0
	right := len(entries) - 1

	for left < right {
		sum := entries[left] + entries[right]
		if sum == target {
			return []int{entries[left], entries[right]}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}

	return []int{}
}

func findThreeEntries(entries []int, target int) []int {
	threeEntries := make([]int, 3)
	for i, e := range entries {
		threeEntries = []int{e}
		twoEntries := findTwoEntries(entries[i+1:], target-e)

		if len(twoEntries) == 2 {
			threeEntries = append(threeEntries, twoEntries...)
			return threeEntries
		}
	}

	return threeEntries
}
