package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var a = [10]int{4, 6, 9, 8, 1, 5, 2, 3, 7, 10}
	var b = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var notEqual = [10]int{4, 6, 9, 8, 1, 5, 2, 3, 7, 10}
	var notEqual1 = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	assert.Equal(t, b, bubbleSort(a))
	assert.NotEqual(t, notEqual, bubbleSort(a))
	assert.NotEqual(t, notEqual1, bubbleSort(a))
}

func TestSimpleSort(t *testing.T) {
	var a = []int{4, 6, 9, 8, 1, 5, 2, 3, 7, 10}
	var b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var notEqual = [10]int{4, 6, 9, 8, 1, 5, 2, 3, 7, 10}
	var notEqual1 = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	assert.Equal(t, b, simpleSort(a))
	assert.NotEqual(t, notEqual, simpleSort(a))
	assert.NotEqual(t, notEqual1, simpleSort(a))
}

func TestInsertSort(t *testing.T) {
	var a = []int{4, 6, 9, 8, 1, 5, 2, 3, 7, 10}
	var b = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var notEqual = [10]int{4, 6, 9, 8, 1, 5, 2, 3, 7, 10}
	var notEqual1 = [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	assert.Equal(t, b, insertSort(a))
	assert.NotEqual(t, notEqual, insertSort(a))
	assert.NotEqual(t, notEqual1, insertSort(a))
}
