package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortingStringArray(t *testing.T) {
	inputArray := []string{
		"1234567890",
		"123456789098765",
		"1234",
		"12345678",
	}
	result := []string{"1234", "12345678", "1234567890", "123456789098765"}
	notEqual := []string{"123456789098765", "1234567890", "12345678", "1234"}
	assert.Equal(t, result, sortingStringArray(inputArray))
	assert.NotEqual(t, notEqual, sortingStringArray(inputArray))
}
