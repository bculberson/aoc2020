package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func runChain(lines []int) int {
	result := map[int]int{0: 1}
	for i := 1; i < len(lines); i++ {
		v := lines[i]
		result[v] = result[v-1] + result[v-2] + result[v-3]
	}
	return result[lines[len(lines)-1]]
}

func main() {
	var lines []int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Error %s in int conversion of: %s", err, scanner.Text())
		}
		lines = append(lines, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error in Scanner: %s", err)
	}

	lines = append(lines, 0)
	sort.Ints(lines)
	lines = append(lines, lines[len(lines)-1]+3)

	count := runChain(lines)
	fmt.Printf("Found %d combinations\n", count)
}
