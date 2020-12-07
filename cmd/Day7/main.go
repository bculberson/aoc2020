package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type bagContainer struct {
	Number int
	Color  string
}

func parseLine(line string) (string, []*bagContainer) {
	words := strings.Split(line, " bags contain ")
	if words[1] == "no other bags." {
		return words[0], nil
	}
	result := make([]*bagContainer, 0)
	for _, bag := range strings.Split(words[1], ", ") {
		number := int(byte(bag[0]) - 8)
		color := strings.Split(bag, " bag")[0][2:]
		result = append(result, &bagContainer{Number: number, Color: color})
	}
	return words[0], result
}

func contains(bagMappings map[string][]*bagContainer, color string, containers map[string]bool) {
	for k := range bagMappings {
		for _, bc := range bagMappings[k] {
			if bc.Color == color {
				containers[k] = true
				contains(bagMappings, k, containers)
			}
		}
	}
}

func main() {
	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error in Scanner: %s", err)
	}

	bagMappings := make(map[string][]*bagContainer)
	for _, line := range lines {
		k, vs := parseLine(line)
		bagMappings[k] = vs
	}

	containers := make(map[string]bool)
	contains(bagMappings, "shiny gold", containers)

	fmt.Printf("count = %v\n", len(containers))
}
