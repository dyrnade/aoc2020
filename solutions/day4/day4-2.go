package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func valid(l []string, s string) bool {
	for _, v := range l {
		x := strings.Split(v, ":")[0]
		if x == s {
			return true
		}
	}
	return false
}

func main() {
	c, err := ioutil.ReadFile("./day4-1.txt")
	if err != nil {
		log.Fatal(err)
	}

	var validFields = []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		// "cid",
	}

	// Firstly, merge all the lines that define a valid or not valid passport
	var passports = []string{}
	var startPoint, endPoint int
	ss := strings.Split(string(c), "\n")
	for k, v := range ss {
		endPoint = k
		if v == "" {
			passports = append(passports, strings.Join(ss[startPoint:endPoint], " "))
			startPoint = endPoint + 1
			continue
		}

		if k == len(ss)-1 {
			passports = append(passports, strings.Join(ss[startPoint:len(ss)], " "))
		}
	}

	validPassportsCount := 0
	for _, passport := range passports {
		valid := true
		passportFields := strings.Split(strings.TrimSpace(passport), " ")
		if len(passportFields) < len(validFields) {
			continue
		}
		for _, f := range validFields {
			if !find(passportFields, f) {
				valid = false
			}
		}
		if valid {
			validPassportsCount++
		}
	}
	fmt.Printf("Valid passport count: %d", validPassportsCount)
}
