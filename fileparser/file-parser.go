package fileparser

import (
	"os"
	"strconv"
	"strings"
)

func ParseToListOfCalories(filePath string) ([]int, error) {
	var (
		elvesCaloriesList []int
		totalElfCalories  int
	)

	caloriesList, err := extractCalories(filePath)
	if err != nil {
		return nil, err
	}

	for food, amountOfCalories := range caloriesList {
		if isLastFoodCaloriesForElf(amountOfCalories) {
			elvesCaloriesList = append(elvesCaloriesList, totalElfCalories)
			totalElfCalories = 0
			continue
		}

		caloriesAsInt, err := strconv.Atoi(amountOfCalories)
		if err != nil {
			return nil, err
		}

		totalElfCalories += caloriesAsInt

		if isLastFoodCaloriesFromList(food, caloriesList) {
			elvesCaloriesList = append(elvesCaloriesList, totalElfCalories)
		}
	}

	return elvesCaloriesList, nil
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
