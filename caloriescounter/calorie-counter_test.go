package caloriescounter_test

import (
	"catarinabombaca/advent-code-22/caloriescounter"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCases struct {
	Name           string
	Input          []int
	ExpectedResult int
}

func TestFindElfWithMostCalories(t *testing.T) {
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
			mostCalories := caloriescounter.FindElfWithMostCalories(test.Input)
			assert.Equal(t, test.ExpectedResult, mostCalories)
		})
	}
}
