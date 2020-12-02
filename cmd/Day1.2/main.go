package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func findSummandsInSliceTo(lines []int, startOffset int, sum int) (int, int, int, error) {
	nextOffset := startOffset + 1
	firstSummand := lines[startOffset]
	for i := startOffset + 1; i < len(lines); i++ {
		secondSummand := lines[i]
		for x := 0; x < len(lines); x++ {
			thirdSummand := lines[x]
			if firstSummand+secondSummand+thirdSummand == sum {
				return firstSummand, secondSummand, thirdSummand, nil
			}
		}
	}
	if nextOffset < len(lines)-1 {
		return findSummandsInSliceTo(lines, nextOffset, sum)
	}
	return 0, 0, 0, errors.New("Not Found")
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

	a, b, c, err := findSummandsInSliceTo(lines, 0, 2020)
	if err != nil {
		log.Fatal("Summands not found!")
	}
	fmt.Printf("FOUND %d * %d * %d = %d\n", a, b, c, a*b*c)
}
