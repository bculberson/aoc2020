package main

import (
	"bufio"
	"errors"
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

func run(instructions []instruction) (int, error) {
	history := make(map[int]bool)
	result := 0

	for ix := 0; ix <= len(instructions); ix++ {
		if ix == len(instructions) {
			return result, nil
		}
		if _, ok := history[ix]; ok {
			return result, errors.New("recursion")
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

	return result, errors.New("unknown")
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

	for ix := 0; ix < len(instructions); ix++ {
		cpy := make([]instruction, len(instructions))
		copy(cpy, instructions)
		if cpy[ix].op == "jmp" {
			cpy[ix].op = "nop"
		} else if cpy[ix].op == "nop" {
			cpy[ix].op = "jmp"
		} else {
			continue
		}
		result, err := run(cpy)
		if err == nil {
			fmt.Printf("result = %d %v\n", result, err)
			break
		}
	}
}
