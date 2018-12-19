package main

import (
	"io/ioutil"
	"strings"
	"strconv"
	"fmt"
)

func main() {

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Can't open input file")
	}

	data := strings.Split(string(input), "\n")

	frequency := 0
	for _, item := range data {
		x, err := strconv.Atoi(item)
		if err == nil {
			frequency += x
		}
	}

	fmt.Println(frequency)

}
