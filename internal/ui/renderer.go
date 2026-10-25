package ui

import (
	"strconv"

	"github.com/Maheshivara/froggy2048/internal/game"
	"github.com/Maheshivara/gofroggy/pkg/gfx"
	"github.com/Maheshivara/gofroggy/pkg/gfx/primitives"
	"github.com/Maheshivara/gofroggy/pkg/gfx/text"
)

type Renderer struct {
	layout   Layout
	fonts    map[int]*text.Font
	fontPath string
}

func (r *Renderer) fontForSize(size int) *text.Font {
	if size < 8 {
		size = 8
	}

	if font, ok := r.fonts[size]; ok {
		return font
	}

	font, _ := text.LoadFont(r.fontPath, float64(size))
	r.fonts[size] = font
	return font
}

func NewRenderer(
	width, height, boardSize int,
	fontPath string,
) *Renderer {
	return &Renderer{
		layout:   NewLayout(width, height, boardSize),
		fontPath: fontPath,
		fonts:    make(map[int]*text.Font),
	}
}

func (r *Renderer) Draw(d *gfx.Drawer, g *game.Game) {
	r.drawBoard(d, g)
	r.drawPanel(d, g)
}

func (r *Renderer) drawBoard(d *gfx.Drawer, g *game.Game) {
	l := r.layout

	r.drawRect(
		d,
		l.BoardX,
		l.BoardY,
		l.BoardPixels,
		l.BoardPixels,
		ColorBoard,
	)

	for row := 0; row < g.BoardSize(); row++ {
		for col := 0; col < g.BoardSize(); col++ {
			x, y := l.TilePosition(row, col)
			r.drawTile(d, x, y, g.TileAt(row, col))
		}
	}
}

func (r *Renderer) drawTile(
	d *gfx.Drawer,
	x, y, value int,
) {
	size := r.layout.TileSize

	r.drawRect(d, x, y, size, size, TileColor(value))

	if value == 0 {
		return
	}

	label := strconv.Itoa(value)

	fontSize := r.layout.TileFontSize
	switch {
	case value >= 1024:
		fontSize = fontSize * 65 / 100
	case value >= 128:
		fontSize = fontSize * 80 / 100
	}

	r.drawCenteredText(
		d,
		label,
		x+size/2,
		y+size/2,
		fontSize,
		TileTextColor(value),
	)
}

func (r *Renderer) drawPanel(d *gfx.Drawer, g *game.Game) {
	l := r.layout
	x := l.PanelX
	y := l.PanelY

	r.drawText(
		d, "SCORE", x, y,
		l.ScoreFontSize, ColorDarkText,
	)

	y += l.ScoreFontSize + 8

	r.drawText(
		d, strconv.Itoa(g.Score()), x, y,
		l.ScoreFontSize, ColorDarkText,
	)

	y += l.ScoreFontSize + 24

	r.drawText(
		d, "TARGET", x, y,
		l.ScoreFontSize, ColorDarkText,
	)
	y += l.ScoreFontSize + 8

	r.drawText(
		d, strconv.Itoa(g.Target()), x, y,
		l.ScoreFontSize, ColorDarkText,
	)

	y += l.ScoreFontSize + 28

	switch g.Status() {
	case game.Won:
		r.drawText(
			d, "YOU WON!", x, y,
			l.ScoreFontSize, ColorDarkText,
		)
	case game.Lost:
		r.drawText(
			d, "GAME OVER", x, y,
			l.ScoreFontSize, ColorDarkText,
		)
	}

	instructionsY := l.Screen.Height - 95

	r.drawText(
		d, "D-PAD: MOVE", x, instructionsY,
		20, ColorDarkText,
	)
	r.drawText(
		d, "START: RESTART", x, instructionsY+25,
		20, ColorDarkText,
	)
	r.drawText(
		d, "SELECT: MENU", x, instructionsY+45,
		20, ColorDarkText,
	)
}

func (r *Renderer) drawRect(
	d *gfx.Drawer,
	x, y, width, height int,
	color uint16,
) {
	rec := primitives.Rectangle{
		Origin: primitives.Point{
			X: x,
			Y: y,
		},
		Width:  width,
		Height: height,
	}
	config := primitives.DrawConfig{
		Color:     color,
		FillColor: color,
		IsFilled:  true,
	}
	d.Primitives.DrawRectangle(rec, config)
}

func (r *Renderer) drawText(
	d *gfx.Drawer,
	value string,
	x, y, fontSize int,
	color uint16,
) {
	font := r.fontForSize(fontSize)

	ascent, _ := font.GetMetrics()

	d.Text.DrawText(
		font,
		value,
		x,
		y+ascent,
		color,
	)
}

func (r *Renderer) drawCenteredText(
	d *gfx.Drawer,
	value string,
	centerX, centerY, fontSize int,
	color uint16,
) {
	font := r.fontForSize(fontSize)

	width, _ := font.MeasureText(value)
	ascent, descent := font.GetMetrics()

	x := centerX - width/2
	baselineY := centerY + (ascent-descent)/2

	d.Text.DrawText(font, value, x, baselineY, color)
}
