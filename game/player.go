package game

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/math"
	"github.com/andrebq/space-invaders/ces/render"
	"github.com/pkg/errors"
)

type (
	// Player is responsible for controlling the player ship
	Player struct {
		*render.Sprite

		direction int32
	}
)

// NewPlayer returns a new player ship
func NewPlayer(shipFile string) (*Player, error) {
	sp, err := render.NewSprite(shipFile, 100)
	if err != nil {
		return nil, errors.Wrap(err, "game:player unable to load player sprite")
	}

	p := &Player{
		Sprite:    sp,
		direction: 1,
	}

	return p, nil
}

// Update implements the interface required by dynamic system
func (p *Player) Update(dt float64, w *ces.World) {
	npos := p.Pos
	change := int32(dt*100) * p.direction
	npos.X += change
	gameWorld := GetWorld(w)

	rect := gameWorld.GetBounds()
	if !math.FullyInside(&rect, &p.Pos) {
		p.reverseDirection()
		npos.X -= 2 * change /* remove the last update */
	}
	p.Pos = npos
}

func (p *Player) reverseDirection() {
	if p.direction > 0 {
		p.direction = -1
		return
	}
	p.direction = 1
}
