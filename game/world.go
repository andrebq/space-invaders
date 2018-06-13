package game

import "github.com/veandco/go-sdl2/sdl"

type (
	World struct {
		bounds sdl.Rect
	}
)

// GetBounds return the size of the world
func (w *World) GetBounds() sdl.Rect {
	return w.bounds
}
