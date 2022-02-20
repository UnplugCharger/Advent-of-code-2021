package main

import (
	"fmt"
	"github.com/echojc/aocutil"
	"log"
)

func main() {
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
