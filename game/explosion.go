package game

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/render"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	// Explosion is responsible for controlling the Explosion ship
	Explosion struct {
		*render.Sprite
	}
)

// NewExplosion returns a new Explosion ship
func NewExplosion(explosionFrames, explosionAnimation string) (*Explosion, error) {
	sp, err := render.NewSprite(explosionFrames, explosionAnimation, 100)
	if err != nil {
		return nil, errors.Wrap(err, "game:Explosion unable to load Explosion sprite")
	}

	p := &Explosion{
		Sprite: sp,
	}

	return p, nil
}

// MoveTo changes the llc (lower-left-corner) of the Explosion
func (p *Explosion) MoveTo(target sdl.Point) {
	p.Pos.X = target.X
	p.Pos.Y = target.Y
}

// Update implements the interface required by dynamic system
func (p *Explosion) Update(dt float64, w *ces.World) {
	p.Sprite.UpdateAnimation(dt)
	if p.Sprite.Cycles() > 0 {
		w.RemoveEntity(p)
	}
}
