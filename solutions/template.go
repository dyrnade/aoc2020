package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	c, err := ioutil.ReadFile("../inputs/day.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Contents: %s", c)
}
