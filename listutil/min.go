package listutil

func Min(list []int) int {
	min := list[0]

	for _, value := range list {
		if value < min {
			min = value
		}
	}

	return min
}
