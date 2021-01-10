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

		if (string(pw[min-1]) == string(letter)) && (string(pw[max-1]) == string(letter)) {
			continue
		} else if (string(pw[min-1]) != string(letter)) && (string(pw[max-1]) != string(letter)) {
			continue
		} else {
			// fmt.Printf("%s ==> %s -> min:%s(%d)   ==== max:%s(%d) \n", line, string(letter), string(pw[min-1]), min, string(pw[max-1]), max)
			validPws++
		}
	}
	fmt.Println("Total valid pws count:", validPws)
}
