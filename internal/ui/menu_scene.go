package ui

import (
	"fmt"

	"github.com/Maheshivara/gofroggy/pkg/gfx"
	"github.com/Maheshivara/gofroggy/pkg/gfx/text"
	"github.com/Maheshivara/gofroggy/pkg/hardware/input"
	"github.com/Maheshivara/gofroggy/pkg/hardware/input/mapping"
	"github.com/Maheshivara/gofroggy/pkg/scene"
)

var _ scene.Scene = (*MenuScene)(nil)

type MenuScene struct {
	manager *scene.SceneManager

	screenWidth  int
	screenHeight int
	fontPath     string

	selectedField int

	boardSizes     []int
	boardSizeIndex int

	targets     []int
	targetIndex int

	font      *text.Font
	titleFont *text.Font

	pendingStart bool
}

func NewMenuScene(
	screenWidth, screenHeight int,
	target int,
	fontPath string,
) *MenuScene {
	m := &MenuScene{
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		fontPath:     fontPath,

		boardSizes: []int{4, 5, 6},
		targets:    []int{512, 1024, 2048, 4096, 8192},

		boardSizeIndex: 0,
		targetIndex:    2,
	}

	for i, value := range m.targets {
		if value == target {
			m.targetIndex = i
			break
		}
	}

	return m
}

func (m *MenuScene) Init(manager *scene.SceneManager) error {
	m.manager = manager

	m.font, _ = text.LoadFont(m.fontPath, 20)
	m.titleFont, _ = text.LoadFont(m.fontPath, 36)

	return nil
}

func (m *MenuScene) Update(
	_ float64,
	_ *input.InputController,
) error {
	if !m.pendingStart {
		return nil
	}

	m.pendingStart = false

	boardSize := m.boardSizes[m.boardSizeIndex]
	target := m.targets[m.targetIndex]

	gameScene := NewGameScene(
		boardSize,
		target,
		m.screenWidth,
		m.screenHeight,
		m.fontPath,
	)

	return m.manager.Switch(gameScene)
}

func (m *MenuScene) HandleInput(
	btn mapping.InputButton,
	pressed bool,
) bool {
	if !pressed {
		return false
	}

	switch btn {
	case mapping.UP:
		m.selectedField--
		if m.selectedField < 0 {
			m.selectedField = 1
		}
		return true

	case mapping.DOWN:
		m.selectedField++
		if m.selectedField > 1 {
			m.selectedField = 0
		}
		return true

	case mapping.LEFT:
		m.changeValue(-1)
		return true

	case mapping.RIGHT:
		m.changeValue(1)
		return true

	case mapping.START:
		m.pendingStart = true
		return true
	}

	return false
}

func (m *MenuScene) changeValue(delta int) {
	switch m.selectedField {
	case 0:
		count := len(m.boardSizes)
		m.boardSizeIndex = (m.boardSizeIndex + delta + count) % count

	case 1:
		count := len(m.targets)
		m.targetIndex = (m.targetIndex + delta + count) % count
	}
}

func (m *MenuScene) Draw(d *gfx.Drawer) {
	centerX := m.screenWidth / 2

	m.drawCentered(
		d,
		m.titleFont,
		"2048",
		centerX,
		m.screenHeight/5,
		ColorDarkText,
	)

	m.drawCentered(
		d,
		m.font,
		"GAME SETTINGS",
		centerX,
		m.screenHeight/5+45,
		ColorDarkText,
	)

	boardSize := m.boardSizes[m.boardSizeIndex]
	target := m.targets[m.targetIndex]

	boardLabel := fmt.Sprintf(
		"BOARD: %d x %d",
		boardSize,
		boardSize,
	)

	targetLabel := fmt.Sprintf(
		"TARGET: %d",
		target,
	)

	boardColor := ColorDarkText
	targetColor := ColorDarkText

	if m.selectedField == 0 {
		boardLabel = "> " + boardLabel + " <"
		boardColor = 0xF800
	}

	if m.selectedField == 1 {
		targetLabel = "> " + targetLabel + " <"
		targetColor = 0xF800
	}

	m.drawCentered(
		d,
		m.font,
		boardLabel,
		centerX,
		m.screenHeight/2-20,
		boardColor,
	)

	m.drawCentered(
		d,
		m.font,
		targetLabel,
		centerX,
		m.screenHeight/2+25,
		targetColor,
	)

	m.drawCentered(
		d,
		m.font,
		"UP/DOWN: OPTION",
		centerX,
		m.screenHeight-85,
		ColorDarkText,
	)

	m.drawCentered(
		d,
		m.font,
		"LEFT/RIGHT: CHANGE",
		centerX,
		m.screenHeight-60,
		ColorDarkText,
	)

	m.drawCentered(
		d,
		m.font,
		"START: PLAY",
		centerX,
		m.screenHeight-35,
		ColorDarkText,
	)
}

func (m *MenuScene) drawCentered(
	d *gfx.Drawer,
	font *text.Font,
	value string,
	centerX, centerY int,
	color uint16,
) {
	width, _ := font.MeasureText(value)
	ascent, descent := font.GetMetrics()

	x := centerX - width/2
	baselineY := centerY + (ascent-descent)/2

	d.Text.DrawText(
		font,
		value,
		x,
		baselineY,
		color,
	)
}

func (m *MenuScene) Pause()   {}
func (m *MenuScene) Resume()  {}
func (m *MenuScene) Destroy() {}
