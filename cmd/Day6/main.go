package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getGroupsAnswers(lines []string) []map[byte]bool {
	result := make([]map[byte]bool, 0)
	groupAnswers := make(map[byte]bool)
	for _, line := range lines {
		if len(line) > 0 {
			for i := 0; i < len(line); i++ {
				groupAnswers[line[i]] = true
			}
		} else {
			result = append(result, groupAnswers)
			groupAnswers = make(map[byte]bool)
		}
	}
	result = append(result, groupAnswers)
	return result
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

	sum := 0
	groupsAnswers := getGroupsAnswers(lines)
	for _, groupAnswers := range groupsAnswers {
		sum += len(groupAnswers)
	}

	fmt.Printf("sum = %d\n", sum)
}
