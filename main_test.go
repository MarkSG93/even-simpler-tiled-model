package main

import "testing"

type Rule struct {
	TileOne     string
	TileTwo     string
	Orientation string
}

type TileType = string

const (
	Land  TileType = "LAND"
	Sea   TileType = "SEA"
	Coast TileType = "COAST"
)

type Tile struct {
	Type  TileType
	Left  Rule
	Right Rule
	Up    Rule
	Down  Rule
}

type Rules = []Rule

func calculateTileName(sample string) string {
	switch sample {
	case "C":
		return "COAST"
	case "L":
		return "LAND"
	case "S":
		return "SEA"
	}

	return ""
}

func generateLeftMostTile(currentTile string, nextTile string) Rule {
	return Rule{TileOne: calculateTileName(currentTile), TileTwo: calculateTileName(nextTile), Orientation: "LEFT"}
}

func waveFunction(sampleInput [][]string) Rules {
	rules := Rules{}
	for _, row := range sampleInput {
		for j, tile := range row {
			if j == 0 {
				rules = append(rules, generateLeftMostTile(tile, row[j+1]))
				continue
			}

			rule := Rule{}
			if j+1 >= len(row) { // on the right most tile
				rule = Rule{TileOne: calculateTileName(tile), TileTwo: calculateTileName(row[j-1]), Orientation: "LEFT"}
			} else {
				rule = Rule{TileOne: calculateTileName(tile), TileTwo: calculateTileName(row[j+1]), Orientation: "RIGHT"}
				rules = append(rules, Rule{TileOne: calculateTileName(tile), TileTwo: calculateTileName(row[j-1]), Orientation: "LEFT"})
			}
			rules = append(rules, rule)
		}
	}
	return rules
}

func TestGenerateRulesFromSampleInput(t *testing.T) {
	sampleInput := [][]string{
		{"L", "C", "S"},
	}
	rules := waveFunction(sampleInput)

	if rules[0].TileOne != "LAND" && rules[0].TileTwo != "COAST" && rules[0].Orientation != "LEFT" {
		t.Errorf("Rule 1 is fucked")
	}

	if rules[1].TileOne != "COAST" && rules[1].TileTwo != "LAND" && rules[1].Orientation != "RIGHT" {
		t.Errorf("Rule 2 is fucked")
	}

	if rules[2].TileOne != "COAST" && rules[2].TileTwo != "SEA" && rules[2].Orientation != "LEFT" {
		t.Errorf("Rule 3 is fucked")
	}

	if rules[3].TileOne != "SEA" && rules[3].TileTwo != "COAST" && rules[3].Orientation != "RIGHT" {
		t.Errorf("Rule 4 is fucked")
	}
}

func TestGenerateRulesFromMultipleRowSampleInput(t *testing.T) {
	sampleInput := [][]string{
		{"L", "C"},
		{"C", "S"},
	}
	rules := waveFunction(sampleInput)
	if rules[2].TileOne != "COAST" && rules[2].TileTwo != "SEA" && rules[2].Orientation != "LEFT" {
		t.Errorf("The rule for the 1st item in the 2nd row is totally fucked")
	}

	if rules[3].TileOne != "SEA" && rules[3].TileTwo != "COAST" && rules[3].Orientation != "RIGHT" {
		t.Errorf("The rule for the 2nd item in the 2nd row is totally fucked")
	}

}

func TestConvertSampleToTileValues(t *testing.T) {
	sampleInput := [][]string{
		{"C", "S", "L"},
	}
	rules := waveFunction(sampleInput)
	if rules[0].TileOne != "COAST" {
		t.Errorf("Tile one did not match COAST")
	}
	if rules[0].TileTwo != "SEA" {
		t.Errorf("Tile two did not match SEA")
	}
	if rules[1].TileTwo != "LAND" {
		t.Errorf("Tile 3 did not match LAND but %s", rules[2].TileOne)
	}
	if rules[2].TileOne != "LAND" {
		t.Errorf("Rule 3 Tile 1 did not match LAND but %s", rules[2].TileOne)
	}
}
