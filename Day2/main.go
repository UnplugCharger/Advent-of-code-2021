package main

import (
	"fmt"
	"github.com/echojc/aocutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := aocutil.NewInputFromFile("session_id")

	if err != nil {
		log.Fatal(err)
	}
	data, err := input.Strings(2021, 2)
	if err != nil {
		log.Fatal(err)
	}
	type instructionsSet struct {
		instructions string
		value        int
	}
	//Read the data in the file into an array

	var instructions []instructionsSet

	//loop  data and append the instructions
	for _, s := range data {
		// split line at space
		splits := strings.Split(s, " ")
		v, err := strconv.Atoi(splits[1])
		if err != nil {
			log.Fatal(err)
		}
		instructions = append(instructions, instructionsSet{splits[0], v})

		h := 0
		d := 0

		for _, i := range instructions {
			switch i.instructions {
			case "forward":
				h += i.value
			case "up":
				d -= i.value
			case "down":
				d += i.value
			default:
				log.Fatal(i)

			}
		}
		fmt.Println(h, d, h*d)
	}

}
