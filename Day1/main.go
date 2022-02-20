package main

import (
	"fmt"
	"github.com/echojc/aocutil"
	"log"
)

func main() {
	all()
}
func depthCalc() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}
	data, err := input.Ints(2021, 1)
	if err != nil {
		log.Fatal(err)
	}
	// set the depth to the first value of data and initialize the count to zero
	depth := data[0]
	count := 0
	// Loop through data and if data is greater than previous increment count
	for i := 1; i < len(data); i++ {
		if data[i] > depth {
			count++
		}
		depth = data[i]
	}
	fmt.Println(count)
}
func depthCalcSlide() {
	input, err := aocutil.NewInputFromFile("session_id")
	if err != nil {
		log.Fatal(err)
	}
	data, err := input.Ints(2021, 1)
	if err != nil {
		log.Fatal(err)
	}
	depth := data[0] + data[1] + data[2]

	count := 0
	for i := 1; i < len(data)-2; i++ {
		curDepth := data[i] + data[i+1] + data[i+2]
		if curDepth > depth {
			count++
		}
		depth = curDepth

	}
	fmt.Println(count)
}
func all() {
	depthCalc()
	depthCalcSlide()
}
