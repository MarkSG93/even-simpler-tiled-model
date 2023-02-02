package main

import (
	"math"
	"testing"

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

func Equal(a, b []TileType) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func shouldAddRule(entries []TileType, newTile TileType) bool {
	return newTile != None && slices.Index(entries, newTile) == -1
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

func TestGenerateRulesFromSampleInput(t *testing.T) {
	sampleInput := [][]string{
		{"L", "C", "S"},
		{"L", "C", "S"},
		{"C", "S", "L"},
	}
	tileRulesMap := WaveFunction(sampleInput)

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

type NumberGenerator func() int
type Entropy func(numberGenerator NumberGenerator, grid [3][3]Square, totalCollapsed int) int
type Square struct {
	Possibilities []TileType
	Type          *TileType
}

func collapse(ruleSet map[TileType]TileRulesList, numberGenerator NumberGenerator, entropy Entropy) [3][3]Square {
	tileTypes := []TileType{Coast, Land, Sea}
	grid := [3][3]Square{}

	// fill all squares with possibilities
	for i, row := range grid {
		for j := 0; j < len(row); j++ {
			grid[i][j] = Square{Possibilities: tileTypes}
		}
	}

	totalCollapsed := 0
	gridSize := len(grid) * len(grid[0])
	for totalCollapsed < gridSize {
		tileNumber := entropy(numberGenerator, grid, totalCollapsed)
		row, col := 0, 0
		if tileNumber != 0 {
			row = int(math.Floor(float64(tileNumber) / float64(len(grid))))
			col = tileNumber % len(grid)
		}

		// decided the tile type
		if len(grid[row][col].Possibilities) < 1 {
			return collapse(ruleSet, numberGenerator, entropy)
		}

		collapsedTileType := &grid[row][col].Possibilities[numberGenerator()]
		grid[row][col].Type = collapsedTileType
		totalCollapsed++

		collapsedTileRuleSet := ruleSet[*collapsedTileType]
		// tile to the left
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
	}

	return grid
}

/*
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
} */

// Give a collapsed tile
// the surrounding tiles have their possibilities removed
// following the rules of the collapsed tile
func TestRemovePossibilities(t *testing.T) {

	numberGenerator := func() int {
		return 2
	}

	entropy := func(ng NumberGenerator, grid [3][3]Square, totalCollapsed int) int {
		if totalCollapsed == 0 {
			return 1
		}

		return 0
	}

	sampleInput := [][]string{
		{"L", "C", "S"},
	}
	ruleSet := WaveFunction(sampleInput)
	grid := collapse(ruleSet, numberGenerator, entropy)
	leftTile := grid[0][0]
	middleTile := grid[0][1]
	rightTile := grid[0][2]
	if *middleTile.Type != Coast {
		t.Errorf("Middle tile doesn't equal Coast, got %+v", *middleTile.Type)
	}
	if !slices.Equal(leftTile.Possibilities, []TileType{Land}) {
		t.Errorf("Left tile doesn't have the possibility of Land, got %+v", leftTile.Possibilities)
	}
	if !slices.Equal(rightTile.Possibilities, []TileType{Sea}) {
		t.Errorf("Right tile doesn't have the possibility of Sea, got %+v", rightTile.Possibilities)
	}
}

func TestRemovePossibilitiesMultipleRows(t *testing.T) {
	numberGenerator := func() int {
		return 2
	}

	entropy := func(ng NumberGenerator, grid [3][3]Square, totalCollapsed int) int {
		if totalCollapsed == 0 {
			return 1
		}

		return 5
	}

	sampleInput := [][]string{
		{"L", "C", "S"},
		{"C", "L", "C"},
	}
	ruleSet := WaveFunction(sampleInput)
	grid := collapse(ruleSet, numberGenerator, entropy)
	middleTile := grid[0][1]
	downTile := grid[1][1]
	rightTile := grid[1][2]
	upTile := grid[0][2]
	if *middleTile.Type != Coast {
		t.Errorf("Middle tile doesn't equal Coast, got %+v", *middleTile.Type)
	}
	if !slices.Equal(downTile.Possibilities, []TileType{Land}) {
		t.Errorf("Down tile doesn't have the possibility of Land, got %+v", downTile.Possibilities)
	}
	if *rightTile.Type != Coast {
		t.Errorf("Right tile doesn't equal Coast, got %+v", *rightTile.Type)
	}
	if !slices.Equal(upTile.Possibilities, []TileType{Land, Sea}) {
		t.Errorf("Up tile doesn't have the possibility of Sea and Land, got %+v", upTile.Possibilities)
	}
}

func TestCollapsesAllSquaresInAGrid(t *testing.T) {
	numberGenerator := func() int {
		return 0
	}

	entropy := func(ng NumberGenerator, grid [3][3]Square, totalCollapsed int) int {
		return totalCollapsed
	}

	sampleInput := [][]string{
		{"L", "C", "S"},
		{"C", "L", "S"},
		{"S", "C", "L"},
	}
	ruleSet := WaveFunction(sampleInput)
	grid := collapse(ruleSet, numberGenerator, entropy)
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
	if *thirdTile.Type != Sea {
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

func getMatchingItems(a []TileType, b []TileType) []TileType {
	hits := map[string]int{
		Land:  0,
		Sea:   0,
		Coast: 0,
	}
	for _, item := range a {
		hits[item] += 1
	}

	for _, item := range b {
		hits[item] += 1
	}

	matchingItems := []TileType{}
	for tileType, hits := range hits {
		if hits > 1 {
			matchingItems = append(matchingItems, tileType)
		}
	}

	return matchingItems
}

func TestGetMatchingItemsInSlices(t *testing.T) {
	a := []TileType{Land, Sea, Coast}
	b := []TileType{Land, Coast}
	matchingItems := getMatchingItems(a, b)

	if !slices.Equal(matchingItems, []TileType{Land, Coast}) {
		t.Errorf("Ur shit whack bro %+v", matchingItems)
	}
}
