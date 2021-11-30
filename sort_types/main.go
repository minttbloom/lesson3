package main

import (
	"fmt"
	"sort"
)

func bubbleSort(a [10]int) [10]int {
	length := len(a)
	for i := 0; i < length-1; i++ {
		for l := i + 1; l < length; l++ {
			if a[i] > a[l] {
				a[i], a[l] = a[l], a[i]
			}
		}
	}

	return a
}

func simpleSort(b []int) []int {
	sort.Ints(b)
	return b
}

func insertSort(x []int) []int {
	var n = len(x)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if x[j-1] > x[j] {
				x[j-1], x[j] = x[j], x[j-1]
			}
			j = j - 1
		}
	}
	return x
}

func main() {
	m := bubbleSort([10]int{0, 10, 6, 56, 34, 5, 9, 8, 31, 10})
	fmt.Println(m)

	s := simpleSort([]int{0, 10, 6, 56, 34, 5, 9, 8, 31, 10})
	fmt.Println(s)

	d := insertSort([]int{0, 10, 6, 56, 34, 5, 9, 8, 31, 10})
	fmt.Println(d)

}