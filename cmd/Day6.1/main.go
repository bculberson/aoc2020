package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getGroupsAnswers(lines []string) ([]map[byte]int, []int) {
	result := make([]map[byte]int, 0)
	groupsSizes := make([]int, 0)

	groupAnswers := make(map[byte]int)
	groupSize := 0
	for _, line := range lines {
		if len(line) > 0 {
			groupSize++
			for i := 0; i < len(line); i++ {
				if _, found := groupAnswers[line[i]]; found {
					groupAnswers[line[i]] = groupAnswers[line[i]] + 1
				} else {
					groupAnswers[line[i]] = 1
				}
			}
		} else {
			result = append(result, groupAnswers)
			groupsSizes = append(groupsSizes, groupSize)
			groupAnswers = make(map[byte]int)
			groupSize = 0
		}
	}
	result = append(result, groupAnswers)
	groupsSizes = append(groupsSizes, groupSize)
	return result, groupsSizes
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

	var sum, ix int
	groupsAnswers, groupsSizes := getGroupsAnswers(lines)
	for _, groupAnswers := range groupsAnswers {
		for _, q := range groupAnswers {
			if q == groupsSizes[ix] {
				sum++
			}
		}
		ix++
	}

	fmt.Printf("sum = %d\n", sum)
}
