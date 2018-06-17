package game

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/render"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	// EndAnimation is responsible for controlling the GameEnd ship
	EndAnimation struct {
		*render.Sprite
	}
)

var (
	endAnimationKey = new(int)
)

// NewGameEnd returns a new GameEnd ship
func NewGameEnd(gameEndFrames, gameEndAnimation string) (*EndAnimation, error) {
	sp, err := render.NewSprite(gameEndFrames, gameEndAnimation, 100)
	if err != nil {
		return nil, errors.Wrap(err, "game:game_end unable to load GameEnd sprite")
	}

	p := &EndAnimation{
		Sprite: sp,
	}

	return p, nil
}

// Key implements indexed entities
func (p *EndAnimation) Key() interface{} {
	return endAnimationKey
}

// MoveTo changes the llc (lower-left-corner) of the GameEnd
func (p *EndAnimation) MoveTo(target sdl.Point) {
	p.Pos.X = target.X
	p.Pos.Y = target.Y
}

// Update implements the interface required by dynamic system
func (p *EndAnimation) Update(dt float64, w *ces.World) {
	p.Sprite.UpdateAnimation(dt)
	if p.Sprite.Cycles() > 10 {
		w.RemoveEntity(p)
	}
}
