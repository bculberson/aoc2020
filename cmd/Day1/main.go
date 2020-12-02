package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func findSummandsInSliceTo(lines []int, startOffset int, sum int) (int, int, error) {
	nextOffset := startOffset + 1
	firstSummand := lines[startOffset]
	for i := startOffset + 1; i < len(lines)-1; i++ {
		secondSummand := lines[i]
		if firstSummand+secondSummand == sum {
			return firstSummand, secondSummand, nil
		}
	}
	if nextOffset < len(lines)-1 {
		return findSummandsInSliceTo(lines, nextOffset+1, sum)
	}
	return 0, 0, errors.New("Not Found")
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

	a, b, err := findSummandsInSliceTo(lines, 0, 2020)
	if err != nil {
		log.Fatal("Summands not found!")
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
