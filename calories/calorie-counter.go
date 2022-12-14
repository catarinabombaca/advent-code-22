package calories

import "sort"

type CaloriesCounter struct {
	elvesCalories []int
}

func NewCounter(calories []int) *CaloriesCounter {
	return &CaloriesCounter{elvesCalories: calories}
}

func (cc *CaloriesCounter) FindElfWithMostCalories() int {
	var mostCalories int

	for _, calories := range cc.elvesCalories {
		if calories >= mostCalories {
			mostCalories = calories
		}
	}

	return mostCalories
}

func (cc *CaloriesCounter) SortElvesCaloriesFromHighToLow() []int {
	sort.Sort(sort.Reverse(sort.IntSlice(cc.elvesCalories)))

	return cc.elvesCalories
}

func (cc *CaloriesCounter) SumFirstNElvesCalories(n int) int {
	var sum int
	for _, calories := range cc.elvesCalories[:n] {
		sum += calories
	}

	return sum
}
