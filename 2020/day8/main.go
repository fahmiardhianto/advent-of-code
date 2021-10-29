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
	bc := bootCode{}
	for scanner.Scan() {
		bc.parseInstruction(scanner.Text())
	}
	bc.trace()
}

type bootCode struct {
	codes       []instruction
	visited     []bool
	accumulator int
}

type instruction struct {
	operation string
	argument  int
}

func (bc *bootCode) parseInstruction(i string) {
	s := strings.Split(i, " ")
	arg, _ := strconv.Atoi(s[1])
	inst := instruction{operation: s[0], argument: arg}
	bc.codes = append(bc.codes, inst)
}

func (bc *bootCode) trace() {
	bc.visited = make([]bool, len(bc.codes))
	line := 0

	for line < len(bc.codes) && !bc.visited[line] {
		code := bc.codes[line]
		bc.visited[line] = true
		switch code.operation {
		case "nop":
			line++
		case "acc":
			bc.accumulator += code.argument
			line++
		case "jmp":
			line += code.argument
		}
	}

	fmt.Println(bc.accumulator)
}
