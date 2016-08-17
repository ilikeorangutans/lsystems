package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkeleton(t *testing.T) {
	input := "Y"
	expected := "XYX"

	var productions Productions = make(map[rune]string)
	productions['Y'] = "XYX"

	result := productions.Apply(input)

	assert.Equal(t, expected, result)
}

func TestRepeated(t *testing.T) {
	input := "B"
	expected := "ABAAB"

	var productions Productions = make(map[rune]string)
	productions['A'] = "AB"
	productions['B'] = "A"

	result := productions.ApplyTimes(input, 4)

	assert.Equal(t, expected, result)
}

func TestLoop(t *testing.T) {
	input := "FX"

	var productions Productions = make(map[rune]string)
	productions['X'] = "X+YF+"
	productions['Y'] = "-FX-Y"
	productions['F'] = ""

	result := productions.ApplyTimes(input, 2)

	assert.Equal(t, "X+YF++-FX-Y+", result)
}
