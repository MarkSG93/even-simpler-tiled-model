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

func TestRotateSlice(t *testing.T) {
	input := [][]string{
		{"L", "L", "L"},
		{"C", "C", "C"},
		{"S", "S", "S"},
	}
	expected := [][]string{
		{"S", "C", "L"},
		{"S", "C", "L"},
		{"S", "C", "L"},
	}

	result := rotate(input)
	if !slices.Equal(result[0], expected[0]) {
		t.Errorf("Expected %+v didn't match result %+v", expected[0], result[0])
	}
	if !slices.Equal(result[1], expected[1]) {
		t.Errorf("Expected %+v didn't match result %+v", expected[0], result[0])
	}
	if !slices.Equal(result[2], expected[2]) {
		t.Errorf("Expected %+v didn't match result %+v", expected[0], result[0])
	}
}

func rotate(matrix [][]string) [][]string {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[m-j-1][i]
			matrix[m-j-1][i] = matrix[m-i-1][n-j-1]
			matrix[m-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = temp
		}
	}

	return matrix
}

func TestApplySinglePattern(t *testing.T) {
	input := [][]string{
		{"L", "L", "L"},
		{"L", "L", "L"},
		{"L", "L", "L"},
	}
	expected := [][]string{
		{"L", "L", "L"},
		{"L", "L", "L"},
		{"L", "L", "L"},
	}

	ng := func(n int) int {
		return 0
	}
	rotator := func(matrix [][]string) [][]string {
		return matrix
	}
	result := overlappingCollapse(input, ng, rotator)
	if !slices.Equal(result[0], expected[0]) {
		t.Errorf("Expected %+v didn't match result %+v", expected[0], result[0])
	}
	if !slices.Equal(result[1], expected[1]) {
		t.Errorf("Expected %+v didn't match result %+v", expected[0], result[0])
	}
	if !slices.Equal(result[2], expected[2]) {
		t.Errorf("Expected %+v didn't match result %+v", expected[0], result[0])
	}
}

func TestApplyRandomPatterns(t *testing.T) {
	input := [][]string{
		{"L", "L", "L"},
		{"L", "L", "L"},
		{"L", "L", "L"},
		{"C", "C", "C"},
		{"C", "C", "C"},
		{"C", "C", "C"},
	}
	expected := [][]string{
		{"C", "C", "C"},
		{"C", "C", "C"},
		{"C", "C", "C"},
	}
	ng := func(n int) int {
		return 1
	}
	rotator := func(matrix [][]string) [][]string {
		return matrix
	}
	result := overlappingCollapse(input, ng, rotator)
	if !slices.Equal(result[0], expected[0]) {
		t.Errorf("Expected %+v didn't match result %+v", expected[0], result[0])
	}
	if !slices.Equal(result[1], expected[1]) {
		t.Errorf("Expected %+v didn't match result %+v", expected[0], result[0])
	}
	if !slices.Equal(result[2], expected[2]) {
		t.Errorf("Expected %+v didn't match result %+v", expected[0], result[0])
	}
}

func TestApplyAndRotatePattern(t *testing.T) {
	input := [][]string{
		{"L", "L", "L"},
		{"C", "C", "C"},
		{"S", "S", "S"},
	}
	expected := [][]string{
		{"S", "C", "L"},
		{"S", "C", "L"},
		{"S", "C", "L"},
	}
	ng := func(n int) int {
		return 0
	}

	result := overlappingCollapse(input, ng, rotate)
	if !slices.Equal(result[0], expected[0]) {
		t.Errorf("Expected %+v didn't match result %+v", expected[0], result[0])
	}
	if !slices.Equal(result[1], expected[1]) {
		t.Errorf("Expected %+v didn't match result %+v", expected[0], result[0])
	}
	if !slices.Equal(result[2], expected[2]) {
		t.Errorf("Expected %+v didn't match result %+v", expected[0], result[0])
	}
}

type Rotator = func(matrix [][]string) [][]string

func overlappingCollapse(sampleInput [][]string, ng NumberGenerator, rotator Rotator) [][]string {
	patterns := generatePatterns(sampleInput)

	output := [][]string{}
	selectedPattern := patterns[ng(len(patterns))]
	rotatedPattern := rotator(selectedPattern)
	for _, row := range rotatedPattern {
		output = append(output, row)
	}

	return output
}
