package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSomething ...
func TestEqual(t *testing.T) {
	var a string = "Hello"
	var b string = "Hello"

	assert.Equal(t, a, b, "The two words should be the same.")
}

// TestSomething ...
func TestNotEqual(t *testing.T) {
	var a string = "Hello1"
	var b string = "Hello"

	assert.NotEqual(t, a, b, "The two words should not be the same.")
}
