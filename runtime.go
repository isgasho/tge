package tge

import (
	physics "github.com/thommil/tge/physics"
	player "github.com/thommil/tge/player"
	renderer "github.com/thommil/tge/renderer"
	ui "github.com/thommil/tge/ui"
)

// App defines API to implement for TGE applications
type App interface {
	Create(settings Settings) error
	Start(runtime Runtime) error
	Resize(width int, height int)
	Resume()
	Render(renderer renderer.Renderer, ui ui.UI, player player.Player)
	Tick(physics physics.Physics)
	Pause()
	Dispose() error
}

// Instanciate is the main entry point
func Instanciate(app App) {
	app.Create(Settings{})
	//app.Start()
	//app.Resume()
	//app.Resize()

	//app.Render()
	//app.Tick()

	//app.Dispose()
}

// Runtime API
type Runtime interface {
}
