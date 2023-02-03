package main

import "fmt"

func main() {
	sampleInput := [][]string{
		{"L", "C", "L"},
		{"C", "S", "C"},
		{"L", "C", "S"},
		{"L", "C", "S"},
	}

	grid := WaveFunction(sampleInput, 144)
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%s ", *col.Type)
		}
		fmt.Printf("\n")
	}
}
