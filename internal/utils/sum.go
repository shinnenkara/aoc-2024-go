package utils

func Sum(values []int) int {
	result := 0
	for i := 0; i < len(values); i++ {
		result += values[i]
	}

	return result
}
