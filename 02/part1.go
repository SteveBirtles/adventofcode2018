package main

import (
	"io/ioutil"
	"strings"
	"fmt"
)

func main() {

	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Can't open input file")
	}

	data := strings.Split(string(input), "\n")

	pairs, triples := 0, 0

	for _, item := range data {
		chars := make(map[string]int, 0)
		for _, char := range item {
			if count, ok := chars[string(char)]; ok {
				chars[string(char)] = count + 1
			} else {
				chars[string(char)] = 1
			}
		}
		fmt.Println(chars)
		for _, v := range chars {
			if v == 2 {
				pairs++
				break
			}
		}
		for _, v := range chars {
			if v == 3 {
				triples++
				break
			}
		}
	}

	fmt.Printf("%d Twos and %d Threes gives %d Checksum\n", pairs, triples, pairs*triples)

}


