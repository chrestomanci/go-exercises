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

// Similar but takes a vardic list instead of an array
func MinArgs(args ...int) int {
	first := true
	var min int

	for _, value := range args {
		if first {
			min = value
			first = false
		} else if value < min {
			min = value
		}
	}

	return min
}
