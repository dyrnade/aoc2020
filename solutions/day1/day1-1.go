package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {

	m := make(map[string]bool)

	c, err := ioutil.ReadFile("../inputs/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(c), "\n")
	for i, v := range lines {
		for ii, vv := range lines {
			iv, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			ivv, err := strconv.Atoi(vv)
			if err != nil {
				log.Fatal(err)
			}

			k1 := fmt.Sprintf("%s+%s", v, vv)
			k2 := fmt.Sprintf("%s+%s", vv, v)
			if i == ii || m[k1] == true || m[k2] == true {
				continue
			} else {
				if iv+ivv == 2020 {
					fmt.Printf("%d * %d = %d \n", iv, ivv, iv*ivv)
					break
				} else {
					m[k1] = true
					m[k2] = true
				}
			}
		}
	}
}
