package ces

import "time"

// Iterate runs one full iteration (input, update, render, renderhud)
func (w *World) Iterate(oldTime time.Time) time.Time {
	now := time.Now()
	secs := now.Sub(oldTime).Seconds()
	w.Input(secs)
	w.Update(secs)
	w.Render(secs)
	w.RenderHUD(secs)
	return time.Now()
}

// Update is when system actually have a chance do change the world
func (w *World) Update(dt float64) {
	for s := range w.systems {
		if s, ok := s.(DynamicSystem); ok {
			s.Update(dt, w)
		}
	}
}

// Input is called to notify about user input
func (w *World) Input(dt float64) {
	for s := range w.systems {
		if s, ok := s.(InputSystem); ok {
			s.Input(dt)
		}
	}
}

// Render is called to notify about user Render
func (w *World) Render(dt float64) {
	for s := range w.systems {
		if s, ok := s.(RenderSystem); ok {
			s.Render(dt)
		}
	}
}

// RenderHUD is called to notify about user input
func (w *World) RenderHUD(dt float64) {
	for s := range w.systems {
		if s, ok := s.(HUDSystem); ok {
			s.RenderHUD(dt)
		}
	}
}
