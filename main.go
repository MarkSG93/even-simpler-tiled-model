package main

import "fmt"

func main() {
	sampleInput := [][]string{
		{"L", "C", "L", "L"},
		{"C", "S", "C", "L"},
		{"L", "L", "C", "S"},
		{"C", "S", "S", "C"},
	}

	grid := WaveFunction(sampleInput, 144)
	totalsOfTileTypes := make(map[TileType]int)
	for _, row := range grid {
		for _, col := range row {
			totalsOfTileTypes[*col.Type]++
			fmt.Printf("%s ", *col.Type)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("%+v", totalsOfTileTypes)
}
