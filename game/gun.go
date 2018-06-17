package game

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/math"
	"github.com/andrebq/space-invaders/ces/render"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	// Gun is responsible for controlling the Gun ship
	Gun struct {
		*render.Sprite

		direction int32
		consumed  bool
	}
)

// NewGun returns a new Gun ship
func NewGun(gunFrames, gunAnimation string) (*Gun, error) {
	sp, err := render.NewSprite(gunFrames, gunAnimation, 98)
	if err != nil {
		return nil, errors.Wrap(err, "game:gun unable to load Gun sprite")
	}

	p := &Gun{
		Sprite:    sp,
		direction: 1,
	}

	return p, nil
}

// BBox implements bboxed interface
func (p *Gun) BBox() sdl.Rect {
	return p.RectAt(p.Pos)
}

// SetDirection changes the direction of the Gun
// positive values move them to the right
// negative to the left
// zero stops the movement
func (p *Gun) SetDirection(v int) {
	switch {
	case v > 0:
		p.direction = 1
	case v < 0:
		p.direction = -1
	default:
		p.direction = 0
	}
}

// MoveTo changes the llc (lower-left-corner) of the Gun
func (p *Gun) MoveTo(target sdl.Point) {
	p.Pos.X = target.X
	p.Pos.Y = target.Y
}

// Update implements the interface required by dynamic system
func (p *Gun) Update(dt float64, w *ces.World) {
	npos := p.Pos
	change := int32(dt*400) * p.direction
	npos.Y += change
	gameWorld := GetWorld(w)
	rect := gameWorld.GetBounds()

	p.Sprite.UpdateAnimation(dt)

	if gameWorld.checkCollision(p) {
		w.RemoveEntity(p)
		return
	}

	if !math.FullyInside(rect, p.Sprite.RectAt(npos)) {
		w.RemoveEntity(p)
		return
	}
	p.Pos = npos
}
