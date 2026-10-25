package game

func mergeLine(line []int) ([]int, int) {
	values := make([]int, 0, len(line))

	for _, value := range line {
		if value != 0 {
			values = append(values, value)
		}
	}

	result := make([]int, 0, len(line))
	score := 0

	for i := 0; i < len(values); i++ {
		if i+1 < len(values) && values[i] == values[i+1] {
			merged := values[i] * 2

			result = append(result, merged)
			score += merged

			i++
		} else {
			result = append(result, values[i])
		}
	}

	for len(result) < len(line) {
		result = append(result, 0)
	}

	return result, score
}

func movePosition(
	direction Direction,
	line int,
	index int,
	size int,
) Position {
	switch direction {
	case Left:
		return Position{Row: line, Col: index}

	case Right:
		return Position{Row: line, Col: size - 1 - index}

	case Up:
		return Position{Row: index, Col: line}

	case Down:
		return Position{Row: size - 1 - index, Col: line}

	default:
		panic("invalid direction")
	}
}

func (b *Board) HasMoves() bool {
	size := b.Size

	if b.HasEmptyCell() {
		return true
	}

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			p := Position{Row: row, Col: col}
			value := b.Get(p)

			np := Position{Row: row, Col: col + 1}
			if col+1 < size && value == b.Get(np) {
				return true
			}

			bn := Position{Row: row + 1, Col: col}
			if row+1 < size && value == b.Get(bn) {
				return true
			}
		}
	}

	return false
}

func (g *Game) Move(direction Direction) bool {
	if g.status != Playing {
		return false
	}

	size := g.BoardSize()
	changed := false
	totalScore := 0

	for i := 0; i < size; i++ {
		line := make([]int, size)

		for j := 0; j < size; j++ {
			p := movePosition(direction, i, j, size)
			line[j] = g.board.Get(p)
		}

		merged, score := mergeLine(line)
		totalScore += score

		for j := 0; j < size; j++ {
			p := movePosition(direction, i, j, size)

			if g.board.Get(p) != merged[j] {
				changed = true
				g.board.Set(p, merged[j])
			}
		}
	}

	if !changed {
		return false
	}

	g.IncreaseScore(totalScore)
	g.SpawnTile()
	g.UpdateStatus()

	return true
}
