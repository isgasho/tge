// +build js

package tge

import (
	log "log"
	"time"
)

// -------------------------------------------------------------------- //
// Runtime implementation
// -------------------------------------------------------------------- //
type browserRuntime struct {
	app      App
	ticker   *time.Ticker
	isPaused bool
}

func (runtime browserRuntime) Stop() {
	runtime.ticker.Stop()
	runtime.app.OnPause()
	runtime.app.OnStop()
}

func doRun(app App, settings *Settings) error {
	log.Println("doRun()")

	// -------------------------------------------------------------------- //
	// Init
	// -------------------------------------------------------------------- //

	// Instanciate Runtime
	browserRuntime := browserRuntime{
		app:      app,
		isPaused: true,
	}

	// -------------------------------------------------------------------- //
	// Ticker Loop
	// -------------------------------------------------------------------- //
	tpsDelay := time.Duration(1000000000 / settings.TPS)
	browserRuntime.ticker = time.NewTicker(tpsDelay)
	defer browserRuntime.ticker.Stop()
	go func() {
		for range browserRuntime.ticker.C {
			if !browserRuntime.isPaused {
				app.OnTick(tpsDelay)
			}
		}
	}()

	return nil
}
