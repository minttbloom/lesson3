package main

import (
	"fmt"
	"strings"
)

func letterCheck(a, b string) bool {
	a = strings.TrimSpace(" ")
	b = strings.TrimSpace(" ")
	if len(a) != len(b) {
		return false
	}
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
	fmt.Println(letterCheck("sadf   tr", "a	sfd rt"))
}
