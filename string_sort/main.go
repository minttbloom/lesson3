package main

import (
	"fmt"
	"sort"
)

func strMap2() map[string]int {
	inputArray := []string{
		"1234567890",
		"123456789098765",
		"1234",
		"12345678",
	}
	mapInputArray := make(map[string]int, len(inputArray))
	for k, v := range inputArray {
		mapInputArray[v] = k
	}

	// sort
	sort.SliceStable(mapInputArray, func(i, j int) bool {
		first := mapInputArray[i]
		return mapInputArray < mapInputArray[j]
	})

	return mapInputArray
}

func main() {
	sortedArray := strMap2()
	fmt.Println(sortedArray)
}
