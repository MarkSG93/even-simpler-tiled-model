package main

import "fmt"

func main() {
	sampleInput := [][]string{
		{"L", "C", "L"},
		{"C", "S", "C"},
		{"L", "C", "S"},
	}

	grid := WaveFunction(sampleInput, 16)
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%s ", *col.Type)
		}
		fmt.Printf("\n")
	}
}
