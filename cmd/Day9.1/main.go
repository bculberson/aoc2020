package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func summandsFound(items []int, length int, sum int) (bool, int, int) {
	for i := 0; i < len(items)-length; i++ {
		tsum := 0
		min := 88881063549210
		max := 0
		for x := i; x < i+length; x++ {
			tsum += items[x]
			if items[x] < min {
				min = items[x]
			}
			if items[x] > max {
				max = items[x]
			}
		}
		if tsum == sum {
			return true, min, max
		}
	}
	return false, 0, 0
}

func main() {

	argsWithoutProg := os.Args[1:]
	matchValue, err := strconv.Atoi(argsWithoutProg[0])
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

	for i := 2; i <= len(lines); i++ {
		found, min, max := summandsFound(lines, i, matchValue)
		if found {
			fmt.Printf("%d found, min %d, max: %d, result: %d \n", matchValue, min, max, min+max)
			break
		}
		continue
	}

}
