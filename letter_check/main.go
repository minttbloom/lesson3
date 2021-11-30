package main

import "fmt"

func letterCheck(a, b string) bool {
	for _, v := range a {
		charIsFind := false
		for _, v2 := range b {
			if v == v2 {
				charIsFind = true
			}
		}
		if !charIsFind {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(letterCheck("asdf", "asfd"))
}
