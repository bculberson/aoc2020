package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func run(lines []string) int {
	var trees, y int
	width := len(lines[0])
	for i := 0; i < len(lines); i++ {
		if lines[i][y%width] == byte('#') {
			trees++
		}
		y = y + 3
	}
	return trees
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

	trees := run(lines)
	fmt.Printf("%d\n", trees)
}
