package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	c, err := ioutil.ReadFile("../inputs/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(c), "\n")
	for _, v := range lines {
		for _, vv := range lines {
			for _, vvv := range lines {
				iv, err := strconv.Atoi(v)
				if err != nil {
					log.Fatal(err)
				}
				ivv, err := strconv.Atoi(vv)
				if err != nil {
					log.Fatal(err)
				}
				ivvv, err := strconv.Atoi(vvv)
				if err != nil {
					log.Fatal(err)
				}
				if iv+ivv+ivvv == 2020 {
					fmt.Printf("%d * %d * %d = %d \n", iv, ivv, ivvv, iv*ivv*ivvv)
					os.Exit(0)
				} else {
					continue
				}

			}
		}
	}
}
