package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	op  string
	arg int
}

func run(instructions []instruction) int {
	history := make(map[int]bool)
	result := 0

	for ix := 0; true; ix++ {
		if _, ok := history[ix]; ok {
			return result
		}
		history[ix] = true

		if instructions[ix].op == "nop" {
			continue
		} else if instructions[ix].op == "acc" {
			result += instructions[ix].arg
		} else if instructions[ix].op == "jmp" {
			ix = ix + instructions[ix].arg - 1
		}
	}

	return result
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

	instructions := make([]instruction, len(lines))
	for ix, line := range lines {
		words := strings.Split(line, " ")
		i, err := strconv.Atoi(words[1])
		if err != nil {
			log.Fatalf("Error %s in int conversion of: %s", err, words[1])
		}
		instructions[ix] = instruction{op: words[0], arg: i}
	}

	result := run(instructions)
	fmt.Printf("result = %d\n", result)
}
