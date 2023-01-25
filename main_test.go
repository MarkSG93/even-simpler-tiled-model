package main

import "testing"

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

func NewTileRules(sampleInput [][]string, tile string, i int, j int) TileRules {
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
			tileRulesList = append(tileRulesList, NewTileRules(sampleInput, tile, i, j))
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