package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	c, err := ioutil.ReadFile("../inputs/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	a := strings.Split(string(c), "\n")
	rl := len(a)
	b := [323]string{}
	for i, k := range a {
		// for o := 0; o < rl; o++ {
		// 	b[i] += k
		// }
		b[i] = strings.Repeat(k, rl)
	}
	var res = 0
	for r := 0; r < rl; r++ {
		if string(b[r][r*3]) == "#" {
			res++
		}
	}
	fmt.Println(res)
}
