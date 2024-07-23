package main

func sortWin(array []int) []int {
	max := 0
	low := 0
	high := 0
	for i, v := range array {
		if v < max {
			max = v
			high = i
		}
	}
	min := max
	for i := len(array) - 1; i >= 0; i-- {
		v := array[i]
		if v > min {
			min = v
			low = i
		}
	}

	return []int{low, high}
}
