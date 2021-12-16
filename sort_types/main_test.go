package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var a = [10]int{4, 6, 9, 8, 1, 5, 2, 3, 7, 10}
	var b = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, b, bubbleSort(a))
}

func TestSimpleSort(t *testing.T) {
	var a = []int{4, 6, 9, 8, 1, 5, 2, 3, 7, 10}
	var b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, b, simpleSort(a))
}

func TestInsertSort(t *testing.T) {
	var a = []int{4, 6, 9, 8, 1, 5, 2, 3, 7, 10}
	var b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, b, insertSort(a))
}
