package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

// Given a sample input
// it generates the patterns in the input
func TestGeneratePatterns(t *testing.T) {
	sampleInput := [][]string{
		{"L", "L", "L"},
		{"C", "C", "C"},
		{"S", "S", "S"},
		{"C", "S", "C"},
		{"L", "L", "S"},
		{"C", "C", "C"},
		{"S", "S", "S"},
		{"S", "S", "S"},
		{"S", "S", "S"},
	}
	expected := [][][]string{
		{{"L", "L", "L"},
			{"C", "C", "C"},
			{"S", "S", "S"}},
		{{"C", "S", "C"},
			{"L", "L", "S"},
			{"C", "C", "C"}},
		{{"S", "S", "S"},
			{"S", "S", "S"},
			{"S", "S", "S"}},
	}
	pattern := generatePatterns(sampleInput)
	if !slices.Equal(expected[0][0], pattern[0][0]) {
		t.Errorf("Pattern %+v didn't match expected %+v", expected, pattern)
	}
	if !slices.Equal(expected[0][1], pattern[0][1]) {
		t.Errorf("Pattern %+v didn't match expected %+v", expected, pattern)
	}
	if !slices.Equal(expected[0][2], pattern[0][2]) {
		t.Errorf("Pattern %+v didn't match expected %+v", expected, pattern)
	}

	if !slices.Equal(expected[1][0], pattern[1][0]) {
		t.Errorf("Pattern %+v didn't match expected %+v", expected, pattern)
	}
	if !slices.Equal(expected[1][1], pattern[1][1]) {
		t.Errorf("Pattern %+v didn't match expected %+v", expected, pattern)
	}
	if !slices.Equal(expected[1][2], pattern[1][2]) {
		t.Errorf("Pattern %+v didn't match expected %+v", expected, pattern)
	}

	if !slices.Equal(expected[2][0], pattern[2][0]) {
		t.Errorf("Pattern %+v didn't match expected %+v", expected, pattern)
	}
	if !slices.Equal(expected[2][1], pattern[2][1]) {
		t.Errorf("Pattern %+v didn't match expected %+v", expected, pattern)
	}
	if !slices.Equal(expected[2][2], pattern[2][2]) {
		t.Errorf("Pattern %+v didn't match expected %+v", expected, pattern)
	}
}

func generatePatterns(sampleInput [][]string) [][][]string {
	width := len(sampleInput)
	height := len(sampleInput[0])

	patterns := [][][]string{}
	for i := 0; i < width/height; i++ {
		var pattern [][]string
		if i == 0 {
			pattern = sampleInput[0:height][0:height]
		}
		pattern = sampleInput[i*height : i*height+height][0:height]
		patterns = append(patterns, pattern)
	}

	return patterns
}
