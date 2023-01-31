package main

import (
	"testing"
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

func sliceIncludes(slice []TileType, needle TileType) bool {
	for _, tileType := range slice {
		if tileType == needle {
			return true
		}
	}

	return false
}

func WaveFunction(sampleInput [][]string) map[TileType]TileRulesList {
	tileRulesMap := make(map[TileType]TileRulesList)
	defaultTileRule := TileRulesList{Up: []TileType{}, Down: []TileType{}, Left: []TileType{}, Right: []TileType{}}
	tileRulesMap[Land] = defaultTileRule
	tileRulesMap[Sea] = defaultTileRule
	tileRulesMap[Coast] = defaultTileRule

	for i, row := range sampleInput {
		for j, tile := range row {
			tileName := calculateTileName(tile)
			newRule := newTileRules(sampleInput, tile, i, j)

			if entry, ok := tileRulesMap[tileName]; ok {
				if newRule.Up != None && !sliceIncludes(entry.Up, newRule.Up) {
					entry.Up = append(entry.Down, newRule.Up)
				}
				if newRule.Down != None && !sliceIncludes(entry.Down, newRule.Down) {
					entry.Down = append(entry.Down, newRule.Down)
				}
				if newRule.Left != None && !sliceIncludes(entry.Left, newRule.Left) {
					entry.Left = append(entry.Left, newRule.Left)
				}
				if newRule.Right != None && !sliceIncludes(entry.Right, newRule.Right) {
					entry.Right = append(entry.Right, newRule.Right)
				}
				tileRulesMap[tileName] = entry
			}
		}
	}
	return tileRulesMap
}

func TestGenerateRulesFromSampleInput(t *testing.T) {
	sampleInput := [][]string{
		{"L", "C", "S"},
		{"L", "C", "S"},
		{"C", "S", "L"},
	}
	tileRulesMap := WaveFunction(sampleInput)

	landRules := tileRulesMap[Land]
	if !sliceIncludes(landRules.Down, Land) || !sliceIncludes(landRules.Down, Coast) {
		t.Errorf("Land rules for Down isn't Land and Coast. Got %+v.", landRules.Down)
	}
	if !sliceIncludes(landRules.Up, Land) || !sliceIncludes(landRules.Up, Sea) {
		t.Errorf("Land rules for up isn't Land and Sea. Got %+v.", landRules.Down)
	}
	if !sliceIncludes(landRules.Right, Coast) {
		t.Errorf("Land rules for Right isn't Coast. Got %+v.", landRules.Down)
	}
	if !sliceIncludes(landRules.Left, Sea) {
		t.Errorf("Land rules for Left isn't Sea. Got %+v.", landRules.Down)
	}

	seaRules := tileRulesMap[Sea]
	if !sliceIncludes(seaRules.Down, Sea) {
		t.Errorf("Sea rules for Down isn't Sea. Got %+v.", seaRules.Down)
	}
	if !sliceIncludes(seaRules.Up, Coast) || !sliceIncludes(seaRules.Up, Sea) {
		t.Errorf("Sea rules for up isn't Coast and Sea. Got %+v.", seaRules.Up)
	}
	if !sliceIncludes(seaRules.Right, Land) {
		t.Errorf("Sea rules for Right isn't Land. Got %+v.", seaRules.Right)
	}
	if !sliceIncludes(seaRules.Left, Coast) {
		t.Errorf("Sea rules for Left isn't Coast. Got %+v.", seaRules.Left)
	}

	coastRules := tileRulesMap[Coast]
	if !sliceIncludes(coastRules.Down, Coast) || !sliceIncludes(coastRules.Down, Sea) {
		t.Errorf("Coast rules for Down isn't Coast and Sea. Got %+v.", coastRules.Down)
	}
	if !sliceIncludes(coastRules.Up, Coast) || !sliceIncludes(coastRules.Up, Land) {
		t.Errorf("Coast rules for up isn't Coast and Land. Got %+v.", coastRules.Up)
	}
	if !sliceIncludes(coastRules.Right, Sea) {
		t.Errorf("Coast rules for Right isn't Sea. Got %+v.", coastRules.Right)
	}
	if !sliceIncludes(coastRules.Left, Land) {
		t.Errorf("Coast rules for Left isn't Land. Got %+v.", coastRules.Left)
	}
}

type NumberGenerator func() int
type Entropy func(numberGenerator NumberGenerator, grid [3][3]Square, totalCollapsed int) int
type Square struct {
	Possibilities []TileType
	Type          *TileType
}

func collapse(ruleSet []TileRules, numberGenerator NumberGenerator, entropy Entropy) [3][3]Square {
	tileTypes := []TileType{Land, Sea, Coast}
	grid := [3][3]Square{}

	// fill all squares with possibilities
	for i, row := range grid {
		for j := 0; j < len(row); j++ {
			grid[i][j] = Square{Possibilities: tileTypes}
		}
	}

	totalCollapsed := 0
	gridSize := len(grid) * len(grid[0])
	for totalCollapsed < 2 {
		tileNumber := entropy(numberGenerator, grid, totalCollapsed)
		row, col := 0, 0
		if tileNumber != 0 {
			row = tileNumber / gridSize
			col = (tileNumber % gridSize) - 1
		}

		// decided the tile type
		grid[row][col].Type = &grid[row][col].Possibilities[numberGenerator()]
		totalCollapsed++

		// tile to the left
		grid[row][col-1].Possibilities = []TileType{ruleSet[0].Left}
		// tile to the right
		if col != len(grid[0])-1 {
			grid[row][col+1].Possibilities = []TileType{ruleSet[0].Right}
		}
		// tile above
		if row != 0 {
			grid[row-1][col].Possibilities = []TileType{ruleSet[0].Up}
		}
		// tile below
		if row != len(grid)-1 {
			grid[row+1][col].Possibilities = []TileType{ruleSet[0].Down}
		}
	}

	return grid
}

// Given a set of rules and an entropy function
// When a square is collapsed
// Collapses the next square given it has the lowest entropy
func TestWaveFunctionCollapseNextLowestEntropy(t *testing.T) {
	ruleOne := TileRules{Type: Land, Up: Land, Down: Sea, Right: Coast, Left: Sea}
	ruleSet := []TileRules{ruleOne}

	// initial square, pick one at random because they are all the same entropy
	numberGenerator := func() int {
		return 0
	}

	entropy := func(ng NumberGenerator, grid [3][3]Square, totalCollapsed int) int {
		if totalCollapsed == 0 {
			return 3
		}

		return 2
	}

	result := collapse(ruleSet, numberGenerator, entropy)
	if *result[0][2].Type != Land {
		t.Errorf("Square did not collapse into LAND but %s instead", *result[0][2].Type)
	}

	// the next square that should be picked based on the lowest entropy
	nextTileType := result[0][1]
	if *nextTileType.Type != Sea {
		t.Errorf("Next tile should be SEA but got %s", *nextTileType.Type)
	}
}

func TestRemovesPossibilitiesGivenASetOfRules(t *testing.T) {
	ruleOne := TileRules{Type: Land, Up: Land, Down: Sea, Right: Coast, Left: Sea}
	ruleTwo := TileRules{Type: Sea, Up: Sea, Down: Sea, Right: Coast, Left: Coast}
	ruleSet := []TileRules{ruleOne, ruleTwo}

	// initial square, pick one at random because they are all the same entropy
	numberGenerator := func() int {
		return 0
	}

	entropy := func(ng NumberGenerator, grid [3][3]Square, totalCollapsed int) int {
		if totalCollapsed == 0 {
			return 3
		}

		return 2
	}

	result := collapse(ruleSet, numberGenerator, entropy)
	if *result[0][2].Type != Land {
		t.Errorf("Square did not collapse into LAND but %s instead", *result[0][2].Type)
	}

	// the next square that should be picked based on the lowest entropy
	nextTileType := result[0][1]
	if *nextTileType.Type != Sea {
		t.Errorf("Next tile should be SEA but got %s", *nextTileType.Type)
	}
}
