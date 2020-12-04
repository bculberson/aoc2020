package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func validateByr(passport map[string]string) bool {
	if byrVal, ok := passport["byr"]; ok {
		byr, err := strconv.Atoi(byrVal)
		if len(byrVal) != 4 || err != nil || byr < 1920 || byr > 2002 {
			return false
		}
	} else {
		return false
	}
	return true
}

func validateIyr(passport map[string]string) bool {
	if iyrVal, ok := passport["iyr"]; ok {
		iyr, err := strconv.Atoi(iyrVal)
		if len(iyrVal) != 4 || err != nil || iyr < 2010 || iyr > 2020 {
			return false
		}
	} else {
		return false
	}
	return true
}

func validateEyr(passport map[string]string) bool {
	if eyrVal, ok := passport["eyr"]; ok {
		eyr, err := strconv.Atoi(eyrVal)
		if len(eyrVal) != 4 || err != nil || eyr < 2020 || eyr > 2030 {
			return false
		}
	} else {
		return false
	}
	return true
}

func validateHgt(passport map[string]string) bool {
	if hgtVal, ok := passport["hgt"]; ok {
		if strings.HasSuffix(hgtVal, "cm") {
			hgtCm := strings.TrimSuffix(hgtVal, "cm")
			hgt, err := strconv.Atoi(hgtCm)
			if len(hgtCm) != 3 || err != nil || hgt < 150 || hgt > 193 {
				return false
			}
		} else if strings.HasSuffix(hgtVal, "in") {
			hgtIn := strings.TrimSuffix(hgtVal, "in")
			hgt, err := strconv.Atoi(hgtIn)
			if len(hgtIn) != 2 || err != nil || hgt < 59 || hgt > 76 {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
	return true
}

func validateHcl(passport map[string]string) bool {
	if hclVal, ok := passport["hcl"]; ok {
		if len(hclVal) == 7 && strings.HasPrefix(hclVal, "#") {
			for i := 1; i < len(hclVal); i++ {
				if !((hclVal[i] >= '0' && hclVal[i] <= '9') || (hclVal[i] >= 'a' && hclVal[i] <= 'f')) {
					return false
				}
			}
		} else {
			return false
		}
	} else {
		return false
	}
	return true
}

func validateEcl(passport map[string]string) bool {
	if eclVal, ok := passport["ecl"]; ok {
		if !(eclVal == "amb" || eclVal == "blu" || eclVal == "brn" || eclVal == "gry" || eclVal == "grn" || eclVal == "hzl" || eclVal == "oth") {
			return false
		}
	} else {
		return false
	}
	return true
}

func validatePid(passport map[string]string) bool {
	if pidVal, ok := passport["pid"]; ok {
		if len(pidVal) != 9 {
			return false
		}
		for i := 0; i < len(pidVal); i++ {
			if pidVal[i] < '0' || pidVal[i] > '9' {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func isValid(passport map[string]string) bool {
	return validateByr(passport) && validateEcl(passport) && validateEyr(passport) &&
		validateHcl(passport) && validateHgt(passport) && validateIyr(passport) && validatePid(passport)
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
