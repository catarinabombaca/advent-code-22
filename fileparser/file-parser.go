package fileparser

import (
	"os"
	"strconv"
	"strings"
)

func ParseToListOfCalories(filePath string) ([]int, error) {
	var (
		elfsCaloriesList []int
		totalElfCalories int
	)

	caloriesList, err := extractCalories(filePath)
	if err != nil {
		return nil, err
	}

	for food, amountOfCalories := range caloriesList {
		if isLastFoodCaloriesForElf(amountOfCalories) {
			elfsCaloriesList = append(elfsCaloriesList, totalElfCalories)
			totalElfCalories = 0
			continue
		}

		caloriesAsInt, err := strconv.Atoi(amountOfCalories)
		if err != nil {
			return nil, err
		}

		totalElfCalories += caloriesAsInt

		if isLastFoodCaloriesFromList(food, caloriesList) {
			elfsCaloriesList = append(elfsCaloriesList, totalElfCalories)
		}
	}

	return elfsCaloriesList, nil
}

func extractCalories(filepath string) ([]string, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(file), "\n"), nil
}

func isLastFoodCaloriesForElf(amount string) bool {
	return amount == ""
}

func isLastFoodCaloriesFromList(calories int, list []string) bool {
	return calories == len(list)-1
}
