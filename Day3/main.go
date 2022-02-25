package main

import (
	"fmt"
	"github.com/echojc/aocutil"
	"log"
)

func main() {
	input, err := aocutil.NewInputFromFile("/home/data2020/Golang/Advent-of-code-2021/Day2/session_id")

	if err != nil {
		log.Fatal(err)
	}
	data, err := input.Strings(2021, 3)
	if err != nil {
		log.Fatal(err)
	}

	var bits = make([]int, len(data[0]))

	for _, n := range data {
		for i := range n {
			if n[i] == '1' {
				bits[i]++
			}
		}
	}
	g := 0b010010010110
	e := 0b101101101001

	fmt.Println(bits, len(data))
	fmt.Println(g, e, g*e)
}
