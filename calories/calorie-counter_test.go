package calories_test

import (
	"catarinabombaca/advent-code-22/calories"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindElfWithMostCalories(t *testing.T) {
	type testCases struct {
		Name           string
		Input          []int
		ExpectedResult int
	}

	tests := []testCases{
		{
			Name:           "Given a list of calories for each elf, when we find the elf with most calories, then the value is that elf's calories is returned",
			Input:          []int{1000, 10000, 3000, 4000, 6000},
			ExpectedResult: 10000,
		},
		{
			Name:           "Given a empty list of calories for each elf, when we find the elf with most calories, then the value 0 calories is returned",
			Input:          []int{},
			ExpectedResult: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			caloriesCounter := calories.NewCounter(test.Input)
			mostCalories := caloriesCounter.FindElfWithMostCalories()

			assert.Equal(t, test.ExpectedResult, mostCalories)
		})
	}
}

func TestSortElvesCaloriesFromHighToLow(t *testing.T) {
	type testCases struct {
		Name           string
		Input          []int
		ExpectedResult []int
	}

	tests := []testCases{
		{
			Name:           "Given a list of calories for each elf, when we sort the list from high to low, then an ordered list of calories is returned",
			Input:          []int{1000, 10000, 3000, 4000, 6000},
			ExpectedResult: []int{10000, 6000, 4000, 3000, 1000},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			caloriesCounter := calories.NewCounter(test.Input)
			sortCalories := caloriesCounter.SortElvesCaloriesFromHighToLow()

			assert.Equal(t, test.ExpectedResult, sortCalories)
		})
	}
}

func TestSumFirstNElvesCalories(t *testing.T) {
	type testCases struct {
		Name           string
		Input          []int
		N              int
		ExpectedResult int
	}

	tests := []testCases{
		{
			Name:           "Given a list of calories for each elf, when we sum each elf calories, then the total sum is returned",
			Input:          []int{1000, 10000, 3000, 4000, 6000},
			N:              3,
			ExpectedResult: 14000,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			caloriesCounter := calories.NewCounter(test.Input)
			totalCalories := caloriesCounter.SumFirstNElvesCalories(test.N)

			assert.Equal(t, test.ExpectedResult, totalCalories)
		})
	}
}
