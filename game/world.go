package game

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	// World holds the simulation parameters
	World struct {
		bounds sdl.Rect

		collidables collidables
	}

	bboxed interface {
		ces.Entity
		BBox() sdl.Rect
	}

	collidable interface {
		bboxed
		Collision(other ces.Entity)
	}
)

var (
	worldKey = new(int)
)

// NewWorld configures a new world
func NewWorld(bounds sdl.Rect) *World {
	return &World{
		bounds:      bounds,
		collidables: make(collidables, 0),
	}
}

// GetWorld returns the game world for the given ces.World
// it will panic if the world isn't available
func GetWorld(w *ces.World) *World {
	entity, ok := w.FindFirstEntity(worldKey)
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

// GetCentered returns the same rectangle but with its X/Y values
// moved to a specific value which will put the center if rectangle A
// in the center of the world bounds
func (w *World) GetCentered(r sdl.Rect) sdl.Rect {
	b := w.GetBounds()
	r.X = (b.W-b.X)/2 - (r.W / 2)
	r.Y = (b.H-b.Y)/2 - (r.H / 2)
	return r
}

func (w *World) addCollidable(c collidable) {
	w.collidables.add(c)
}

func (w *World) removeCollidable(c collidable) {
	w.collidables.remove(c)
}

func (w *World) checkCollision(e bboxed) bool {
	box := e.BBox()
	var collision bool
	for _, v := range w.collidables {
		vbox := v.BBox()
		if _, ok := vbox.Intersect(&box); ok {
			v.Collision(e)
			collision = true
		}
	}
	return collision
}
