package listutil

func Max(list []int) int {
	max := list[0]

	for _, value := range list {
		if value > max {
			max = value
		}
	}

	return max
}
