package main

import (
	"fmt"
	sort "sort"
	"strings"
)

func letterCheck(a, b string) bool {
	a = strings.Join(strings.Fields(a), "")
	b = strings.Join(strings.Fields(b), "")
	if len(a) != len(b) {
		return false
	}
	ab := []byte(a)
	bb := []byte(b)

	sort.Slice(ab, func(i, j int) bool {
		return ab[i] < ab[j]
	})
	sort.Slice(bb, func(i, j int) bool {
		return bb[i] < bb[j]
	})

	for i, _ := range ab {
		if ab[i] != bb[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(letterCheck("b a a", "ab  b"))
}
