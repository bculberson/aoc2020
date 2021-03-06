package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func isValid(line string) bool {
	words := strings.Split(line, " ")
	lens := strings.Split(words[0], "-")
	lower, _ := strconv.Atoi(lens[0])
	upper, _ := strconv.Atoi(lens[1])
	char := words[1][0]
	password := words[2]
	if (password[lower-1] == char || password[upper-1] == char) && !(password[lower-1] == char && password[upper-1] == char) {
		return true
	}
	return false
}

func main() {
	var lines []string
	var validPasswords int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if isValid(line) {
			validPasswords++
			lines = append(lines, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error in Scanner: %s", err)
	}

	fmt.Printf("%d valid passwords\n", validPasswords)
}
