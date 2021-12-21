package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLenght(t *testing.T) {
	assert.Equal(t, true, letterCheck("a", "a"))
	assert.NotEqual(t, true, letterCheck("a", "b"))
	assert.NotEqual(t, true, letterCheck("a", " "))
}
