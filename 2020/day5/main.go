package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

	max := -math.MaxUint32
	seats := make([]int, 0)
	for scanner.Scan() {
		boardingPass := scanner.Text()

		row := boardingPass[:7]
		row = strings.ReplaceAll(row, "F", "0")
		row = strings.ReplaceAll(row, "B", "1")

		column := boardingPass[7:]
		column = strings.ReplaceAll(column, "R", "1")
		column = strings.ReplaceAll(column, "L", "0")

		r, _ := strconv.ParseUint(row, 2, 32)
		c, _ := strconv.ParseUint(column, 2, 32)
		v := int(r*8 + c)
		seats = append(seats, v)

		if v > max {
			max = v
		}
	}
	fmt.Println(max)

	sort.Ints(seats)
	for i, v := range seats {
		if i > 0 && i < len(seats)-1 {
			if seats[i+1]-v > 1 {
				fmt.Println("my seat: ", v+1)
			}
		}
	}
}
