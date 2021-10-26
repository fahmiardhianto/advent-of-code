package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	rs := newRuleSet()
	for scanner.Scan() {
		rs.parseLine(scanner.Text())
	}

	result := rs.findBagsContain("shiny gold")
	fmt.Println(len(result))
}

type ruleset struct {
	bags map[string][]string
}

func newRuleSet() ruleset {
	return ruleset{bags: make(map[string][]string)}
}

func (r *ruleset) parseLine(line string) {
	rule := strings.Split(line, " bags contain ")
	bag := rule[0]
	contents := strings.Split(rule[1], ", ")
	for _, c := range contents {
		c = strings.TrimLeftFunc(c, func(r rune) bool {
			return unicode.IsNumber(r)
		})
		c = strings.TrimSuffix(c, ".")
		c = strings.TrimSuffix(c, "bag")
		c = strings.TrimSuffix(c, "bags")
		c = strings.TrimSpace(c)
		r.bags[c] = append(r.bags[c], bag)
	}
}

func (r *ruleset) findBagsContain(content string) map[string]bool {
	result := make(map[string]bool)

	stacks := r.bags[content]
	var bag string

	for len(stacks) > 0 {
		bag, stacks = stacks[len(stacks)-1], stacks[:len(stacks)-1]
		result[bag] = true
		for _, c := range r.bags[bag] {
			stacks = append(stacks, c)
		}
	}

	return result
}
