package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func summandsFound(items []int, start int, end int, sum int) bool {
	for i := start; i < end; i++ {
		for x := i; x < end; x++ {
			if items[i]+items[x] == sum {
				return true
			}
		}
	}
	return false
}

func main() {

	argsWithoutProg := os.Args[1:]
	lookback, err := strconv.Atoi(argsWithoutProg[0])
	if err != nil {
		log.Fatalf("Error %s in int conversion of: %s", err, argsWithoutProg[0])
	}

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

	for i := lookback; i < len(lines); i++ {
		if summandsFound(lines, i-lookback, i, lines[i]) {
			continue
		}
		fmt.Printf("%d not found \n", lines[i])
		break
	}

}
