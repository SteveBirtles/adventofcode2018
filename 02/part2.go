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

	outer:
	for i, item1 := range data {
		fmt.Print(".")
		for j, item2 := range data {
			if i == j { continue }
			diffs, diffat := 0, 0
			for k := range item1 {
				u := item1[k]
				v := item2[k]
				if u != v {
					diffs++
					diffat = k
				}
			}
			if diffs == 1 {
				fmt.Println()
				fmt.Println(item1)
				fmt.Println(item2)
				fmt.Printf("%s%s", item1[:diffat], item1[diffat+1:])

				break outer
			}
		}
	}

}


