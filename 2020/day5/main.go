package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	max := -math.MaxUint32
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

		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}
