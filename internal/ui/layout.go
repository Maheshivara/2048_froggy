package ui

type ScreenSize struct {
	Width  int
	Height int
}

type Layout struct {
	Screen ScreenSize

	BoardX      int
	BoardY      int
	BoardPixels int
	TileSize    int
	Gap         int

	PanelX     int
	PanelY     int
	PanelWidth int

	TitleFontSize int
	ScoreFontSize int
	TileFontSize  int
}

func NewLayout(width, height, boardSize int) Layout {
	screen := ScreenSize{width, height}

	if boardSize <= 0 {
		boardSize = 4
	}

	padding := 12
	panelGap := 20
	tileGap := 6

	boardPixels := min(
		screen.Height-2*padding,
		screen.Width*55/100,
	)

	tileSize := (boardPixels - (boardSize+1)*tileGap) / boardSize
	panelX := padding + boardPixels + panelGap

	return Layout{
		Screen:      screen,
		BoardX:      padding,
		BoardY:      (screen.Height - boardPixels) / 2,
		BoardPixels: boardPixels,
		TileSize:    tileSize,
		Gap:         tileGap,
		PanelX:      panelX,
		PanelY:      padding,
		PanelWidth:  screen.Width - panelX - padding,

		TitleFontSize: screen.Height / 12,
		ScoreFontSize: screen.Height / 14,
		TileFontSize:  max(10, tileSize/3),
	}
}

func (l Layout) TilePosition(row, col int) (int, int) {
	return l.BoardX + l.Gap + col*(l.TileSize+l.Gap),
		l.BoardY + l.Gap + row*(l.TileSize+l.Gap)
}
