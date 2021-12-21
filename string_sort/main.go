package main

import (
	"fmt"
	"sort"
)

func sortingStringArray(inputArray []string) []string {
	sort.SliceStable(inputArray, func(i, j int) bool {
		return inputArray[i] < inputArray[j]
	})
	return inputArray
}

func main() {
	inputArray := []string{
		"1234567890",
		"123456789098765",
		"1234",
		"12345678",
	}
	sortedArray := sortingStringArray(inputArray)
	fmt.Println(sortedArray)
}
