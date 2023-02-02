package main

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestGenerateRulesFromSampleInput(t *testing.T) {
	sampleInput := [][]string{
		{"L", "C", "S"},
		{"L", "C", "S"},
		{"C", "S", "L"},
	}
	tileRulesMap := generateRuleSet(sampleInput)

	landRules := tileRulesMap[Land]
	if slices.Index(landRules.Down, Land) == -1 || slices.Index(landRules.Down, Coast) == -1 {
		t.Errorf("Land rules for Down isn't Land and Coast. Got %+v.", landRules.Down)
	}
	if slices.Index(landRules.Up, Land) == -1 || slices.Index(landRules.Up, Sea) == -1 {
		t.Errorf("Land rules for up isn't Land and Sea. Got %+v.", landRules.Down)
	}
	if slices.Index(landRules.Right, Coast) == -1 {
		t.Errorf("Land rules for Right isn't Coast. Got %+v.", landRules.Down)
	}
	if slices.Index(landRules.Left, Sea) == -1 {
		t.Errorf("Land rules for Left isn't Sea. Got %+v.", landRules.Down)
	}

	seaRules := tileRulesMap[Sea]
	if slices.Index(seaRules.Down, Sea) == -1 {
		t.Errorf("Sea rules for Down isn't Sea. Got %+v.", seaRules.Down)
	}
	if slices.Index(seaRules.Up, Coast) == -1 || slices.Index(seaRules.Up, Sea) == -1 {
		t.Errorf("Sea rules for up isn't Coast and Sea. Got %+v.", seaRules.Up)
	}
	if slices.Index(seaRules.Right, Land) == -1 {
		t.Errorf("Sea rules for Right isn't Land. Got %+v.", seaRules.Right)
	}
	if slices.Index(seaRules.Left, Coast) == -1 {
		t.Errorf("Sea rules for Left isn't Coast. Got %+v.", seaRules.Left)
	}

	coastRules := tileRulesMap[Coast]
	if slices.Index(coastRules.Down, Coast) == -1 || slices.Index(coastRules.Down, Sea) == -1 {
		t.Errorf("Coast rules for Down isn't Coast and Sea. Got %+v.", coastRules.Down)
	}
	if slices.Index(coastRules.Up, Coast) == -1 || slices.Index(coastRules.Up, Land) == -1 {
		t.Errorf("Coast rules for up isn't Coast and Land. Got %+v.", coastRules.Up)
	}
	if slices.Index(coastRules.Right, Sea) == -1 {
		t.Errorf("Coast rules for Right isn't Sea. Got %+v.", coastRules.Right)
	}
	if slices.Index(coastRules.Left, Land) == -1 {
		t.Errorf("Coast rules for Left isn't Land. Got %+v.", coastRules.Left)
	}
}

func TestCollapsesAllSquaresInAGrid(t *testing.T) {
	numberGenerator := func() int {
		return 0
	}

	entropy := func(grid [][]Square, totalCollapsed int) int {
		return totalCollapsed
	}

	sampleInput := [][]string{
		{"L", "C", "S"},
		{"C", "L", "S"},
		{"S", "C", "L"},
	}
	ruleSet := generateRuleSet(sampleInput)
	grid := collapse(ruleSet, numberGenerator, entropy, 9)
	firstTile := grid[0][0]
	secondTile := grid[0][1]
	thirdTile := grid[0][2]
	fourthTile := grid[1][0]
	fifthTile := grid[1][1]
	sixthTile := grid[1][2]
	seventhTile := grid[2][0]
	eigthTile := grid[2][1]
	ninthTile := grid[2][2]
	if *firstTile.Type != Coast {
		t.Errorf("First tile fucked: %+v", *firstTile.Type)
	}
	if *secondTile.Type != Land {
		t.Errorf("Second tile fucked: %+v", *secondTile.Type)
	}
	if *thirdTile.Type != Coast {
		t.Errorf("Third tile fucked: %+v", *thirdTile.Type)
	}
	if *fourthTile.Type != Land {
		t.Errorf("Fourth tile fucked: %+v", *fourthTile.Type)
	}
	if *fifthTile.Type != Coast {
		t.Errorf("Fifth tile fucked: %+v", *fifthTile.Type)
	}
	if *sixthTile.Type != Land {
		t.Errorf("Sixth tile fucked: %+v", *sixthTile.Type)
	}
	if *seventhTile.Type != Coast {
		t.Errorf("Seventh tile fucked: %+v", *seventhTile.Type)
	}
	if *eigthTile.Type != Land {
		t.Errorf("Eight tile fucked: %+v", *eigthTile.Type)
	}
	if *ninthTile.Type != Coast {
		t.Errorf("Ninth tile fucked: %+v", *ninthTile.Type)
	}
}

func TestGetMatchingItemsInSlices(t *testing.T) {
	a := []TileType{Land, Sea, Coast}
	b := []TileType{Land, Coast}
	matchingItems := getMatchingItems(a, b)

	if !slices.Equal(matchingItems, []TileType{Coast, Land}) {
		t.Errorf("Ur shit whack bro %+v", matchingItems)
	}
}
