package main

import (
	"bufio"
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

	slopes := []*point{
		{slopeRight: 1, slopeDown: 1},
		{slopeRight: 3, slopeDown: 1},
		{slopeRight: 5, slopeDown: 1},
		{slopeRight: 7, slopeDown: 1},
		{slopeRight: 1, slopeDown: 2},
	}
	lineNum := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		for _, p := range slopes {
			if lineNum == p.y {
				if line[p.x] == "#" {
					p.treeCount++
				}

				p.nextPoint(len(line))
			}
		}
		lineNum++
	}

	multiplication := 1
	for _, p := range slopes {
		fmt.Println(p.treeCount)
		multiplication *= p.treeCount
	}

	fmt.Println(multiplication)
}

type point struct {
	x          int
	y          int
	slopeRight int
	slopeDown  int
	treeCount  int
}

func (p *point) nextPoint(width int) *point {
	p.x += p.slopeRight
	if p.x > width-1 {
		p.x %= width
	}
	p.y += p.slopeDown
	return p
}
