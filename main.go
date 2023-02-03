package main

import (
	"fmt"

	"github.com/DrSmithFr/go-console/pkg/output"
)

func main() {
	sampleInput := [][]string{
		{"L", "C", "L", "L"},
		{"C", "S", "C", "L"},
		{"L", "L", "C", "S"},
		{"C", "S", "S", "C"},
	}

	out := output.NewConsoleOutput(true, nil)
	grid := WaveFunction(sampleInput, 144)
	totalsOfTileTypes := make(map[TileType]int)
	for _, row := range grid {
		for _, col := range row {
			totalsOfTileTypes[*col.Type]++
			if *col.Type == Land {
				out.Write("<info>L  <info>")
			}
			if *col.Type == Coast {
				out.Write("<fg=yellow>C  </>")
			}
			if *col.Type == Sea {
				out.Write("<fg=blue>S  </>")
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("%+v", totalsOfTileTypes)
}
