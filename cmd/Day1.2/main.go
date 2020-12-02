package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func findSummandsInSliceTo(target []int, data []int, sum int) (int, int, int, error) {
	for i := 0; i < len(data); i++ {
		newTarget := make([]int, len(target)+1)
		copy(newTarget, target)
		newTarget[len(target)] = data[i]
		newData := data[i+1:]
		if len(newTarget) == 3 {
			if (newTarget[0] + newTarget[1] + newTarget[2]) == sum {
				return newTarget[0], newTarget[1], newTarget[2], nil
			}
		} else if len(newTarget) < 3 {
			a, b, c, err := findSummandsInSliceTo(newTarget, newData, sum)
			if err == nil {
				return a, b, c, nil
			}
		}
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

	a, b, c, err := findSummandsInSliceTo([]int{}, lines, 2020)
	if err != nil {
		log.Fatal("Summands not found!")
	}
	fmt.Printf("%d * %d * %d = %d\n", a, b, c, a*b*c)
}
