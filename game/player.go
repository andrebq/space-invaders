package game

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/render"
	"github.com/pkg/errors"
)

type (
	// Player is responsible for controlling the player ship
	Player struct {
		*render.Sprite
	}
)

// NewPlayer returns a new player ship
func NewPlayer(shipFile string) (*Player, error) {
	sp, err := render.NewSprite(shipFile, 100)
	if err != nil {
		return nil, errors.Wrap(err, "game:player unable to load player sprite")
	}

	p := &Player{
		Sprite: sp,
	}

	return p, nil
}

// Update implements the interface required by dynamic system
func (p *Player) Update(dt float64, w *ces.World) {
	p.Sprite.Pos.X += int32(dt * 100)
}