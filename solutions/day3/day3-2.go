package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func treeCount(l [323]string, x, y int) int {
	var res = 0
	rl := len(l)
	for r := 0; r < rl; r++ {
		if r*y > rl {
			break
		}
		if string(l[r*y][r*x]) == "#" {
			res++
		}
	}
	return res
}
func main() {
	c, err := ioutil.ReadFile("./day3-2.txt")
	if err != nil {
		log.Fatal(err)
	}
	a := strings.Split(string(c), "\n")

	b := [323]string{}
	rl := len(a)
	for i, k := range a {
		// for o := 0; o < rl; o++ {
		// 	b[i] += k
		// }
		b[i] = strings.Repeat(k, rl)
	}

	fmt.Println(treeCount(b, 1, 1), treeCount(b, 3, 1), treeCount(b, 5, 1), treeCount(b, 7, 1), treeCount(b, 1, 2))
	fmt.Println(treeCount(b, 1, 1) * treeCount(b, 3, 1) * treeCount(b, 5, 1) * treeCount(b, 7, 1) * treeCount(b, 1, 2))
}
