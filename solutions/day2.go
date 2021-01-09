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

	c, err := ioutil.ReadFile("../inputs/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(c), "\n")
	for i, v := range lines {
		for ii, vv := range lines {
			for iii, vvv := range lines {
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
				k1 := fmt.Sprintf("%s+%s+%s", v, vv, vvv)
				k2 := fmt.Sprintf("%s+%s+%s", v, vvv, vv)
				k3 := fmt.Sprintf("%s+%s+%s", vv, v, vvv)
				k4 := fmt.Sprintf("%s+%s+%s", vv, vvv, v)
				k5 := fmt.Sprintf("%s+%s+%s", vvv, vv, v)
				k6 := fmt.Sprintf("%s+%s+%s", vvv, v, vv)

				if i == ii || i == iii || m[k1] == true || m[k2] == true || m[k3] == true || m[k4] == true || m[k5] == true || m[k6] == true {
					continue
				} else {
					if iv+ivv+ivvv == 2020 {
						fmt.Printf("%d * %d * %d = %d \n", iv, ivv, ivvv, iv*ivv*ivvv)
						break
					} else {
						m[k1], m[k2], m[k3], m[k4], m[k5], m[k6] = true, true, true, true, true, true
					}
				}
			}
		}
	}
}
