package main

import (
	"regexp"
	"strconv"
	"strings"
	"testing"
)

func ValidatePassports(validateFn func(map[string]string) bool) func([]string) int {
	return func(input []string) int {
		validCount := 0

		passport := make(map[string]string)
		for _, line := range input {
			if len(line) == 0 {
				if validateFn(passport) {
					validCount++
				}
				passport = make(map[string]string)
				continue
			}
			for _, kv := range strings.Split(line, " ") {
				kvarr := strings.Split(kv, ":")
				passport[kvarr[0]] = kvarr[1]
			}
		}

		return validCount
	}
}

func hasAllFields(passport map[string]string) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, required := range requiredFields {
		if _, ok := passport[required]; !ok {
			return false
		}
	}
	return true
}

func validPassportFields(passport map[string]string) bool {
	if byr, err := strconv.Atoi(passport["byr"]); err != nil || byr < 1920 || byr > 2002 {
		return false
	}

	if iyr, err := strconv.Atoi(passport["iyr"]); err != nil || iyr < 2010 || iyr > 2020 {
		return false
	}

	if eyr, err := strconv.Atoi(passport["eyr"]); err != nil || eyr < 2020 || eyr > 2030 {
		return false
	}

	if strings.Contains(passport["hgt"], "cm") {
		if hgt, err := strconv.Atoi(strings.ReplaceAll(passport["hgt"], "cm", "")); err != nil || hgt < 150 || hgt > 193 {
			return false
		}
	} else if strings.Contains(passport["hgt"], "in") {
		if hgt, err := strconv.Atoi(strings.ReplaceAll(passport["hgt"], "in", "")); err != nil || hgt < 59 || hgt > 76 {
			return false
		}
	} else {
		return false
	}

	re := regexp.MustCompile(`#[0-9a-f]{6}`)
	if !re.MatchString(passport["hcl"]) {
		return false
	}

	if passport["ecl"] != "amb" && passport["ecl"] != "blu" && passport["ecl"] != "brn" && passport["ecl"] != "gry" && passport["ecl"] != "grn" && passport["ecl"] != "hzl" && passport["ecl"] != "oth" {
		return false
	}

	pidre := regexp.MustCompile(`[0-9]{9}`)
	if len(passport["pid"]) != 9 || !pidre.MatchString(passport["pid"]) {
		return false
	}

	return true
}

func TestAOCDay4(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		fn       func([]string) int
		expected int
	}{
		{name: "part1sample", input: []string{
			"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
			"byr:1937 iyr:2017 cid:147 hgt:183cm",
			"",
			"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
			"hcl:#cfa07d byr:1929",
			"",
			"hcl:#ae17e1 iyr:2013",
			"eyr:2024",
			"ecl:brn pid:760753108 byr:1931",
			"hgt:179cm",
			"",
			"hcl:#cfa07d eyr:2025 pid:166559648",
			"iyr:2011 ecl:brn hgt:59in",
		}, fn: ValidatePassports(hasAllFields), expected: 2},
		{name: "part1", input: readStringsInput("day4"), fn: ValidatePassports(hasAllFields), expected: 204},
		{name: "part2sampleInvalid", input: []string{
			"eyr:1972 cid:100",
			"hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926",
			"",
			"iyr:2019",
			"hcl:#602927 eyr:1967 hgt:170cm",
			"ecl:grn pid:012533040 byr:1946",
			"",
			"hcl:dab227 iyr:2012",
			"ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277",
			"",
			"hgt:59cm ecl:zzz",
			"eyr:2038 hcl:74454a iyr:2023",
			"pid:3556412378 byr:2007",
			"",
		}, fn: ValidatePassports(validPassportFields), expected: 0},
		{name: "part2sampleValid", input: []string{
			"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980",
			"hcl:#623a2f",
			"",
			"eyr:2029 ecl:blu cid:129 byr:1989",
			"iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm",
			"",
			"hcl:#888785",
			"hgt:164cm byr:2001 iyr:2015 cid:88",
			"pid:545766238 ecl:hzl",
			"eyr:2022",
			"",
			"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719",
			"",
		}, fn: ValidatePassports(validPassportFields), expected: 4},
		{name: "part2", input: readStringsInput("day4"), fn: ValidatePassports(validPassportFields), expected: 179},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := tt.fn(tt.input); res != tt.expected {
				t.Fatalf("Expected: %d, Actual: %d", tt.expected, res)
			}
		})
	}
}
