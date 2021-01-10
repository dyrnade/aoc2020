package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	c, err := ioutil.ReadFile("./day2-1.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(c), "\n")

	var validPws int
	for _, ln := range lines {
		line := strings.Split(string(ln), " ")
		l := strings.Split(string(line[0]), "-")
		min, err := strconv.Atoi(l[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(l[1])
		if err != nil {
			log.Fatal(err)
		}
		letter := string(line[1])[0]
		pw := string(line[2])
		var c int
		for _, v := range pw {
			if byte(v) == letter {
				c++
			}
		}
		if c >= min && c <= max {
			validPws++
		}
		fmt.Printf("line: %s\n count: %d\n", line, c)

	}
	fmt.Println(validPws)
}
