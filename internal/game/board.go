package game

type Position struct {
	Row int
	Col int
}

type Board struct {
	Size  int
	Cells [][]int
}

func NewBoard(size int) (*Board, error) {
	if size < 2 {
		return nil, newInvalidBoardSizeError("the board size must be at least 2")
	}

	board := &Board{
		Size:  size,
		Cells: make([][]int, size),
	}

	for i := range board.Cells {
		board.Cells[i] = make([]int, size)
	}

	return board, nil
}

func (b *Board) Get(p Position) int {
	return b.Cells[p.Row][p.Col]
}

func (b *Board) Set(p Position, value int) {
	b.Cells[p.Row][p.Col] = value
}

func (b *Board) IsEmpty(p Position) bool {
	return b.Get(p) == 0
}

func (b *Board) Reset() {
	for row := range b.Cells {
		clear(b.Cells[row])
	}
}

func (b *Board) GetEmptyCells() []Position {
	var empty []Position
	for row := range b.Cells {
		for col := range b.Cells[row] {
			p := Position{Row: row, Col: col}
			if b.IsEmpty(p) {
				empty = append(empty, p)
			}
		}
	}

	return empty
}

func (b *Board) HasEmptyCell() bool {
	for row := range b.Cells {
		for col := range b.Cells[row] {
			p := Position{Row: row, Col: col}
			if b.IsEmpty(p) {
				return true
			}
		}
	}

	return false
}

func (b *Board) HasTile(value int) bool {
	for row := 0; row < b.Size; row++ {
		for col := 0; col < b.Size; col++ {
			p := Position{Row: row, Col: col}
			if b.Get(p) == value {
				return true
			}
		}
	}

	return false
}
