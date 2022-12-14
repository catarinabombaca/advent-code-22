package caloriescounter

func FindElfWithMostCalories(elfsCalories []int) int {
	var mostCalories int

	for _, calories := range elfsCalories {
		if calories >= mostCalories {
			mostCalories = calories
		}
	}

	return mostCalories
}
