package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	length := max - min
	result := make([]int, length)
	for i := range result {
		result[i] = min + i
	}
	return result
}
