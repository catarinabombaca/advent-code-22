package main

import (
	"catarinabombaca/advent-code-22/caloriescounter"
	"catarinabombaca/advent-code-22/fileparser"
	"fmt"
	"log"
)

func main() {
	//Day 1
	//Part 1
	elfsCalories, err := fileparser.ParseToListOfCalories("inputs/input.txt")
	if err != nil {
		log.Fatalf("error parsing the file: %v", err)
	}

	mostCalories := caloriescounter.FindElfWithMostCalories(elfsCalories)
	fmt.Printf("Day 1, Part 1: %v", mostCalories)
}
