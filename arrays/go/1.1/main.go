package main

import "fmt"

func reverse(s []int) []int {
	for i := len(s)/2 - 1; i >= 0; i-- {
		opp := len(s) - 1 - i
		s[i], s[opp] = s[opp], s[i]
	}
	return s
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	prefixes := []int{}
	suffixes := []int{}

	for i := range input {
		if len(prefixes) == 0 {
			prefixes = append(prefixes, input[i])
		} else {
			prefixes = append(prefixes, prefixes[len(prefixes)-1]*input[i])
		}
	}

	reverseInput := make([]int, len(input))
	copy(reverseInput, input)
	reverseInput = reverse(reverseInput)
	for i := range reverseInput {
		if len(suffixes) == 0 {
			suffixes = append(suffixes, reverseInput[i])
		} else {
			suffixes = append(suffixes, suffixes[len(suffixes)-1]*reverseInput[i])
		}
	}
	suffixes = reverse(suffixes)

	output := []int{}

	for i := range input {
		if i == 0 {
			output = append(output, suffixes[i+1])
		} else if i == len(input)-1 {
			output = append(output, prefixes[i-1])
		} else {
			output = append(output, prefixes[i-1]*suffixes[i+1])
		}
	}
	fmt.Println(output)
}
