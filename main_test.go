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

func collapse(ruleSet []TileRules, numberGenerator NumberGenerator) [1][2]string {
	tileTypes := []TileType{Land, Sea, Coast}
	grid := [1][2]string{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			randomNumber := numberGenerator()
			tileType := tileTypes[randomNumber]

			rule := ruleSet[0]
			for {
				if j == 0 {
					break
				}
				previousTile := grid[i][j-1]
				if rule.Type == previousTile && tileType == rule.Right {
					break
				}
				randomNumber++
				tileType = tileTypes[randomNumber]
			}
			grid[i][j] = tileType
		}
	}

	return grid
}

// Given a set of rules
// Collapses a square into a tile following those rules
func TestWaveFunctionCollapseSingleSquare(t *testing.T) {
	ruleOne := TileRules{Type: Land, Up: Land, Down: Sea, Right: Coast, Left: Sea}
	ruleSet := []TileRules{ruleOne}

	numberGenerator := func() int {
		return 0
	}

	result := collapse(ruleSet, numberGenerator)
	if result[0][0] != "LAND" {
		t.Errorf("Square did not collapse into LAND but %s instead", result[0][0])
	}
	/*
		[
			[L, C, S],
			[S, C, L]
		]
	*/
}

func TestWaveFunctionCollapseMultipleSquares(t *testing.T) {
	ruleOne := TileRules{Type: Land, Up: Land, Down: Sea, Right: Coast, Left: Sea}
	ruleSet := []TileRules{ruleOne}

	numberGenerator := func() int {
		return 0
	}

	result := collapse(ruleSet, numberGenerator)
	if result[0][0] != Land {
		t.Errorf("Square did not collapse into LAND but %s instead", result[0][0])
	}

	nextTileType := result[0][1]
	if nextTileType != Coast {
		t.Errorf("Next tile should be COAST but got %s", nextTileType)
	}
	/*
		[
			[L, C, S],
			[S, C, L]
		]
	*/
}
