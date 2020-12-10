package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func runChain(lines []int) (int, int) {
	var num1, num3 int

	for x := 1; x < len(lines); x++ {
		if lines[x]-lines[x-1] == 1 {
			num1++
		} else if lines[x]-lines[x-1] == 3 {
			num3++
		}
	}
	return num1, num3
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

	count1, count3 := runChain(lines)
	count3++
	fmt.Printf("Found %d of 1s and %d of 3s, result: %d\n", count1, count3, count1*count3)
}
