package game

import (
	"math/rand"
)

type Direction uint8

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Status uint8

const (
	Playing Status = iota
	Won
	Lost
)

type SpawningTile struct {
	Number int
	Weight int
}

var spawningTileOptions = [...]SpawningTile{SpawningTile{Number: 2, Weight: 9}, SpawningTile{Number: 4, Weight: 1}}

func normalizeTarget(target int) int {
	if target < 4 || (target&(target-1)) != 0 {
		return 2048
	}

	return target
}

type Game struct {
	board       *Board
	score       int
	target      int
	status      Status
	totalWeight int
	rng         *rand.Rand
}

func New(size, target int) (*Game, error) {
	return NewWithSeed(size, target, rand.Int63())
}

func NewWithSeed(size, target int, seed int64) (*Game, error) {
	target = normalizeTarget(target)

	board, err := NewBoard(size)
	if err != nil {
		return nil, err
	}

	var totalWeight int
	for _, tile := range spawningTileOptions {
		totalWeight += int(tile.Weight)
	}

	g := &Game{
		target:      target,
		board:       board,
		status:      Playing,
		totalWeight: totalWeight,
		rng:         rand.New(rand.NewSource(seed)),
	}

	g.SpawnTile()
	g.SpawnTile()

	return g, nil
}

func (g *Game) Status() Status {
	return g.status
}

func (g *Game) IsBoardFull() bool {
	return !g.board.HasEmptyCell()
}

func (g *Game) Score() int {
	return g.score
}

func (g *Game) Target() int {
	return g.target
}

func (g *Game) BoardTiles() [][]int {
	tiles := make([][]int, g.board.Size)

	for i, row := range g.board.Cells {
		tiles[i] = append([]int(nil), row...)
	}

	return tiles
}

func (g *Game) TileAt(row, col int) int {
	p := Position{Row: row, Col: col}
	return g.board.Get(p)
}

func (g *Game) BoardSize() int {
	return g.board.Size
}

func (g *Game) IncreaseScore(i int) {
	if i < 0 {
		return
	}
	g.score += i
}

func (g *Game) SpawnTile() bool {
	if g.Status() != Playing {
		return false
	}

	emptyCells := g.board.GetEmptyCells()

	if len(emptyCells) == 0 {
		return false
	}

	pos := emptyCells[g.rng.Intn(len(emptyCells))]
	number := g.randomTileNumber()

	g.board.Set(pos, int(number))

	return true

}

func (g *Game) randomTileNumber() int {
	if g.totalWeight == 0 {
		return 0
	}

	roll := g.rng.Intn(g.totalWeight)

	for _, tile := range spawningTileOptions {
		if roll < tile.Weight {
			return tile.Number
		}

		roll -= tile.Weight
	}

	panic("unreachable: invalid tile weights")
}

func (g *Game) UpdateStatus() {
	if g.status != Playing {
		return
	}

	if g.board.HasTile(g.target) {
		g.status = Won
		return
	}

	if !g.board.HasMoves() {
		g.status = Lost
	}
}
