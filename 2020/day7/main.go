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
	rs := newRuleSet()
	for scanner.Scan() {
		rs.parseLine(scanner.Text())
	}

	result := rs.findBagsContain("shiny gold")
	fmt.Println(len(result))

	fmt.Println(rs.countBags("shiny gold"))
}

type ruleset struct {
	bagsContain map[string][]string
	bagsContent map[string][]contentDetail
}

type contentDetail struct {
	color       string
	requiredQty int
}

func newRuleSet() ruleset {
	return ruleset{
		bagsContain: make(map[string][]string),
		bagsContent: make(map[string][]contentDetail),
	}
}

func (r *ruleset) parseLine(line string) {
	rule := strings.Split(line, " bags contain ")
	bag := rule[0]
	contents := strings.Split(rule[1], ", ")
	for _, c := range contents {
		c = strings.TrimSuffix(c, ".")
		c = strings.TrimSuffix(c, "bag")
		c = strings.TrimSuffix(c, "bags")
		c = strings.TrimSpace(c)
		content := strings.SplitN(c, " ", 2)
		bagColor := content[1]
		qty, _ := strconv.Atoi(content[0])

		if bagColor == "other" {
			continue
		}
		r.bagsContain[bagColor] = append(r.bagsContain[bagColor], bag)
		r.bagsContent[bag] = append(r.bagsContent[bag], contentDetail{color: bagColor, requiredQty: qty})
	}
}

func (r *ruleset) findBagsContain(content string) map[string]bool {
	result := make(map[string]bool)

	stacks := r.bagsContain[content]
	var bag string

	for len(stacks) > 0 {
		bag, stacks = stacks[len(stacks)-1], stacks[:len(stacks)-1]
		result[bag] = true
		for _, c := range r.bagsContain[bag] {
			stacks = append(stacks, c)
		}
	}

	return result
}

func (r *ruleset) countBags(bag string) int {
	contents := r.bagsContent[bag]
	if len(contents) == 0 {
		return 0
	}

	var sum int
	for _, s := range contents {
		bagQty := s.requiredQty
		contentQty := r.countBags(s.color)
		sum += bagQty*contentQty + bagQty
	}

	return sum
}
