package game

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	// World holds the simulation parameters
	World struct {
		bounds sdl.Rect
	}
)

var (
	worldKey = new(int)
)

// NewWorld configures a new world
func NewWorld(bounds sdl.Rect) *World {
	return &World{
		bounds: bounds,
	}
}

// GetWorld returns the game world for the given ces.World
// it will panic if the world isn't available
func GetWorld(w *ces.World) *World {
	entity, ok := w.FindEntity(worldKey)
	if !ok {
		panic("game:world unable to find world entity")
	}

	actualW, ok := entity.(*World)
	if !ok {
		panic("game:world entity associated with world key isn't a *World object")
	}
	return actualW
}

// Key is the value used to index this object
func (w *World) Key() interface{} {
	return worldKey
}

// GetBounds return the size of the world
func (w *World) GetBounds() sdl.Rect {
	return w.bounds
}
