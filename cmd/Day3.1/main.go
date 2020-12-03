package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func runXY(lines []string, dx int, dy int) (int, error) {
	var trees, x int
	width := len(lines[0])
	for y := 0; y < len(lines); y = y + dy {
		if lines[y][x%width] == byte('#') {
			trees++
		}
		x = x + dx
	}
	return trees, nil
}

func doRunXY(lines []string, dx int, dy int) int {
	trees, err := runXY(lines, dx, dy)
	if err != nil {
		log.Fatalf("Error in Run: %s", err)
	}
	fmt.Printf("right %d, down %d = %d\n", dx, dy, trees)
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

	answer := 1
	answer = answer * doRunXY(lines, 1, 1)
	answer = answer * doRunXY(lines, 3, 1)
	answer = answer * doRunXY(lines, 5, 1)
	answer = answer * doRunXY(lines, 7, 1)
	answer = answer * doRunXY(lines, 1, 2)
	fmt.Printf("answer = %d\n", answer)
}
