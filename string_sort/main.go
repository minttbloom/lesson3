package main

import "fmt"

func strMap2() map[string]int {
	inputArray := []string {
		"1234567890",
		"123456789098765",
		"1234",
		"12345678",
	}
	m := make(map[string]int, len(inputArray))
	for k, v := range inputArray {
		m[v] = k
	}
	return m
}

func main() {
	sortedArray := strMap2()
	fmt.Println(sortedArray)
}
