package main

import (
	"fmt"
	"time"

	"github.com/Maheshivara/froggy2048/internal/debug"
	"github.com/Maheshivara/froggy2048/internal/ui"
	"github.com/Maheshivara/gofroggy/pkg/gfx"
	"github.com/Maheshivara/gofroggy/pkg/hardware/display"
	"github.com/Maheshivara/gofroggy/pkg/hardware/input"
	"github.com/Maheshivara/gofroggy/pkg/hardware/input/mapping"
	"github.com/Maheshivara/gofroggy/pkg/scene"
)

const (
	FontPath = ""
	Target   = 2048
)

var displayConfig = display.Config{
	Backend:         display.BackendDirect,
	Width:           640,
	Height:          480,
	Scaling:         display.ScalingCenter,
	TrackDrawDamage: true,
	BackgroundColor: ui.ColorBackground,
	VSync:           false,
	UpdateScreen:    false,
	HideSystemOSD:   true,
}

type Application struct {
	Display display.DisplayController
	Input   *input.InputController
	Drawer  *gfx.Drawer
	Scenes  *scene.SceneManager

	Running bool

	prevButtons map[mapping.InputButton]bool
}

func (app *Application) Init() error {
	debug.Logf("Init: starting display initialization")
	disp, err := display.NewDisplayController(displayConfig)
	if err != nil {
		return fmt.Errorf("display initialization: %w", err)
	}
	app.Display = disp
	debug.Logf("Display initialized")

	width, height := disp.LogicalSize()

	controller, err := input.NewInputController(nil)
	if err != nil {
		return fmt.Errorf("input initialization: %w", err)
	}
	app.Input = controller
	debug.Logf("Input initialized")

	app.Drawer = gfx.NewDrawer(disp)
	app.Scenes = scene.NewSceneManager()

	app.prevButtons = make(map[mapping.InputButton]bool)

	menuScene := ui.NewMenuScene(
		displayConfig.Width,
		displayConfig.Height,
		Target,
		FontPath,
	)

	if err := app.Scenes.Push(menuScene); err != nil {
		return fmt.Errorf("menu scene initialization: %w", err)
	}

	app.Running = true
	debug.Logf("Init: completed; logical screen=%dx%d", width, height)
	return nil
}

func (app *Application) Update(dt float64) error {
	app.Input.Update()

	if app.Input.IsPressed(mapping.SELECT) &&
		app.Input.IsPressed(mapping.START) {
		debug.Logf("Exit requested: SELECT + START")
		app.Running = false
		return nil
	}

	app.forwardInput()

	return app.Scenes.Update(dt, app.Input)
}

func (app *Application) forwardInput() {
	buttons := [...]mapping.InputButton{
		mapping.UP,
		mapping.DOWN,
		mapping.LEFT,
		mapping.RIGHT,
		mapping.START,
		mapping.SELECT,
	}

	selectHeld := app.Input.IsPressed(mapping.SELECT)

	for _, btn := range buttons {
		current := app.Input.IsPressed(btn)
		previous := app.prevButtons[btn]

		if btn == mapping.START && selectHeld {
			app.prevButtons[btn] = current
			continue
		}

		if current != previous {
			debug.Logf("Input: button=%v pressed=%t", btn, current)
			app.Scenes.HandleInput(btn, current)
		}

		app.prevButtons[btn] = current
	}
}

func (app *Application) Draw() error {
	app.Display.Clear(ui.ColorBackground)

	app.Scenes.Draw(app.Drawer)

	if err := app.Display.Update(); err != nil {
		return fmt.Errorf("display update: %w", err)
	}

	return nil
}

func (app *Application) Run() error {

	ticker := time.NewTicker(time.Second / 60)
	defer ticker.Stop()

	const dt = 1.0 / 60.0

	debug.Logf("Run: entering main loop")

	for app.Running {
		if err := app.Update(dt); err != nil {
			return err
		}

		if !app.Running {
			break
		}

		if err := app.Draw(); err != nil {
			return err
		}

		<-ticker.C
	}

	debug.Logf("Run: main loop ended")
	return nil
}

func (app *Application) Close() {
	debug.Logf("Close: releasing resources")

	if app.Scenes != nil {
		app.Scenes.Pop()
	}

	if app.Input != nil {
		if err := app.Input.Close(); err != nil {
			debug.Logf("Close ERROR: input: %v", err)
		}
	}

	if app.Display != nil {
		app.Display.Close()
	}
}

func main() {
	debug.Init()
	defer debug.Close()
	defer func() {
		if r := recover(); r != nil {
			debug.Panic(r)
		}
	}()
	debug.Logf("=== 2048 starting ===")

	app := &Application{}
	defer app.Close()

	if err := app.Init(); err != nil {
		debug.Logf("Init ERROR: %v", err)
		return
	}

	if err := app.Run(); err != nil {
		debug.Logf("Run ERROR: %v", err)
		return
	}
	debug.Logf("=== 2048 finished ===")
}
