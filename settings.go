// Copyright (c) 2019 Thomas MILLET. All rights reserved.

package tge

// EventMask defines mask event for enable/disable events receivers at runtime level
type EventMask int

const (
	// AllEventsDisable disables all input events on App
	AllEventsDisable = 0x00
	// MouseButtonEventEnabled enabled mouse buttons events receiver on App
	MouseButtonEventEnabled = 0x01
	// MouseMotionEventEnabled enabled mouse motion events receiver on App
	MouseMotionEventEnabled = 0x02
	// ScrollEventEnabled enabled scroll event receiver on App
	ScrollEventEnabled = 0x04
	// KeyEventEnabled enabled key event receiver on App
	KeyEventEnabled = 0x08
	// AllEventsEnabled enables all input events on App
	AllEventsEnabled = 0xFFFF
)

// Settings definition of TGE application
type Settings struct {
	// Name of the App
	Name string `json:"name" yaml:"name"`
	// Fullscreen indicates if the app must be run in fullscreen mode
	Fullscreen bool `json:"fullscreen" yaml:"fullscreen"`
	// Width of the window if run windowed only
	Width int `json:"width" yaml:"width"`
	// Height of the window if run windowed only
	Height int `json:"height" yaml:"height"`
	// EventMask allows to enabled/disable events receiver on Runtime
	EventMask EventMask `json:"event_mask" yaml:"event_mask"`
}

// Default settings
var defaultSettings = Settings{
	Name:       "TGE Application",
	Fullscreen: false,
	Width:      640,
	Height:     480,
	EventMask:  AllEventsEnabled,
}
