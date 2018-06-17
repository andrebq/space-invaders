package game

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/math"
	"github.com/andrebq/space-invaders/ces/render"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	// Player is responsible for controlling the player ship
	Player struct {
		*render.Sprite

		direction int32

		// how many time until the next fire
		gunCooldown float64
	}
)

var (
	playerKey = new(int)
)

// NewPlayer returns a new player ship
func NewPlayer(shipFrames, shipAnimation string) (*Player, error) {
	sp, err := render.NewSprite(shipFrames, shipAnimation, 100)
	if err != nil {
		return nil, errors.Wrap(err, "game:player unable to load player sprite")
	}

	p := &Player{
		Sprite:    sp,
		direction: 0,
	}

	return p, nil
}

// Key implements indexed entities
func (p *Player) Key() interface{} {
	return playerKey
}

// MoveTo changes the llc (lower-left-corner) of the player
func (p *Player) MoveTo(target sdl.Point) {
	p.Pos.X = target.X
	p.Pos.Y = target.Y
}

// KeyDown is called when the player starts to change to a given direction
func (p *Player) KeyDown(ev *sdl.KeyboardEvent, w *ces.World) {
	switch ev.Keysym.Sym {
	case sdl.K_LEFT:
		p.direction = -1
	case sdl.K_RIGHT:
		p.direction = 1
	}
}

// KeyUp makes the player stop the movement
func (p *Player) KeyUp(ev *sdl.KeyboardEvent, w *ces.World) {
	switch ev.Keysym.Sym {
	case sdl.K_LEFT:
		if p.direction == -1 {
			p.direction = 0
		}
	case sdl.K_RIGHT:
		if p.direction == 1 {
			p.direction = 0
		}
	}
}

// Update implements the interface required by dynamic system
func (p *Player) Update(dt float64, w *ces.World) {
	_, enemiesAlive := w.FindAllEntities(enemyKey)
	if !enemiesAlive {
		p.direction = 0
	}

	npos := p.Pos
	change := int32(dt*300) * p.direction
	npos.X += change
	gameWorld := GetWorld(w)
	rect := gameWorld.GetBounds()

	if sdl.GetKeyboardState()[sdl.SCANCODE_SPACE] > 0 {
		p.tryFireGun(w)
	}

	p.Sprite.UpdateAnimation(dt)

	if !math.FullyInside(rect, p.Sprite.RectAt(npos)) && enemiesAlive {
		// do not change the position
		return
	} else if math.FullyOutside(rect, p.Sprite.RectAt(npos)) && !enemiesAlive {
		CreateYouWin(w)
		w.RemoveEntity(p)
	}
	p.gunCooldown -= dt

	if !enemiesAlive {
		change = int32(dt * 500)
		npos.Y += change
	}

	p.Pos = npos
}

func (p *Player) tryFireGun(w *ces.World) {
	if p.gunCooldown <= 0 {
		CreatePlayerGun(w, p)
		p.gunCooldown = 1.0 / 5 /* 100 shots per second */
	}
}
