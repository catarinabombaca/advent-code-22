package main

import (
	"catarinabombaca/advent-code-22/calories"
	"catarinabombaca/advent-code-22/fileparser"
	"fmt"
	"log"
)

func main() {
	//Day 1
	//Part 1
	elvesCalories, err := fileparser.ParseToListOfCalories("inputs/input.txt")
	if err != nil {
		log.Fatalf("error parsing the file: %v", err)
	}

	caloriesCounter := calories.NewCounter(elvesCalories)
	mostCalories := caloriesCounter.FindElfWithMostCalories()

	fmt.Printf("Day 1, Part 1: %v \n", mostCalories)

	//Part 2
	caloriesCounter.SortElvesCaloriesFromHighToLow()
	totalCaloriesForTop3Elves := caloriesCounter.SumFirstNElvesCalories(3)

	fmt.Printf("Day 1, Part 2 %v \n", totalCaloriesForTop3Elves)
}
