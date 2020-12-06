package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getRow(line string) int {
	urow := 127
	lrow := 0
	for x := 0; x < 7; x++ {
		if line[x] == 'F' {
			urow = urow - (urow-lrow)/2 - 1
		} else {
			lrow = lrow + (urow-lrow)/2 + 1
		}
	}
	return lrow
}

func getCol(line string) int {
	ucol := 7
	lcol := 0
	for x := 7; x < 10; x++ {
		if line[x] == 'L' {
			ucol = ucol - (ucol-lcol)/2 - 1
		} else {
			lcol = lcol + (ucol-lcol)/2 + 1
		}
	}
	return ucol
}

func getSeatInfo(line string) (int, int, int) {
	row := getRow(line)
	col := getCol(line)
	id := row*8 + col
	return row, col, id
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

	var seatID int
	for _, line := range lines {
		_, _, tSeatID := getSeatInfo(line)
		if tSeatID > seatID {
			seatID = tSeatID
		}
	}
	fmt.Printf("%d\n", seatID)
}
