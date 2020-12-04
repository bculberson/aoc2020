package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parse(lines []string) []map[string]string {
	result := make([]map[string]string, 0)
	passport := make((map[string]string))
	for _, line := range lines {
		kvs := strings.Split(line, " ")
		for _, kv := range kvs {
			kv := strings.Split(kv, ":")
			if len(kv) < 2 {
				result = append(result, passport)
				passport = make((map[string]string))
			} else {
				passport[kv[0]] = kv[1]
			}
		}
	}
	result = append(result, passport)
	return result
}

func isValid(passport map[string]string) bool {
	if _, ok := passport["byr"]; ok {
		if _, ok := passport["iyr"]; ok {
			if _, ok := passport["eyr"]; ok {
				if _, ok := passport["hgt"]; ok {
					if _, ok := passport["hcl"]; ok {
						if _, ok := passport["ecl"]; ok {
							if _, ok := passport["pid"]; ok {
								return true
							}
						}
					}
				}
			}
		}
	}
	return false
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

	passports := parse(lines)
	ok := 0
	for _, p := range passports {
		if isValid(p) {
			ok++
		}
	}
	fmt.Printf("%d\n", ok)
}
