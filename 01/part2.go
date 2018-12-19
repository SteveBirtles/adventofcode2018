package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

func contains(x int, col []int) bool {
	for _, y := range col {
		if y == x {
			return true
		}
	}
	return false
}

func main() {

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Can't open input file")
	}

	data := strings.Split(string(input), "\n")

	freqs := make([]int, 0)
	frequency := 0

outer:
	for {
		for _, item := range data {
			x, err := strconv.Atoi(item)
			if err == nil {
				frequency += x
				if contains(frequency, freqs) {
					fmt.Println(frequency)
					break outer
				}
				freqs = append(freqs, frequency)
			}
		}
	}

}
