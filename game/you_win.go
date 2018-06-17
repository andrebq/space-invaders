package game

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/render"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	// YouWin is responsible for controlling the YouWin ship
	YouWin struct {
		*render.Sprite
	}
)

var (
	youWinKey = new(int)
)

// NewYouWin returns a new YouWin ship
func NewYouWin(youWinFrames, youWinAnimation string) (*YouWin, error) {
	sp, err := render.NewSprite(youWinFrames, youWinAnimation, 100)
	if err != nil {
		return nil, errors.Wrap(err, "game:you_win unable to load YouWin sprite")
	}

	p := &YouWin{
		Sprite: sp,
	}

	return p, nil
}

// Key implements indexed entities
func (p *YouWin) Key() interface{} {
	return youWinKey
}

// MoveTo changes the llc (lower-left-corner) of the YouWin
func (p *YouWin) MoveTo(target sdl.Point) {
	p.Pos.X = target.X
	p.Pos.Y = target.Y
}

// Update implements the interface required by dynamic system
func (p *YouWin) Update(dt float64, w *ces.World) {
	p.Sprite.UpdateAnimation(dt)
	if p.Sprite.Cycles() > 10 {
		w.RemoveEntity(p)
	}
}
