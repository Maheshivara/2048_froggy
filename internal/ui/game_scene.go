package ui

import (
	"github.com/Maheshivara/froggy2048/internal/game"
	"github.com/Maheshivara/gofroggy/pkg/gfx"
	"github.com/Maheshivara/gofroggy/pkg/hardware/input"
	"github.com/Maheshivara/gofroggy/pkg/hardware/input/mapping"
	"github.com/Maheshivara/gofroggy/pkg/scene"
)

type GameScene struct {
	manager  *scene.SceneManager
	game     *game.Game
	renderer *Renderer

	boardSize int
	target    int

	screenWidth  int
	screenHeight int
	fontPath     string

	pendingRestart bool
	pendingBack    bool
}

var _ scene.Scene = (*GameScene)(nil)

func NewGameScene(
	boardSize, target, screenWidth, screenHeight int,
	fontPath string,
) *GameScene {
	return &GameScene{
		boardSize:    boardSize,
		target:       target,
		screenWidth:  screenWidth,
		screenHeight: screenHeight,
		fontPath:     fontPath,
		renderer: NewRenderer(
			screenWidth,
			screenHeight,
			boardSize,
			fontPath,
		),
	}
}

func (s *GameScene) Init(mgr *scene.SceneManager) error {
	s.manager = mgr
	return s.Restart()
}

func (s *GameScene) Restart() error {
	g, err := game.New(s.boardSize, s.target)
	if err != nil {
		return err
	}

	s.game = g
	return nil
}

func (s *GameScene) Update(
	_ float64,
	_ *input.InputController,
) error {
	if s.pendingBack {
		s.pendingBack = false

		menuScene := NewMenuScene(
			s.screenWidth,
			s.screenHeight,
			s.target,
			s.fontPath,
		)

		return s.manager.Switch(menuScene)
	}

	if s.pendingRestart {
		s.pendingRestart = false
		return s.Restart()
	}

	return nil
}

func (s *GameScene) Draw(d *gfx.Drawer) {
	if s.game == nil {
		return
	}

	s.renderer.Draw(d, s.game)
}

func (s *GameScene) HandleInput(
	btn mapping.InputButton,
	pressed bool,
) bool {
	if !pressed || s.game == nil {
		return false
	}

	if btn == mapping.SELECT {
		s.pendingBack = true
		return true
	}

	if btn == mapping.START {
		s.pendingRestart = true
		return true
	}

	if s.pendingBack || s.pendingRestart ||
		s.game.Status() != game.Playing {
		return false
	}

	switch btn {
	case mapping.UP:
		s.game.Move(game.Up)
	case mapping.DOWN:
		s.game.Move(game.Down)
	case mapping.LEFT:
		s.game.Move(game.Left)
	case mapping.RIGHT:
		s.game.Move(game.Right)
	default:
		return false
	}

	return true
}

func (s *GameScene) Pause() {}

func (s *GameScene) Resume() {}

func (s *GameScene) Destroy() {
	s.game = nil
	s.renderer = nil
}
