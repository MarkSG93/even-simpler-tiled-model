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

func WaveFunction(sampleInput [][]string) []TileRules {
	tileRulesList := []TileRules{}
	for i, row := range sampleInput {
		for j, tile := range row {
			tileRulesList = append(tileRulesList, newTileRules(sampleInput, tile, i, j))
		}
	}
	return tileRulesList
}

func TestGenerateRulesFromSampleInput(t *testing.T) {
	sampleInput := [][]string{
		{"L", "C", "S"},
		{"L", "C", "S"},
		{"C", "S", "L"},
	}
	tileRulesList := WaveFunction(sampleInput)

	tileRuleOne := tileRulesList[0]
	if tileRuleOne.Type != "LAND" || tileRuleOne.Down != "LAND" || tileRuleOne.Up != "" || tileRuleOne.Right != "COAST" || tileRuleOne.Left != "" {
		t.Errorf("Tile rule 1 invalid %+v", tileRuleOne)
	}

	tileRuleTwo := tileRulesList[1]
	if tileRuleTwo.Type != "COAST" || tileRuleTwo.Down != "COAST" || tileRuleTwo.Up != "" || tileRuleTwo.Right != "SEA" || tileRuleTwo.Left != "LAND" {
		t.Errorf("Tile rule 2 invalid %+v", tileRuleTwo)
	}

	tileRuleThree := tileRulesList[2]
	if tileRuleThree.Type != "SEA" || tileRuleThree.Down != "SEA" || tileRuleThree.Up != "" || tileRuleThree.Right != "" || tileRuleThree.Left != "COAST" {
		t.Errorf("Tile rule 3 invalid %+v", tileRuleThree)
	}

	tileRuleFour := tileRulesList[3]
	if tileRuleFour.Type != "LAND" || tileRuleFour.Down != "COAST" || tileRuleFour.Up != "LAND" || tileRuleFour.Right != "COAST" || tileRuleFour.Left != "" {
		t.Errorf("Tile rule 4 invalid %+v", tileRuleFour)
	}

	tileRuleFive := tileRulesList[4]
	if tileRuleFive.Type != "COAST" || tileRuleFive.Down != "SEA" || tileRuleFive.Up != "COAST" || tileRuleFive.Right != "SEA" || tileRuleFive.Left != "LAND" {
		t.Errorf("Tile rule 5 invalid %+v", tileRuleFive)
	}

	tileRuleSix := tileRulesList[5]
	if tileRuleSix.Type != "SEA" || tileRuleSix.Down != "LAND" || tileRuleSix.Up != "SEA" || tileRuleSix.Right != "" || tileRuleSix.Left != "COAST" {
		t.Errorf("Tile rule 6 invalid %+v", tileRuleSix)
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
