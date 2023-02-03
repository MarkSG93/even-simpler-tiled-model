package main

import (
	"math"
	"math/rand"
	"sort"

	"golang.org/x/exp/slices"
)

type TileType = string

const (
	Land  TileType = "LAND"
	Sea   TileType = "SEA"
	Coast TileType = "COAST"
	None           = ""
)

type TileRules struct {
	Type  TileType
	Left  TileType
	Right TileType
	Up    TileType
	Down  TileType
}

type TileRulesList struct {
	Left  []TileType
	Right []TileType
	Up    []TileType
	Down  []TileType
}

type NumberGenerator func(n int) int
type Entropy func(grid [][]Square, totalCollapsed int) int
type Square struct {
	Possibilities []TileType
	Type          *TileType
}
type RuleSet = map[TileType]TileRulesList

func WaveFunction(sampleInput [][]string, gridArea int) [][]Square {
	entropy := func(grid [][]Square, totalCollapsed int) int {
		return totalCollapsed
	}

	numberGenerator := func(n int) int {
		if n == 0 {
			return n
		}
		return rand.Intn(n)
	}

	ruleSet := generateRuleSet(sampleInput)
	return collapse(ruleSet, numberGenerator, entropy, gridArea)
}

func collapse(ruleSet RuleSet, numberGenerator NumberGenerator, entropy Entropy, gridArea int) [][]Square {
	tileTypes := []TileType{Coast, Land, Sea}

	// fill all squares with possibilities
	gridWidth := int(math.Sqrt(float64(gridArea)))
	grid := newGrid(gridWidth, tileTypes)

	totalCollapsed := 0
	for totalCollapsed < gridArea {
		tileNumber := entropy(grid, totalCollapsed)
		row, col := calculateRowAndColumn(tileNumber, gridWidth)

		// Is there a contradiction?
		if len(grid[row][col].Possibilities) < 1 {
			return collapse(ruleSet, numberGenerator, entropy, gridArea)
		}

		// decided the tile type
		collapsedTileType := &grid[row][col].Possibilities[numberGenerator(len(grid[row][col].Possibilities))]
		grid[row][col].Type = collapsedTileType
		totalCollapsed++

		collapsedTileRuleSet := ruleSet[*collapsedTileType]
		grid = updateGridPossibilities(grid, row, col, collapsedTileRuleSet)
	}

	return grid
}

func newGrid(gridWidth int, tileTypes []TileType) [][]Square {
	grid := [][]Square{}
	for i := 0; i < gridWidth; i++ {
		row := make([]Square, gridWidth)
		for x := 0; x < gridWidth; x++ {
			row[x] = Square{Possibilities: tileTypes}
		}
		grid = append(grid, row)
	}

	return grid
}

func calculateRowAndColumn(index int, gridWidth int) (row, col int) {
	row, col = 0, 0
	if index != 0 {
		row = int(math.Floor(float64(index) / float64(gridWidth)))
		col = index % gridWidth
	}

	return row, col
}

func updateGridPossibilities(grid [][]Square, row int, col int, collapsedTileRuleSet TileRulesList) [][]Square {
	// left most tile
	if col-1 >= 0 {
		grid[row][col-1].Possibilities = getMatchingItems(collapsedTileRuleSet.Left, grid[row][col-1].Possibilities)
	}
	// tile to the right
	if col != len(grid[0])-1 {
		grid[row][col+1].Possibilities = getMatchingItems(collapsedTileRuleSet.Right, grid[row][col+1].Possibilities)
	}
	// tile above
	if row != 0 {
		grid[row-1][col].Possibilities = getMatchingItems(collapsedTileRuleSet.Up, grid[row-1][col].Possibilities)
	}
	// tile below
	if row != len(grid)-1 {
		grid[row+1][col].Possibilities = getMatchingItems(collapsedTileRuleSet.Down, grid[row+1][col].Possibilities)
	}

	return grid
}

func shouldAddRule(entries []TileType, newTile TileType) bool {
	return newTile != None && slices.Index(entries, newTile) == -1
}

func generateRuleSet(sampleInput [][]string) RuleSet {
	tileRulesMap := make(RuleSet)
	defaultTileRule := TileRulesList{Up: []TileType{}, Down: []TileType{}, Left: []TileType{}, Right: []TileType{}}
	tileRulesMap[Land] = defaultTileRule
	tileRulesMap[Sea] = defaultTileRule
	tileRulesMap[Coast] = defaultTileRule

	for i, row := range sampleInput {
		for j, tile := range row {
			tileName := calculateTileName(tile)
			newRule := newTileRules(sampleInput, tile, i, j)

			entry, ok := tileRulesMap[tileName]
			if !ok {
				continue
			}

			if shouldAddRule(entry.Up, newRule.Up) {
				entry.Up = append(entry.Down, newRule.Up)
			}
			if shouldAddRule(entry.Down, newRule.Down) {
				entry.Down = append(entry.Down, newRule.Down)
			}
			if shouldAddRule(entry.Left, newRule.Left) {
				entry.Left = append(entry.Left, newRule.Left)
			}
			if shouldAddRule(entry.Right, newRule.Right) {
				entry.Right = append(entry.Right, newRule.Right)
			}
			tileRulesMap[tileName] = entry
		}
	}
	return tileRulesMap
}

func calculateTileName(sample string) TileType {
	switch sample {
	case "C":
		return Coast
	case "L":
		return Land
	case "S":
		return Sea
	}

	return None
}

func newTileRules(sampleInput [][]string, tile string, i int, j int) TileRules {
	row := sampleInput[i]
	tileRules := TileRules{Type: calculateTileName(tile)}
	if i+1 < len(sampleInput) {
		tileRules.Down = calculateTileName(sampleInput[i+1][j])
	}

	if i > 0 {
		tileRules.Up = calculateTileName(sampleInput[i-1][j])
	}

	if j == 0 {
		tileRules.Right = calculateTileName(row[j+1])
		tileRules.Left = ""
		return tileRules
	}

	if j+1 >= len(row) { // on the right most tile
		tileRules.Right = ""
		tileRules.Left = calculateTileName(row[j-1])
	} else { // middle tiles
		tileRules.Left = calculateTileName(row[j-1])
		tileRules.Right = calculateTileName(row[j+1])
	}

	return tileRules
}

func getMatchingItems(a []TileType, b []TileType) []TileType {
	hits := map[string]int{
		Land:  0,
		Sea:   0,
		Coast: 0,
	}

	for tileType := range hits {
		if slices.Index(a, tileType) != -1 {
			hits[tileType] += 1
		}

		if slices.Index(b, tileType) != -1 {
			hits[tileType] += 1
		}
	}

	matchingItems := []TileType{}
	for tileType, hits := range hits {
		if hits > 1 {
			matchingItems = append(matchingItems, tileType)
		}
	}
	sort.Strings(matchingItems)
	return matchingItems
}
