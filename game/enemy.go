package game

import (
	"math/rand"

	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/math"
	"github.com/andrebq/space-invaders/ces/render"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	// Enemy is responsible for controlling the Enemy ship
	Enemy struct {
		*render.Sprite

		direction    int32
		reverseCount int
		killed       bool
	}
)

var (
	enemyKey = new(int)
)

// NewEnemy returns a new Enemy ship
func NewEnemy(enemyFrames, enemyAnimation string) (*Enemy, error) {
	sp, err := render.NewSprite(enemyFrames, enemyAnimation, 99)
	if err != nil {
		return nil, errors.Wrap(err, "game:Enemy unable to load Enemy sprite")
	}

	p := &Enemy{
		Sprite:    sp,
		direction: randDirection(rand.Float32()),
	}

	return p, nil
}

// Key implements index entities
func (p *Enemy) Key() interface{} {
	return enemyKey
}

// OnAdd implements lifecycle aware interface
func (p *Enemy) OnAdd(w *ces.World) {
	GetWorld(w).addCollidable(p)
}

// OnRemove implements lifecycle aware interface
func (p *Enemy) OnRemove(w *ces.World) {
	GetWorld(w).removeCollidable(p)
}

// BBox returns the bounding box for ths object
func (p *Enemy) BBox() sdl.Rect {
	return p.RectAt(p.Pos)
}

// Collision is called when a colliction with another object happens
func (p *Enemy) Collision(e ces.Entity) {
	g, ok := e.(*Gun)
	if !ok {
		return
	}
	if g.consumed {
		return
	}
	// consumes the bullet to prevent it from
	// killing another ET
	g.consumed = true
	p.killed = true
}

// SetDirection changes the direction of the enemy
// positive values move them to the right
// negative to the left
// zero stops the movement
func (p *Enemy) SetDirection(v int) {
	switch {
	case v > 0:
		p.direction = 1
	case v < 0:
		p.direction = -1
	default:
		p.direction = 0
	}
}

// MoveTo changes the llc (lower-left-corner) of the Enemy
func (p *Enemy) MoveTo(target sdl.Point) {
	p.Pos.X = target.X
	p.Pos.Y = target.Y
}

// Update implements the interface required by dynamic system
func (p *Enemy) Update(dt float64, w *ces.World) {
	if p.killed {
		w.RemoveEntity(p)
		CreateExplosion(w, p.Pos)
	}

	if p.reverseCount > 0 {
		p.nextLane()
	}

	npos := p.Pos

	change := int32(dt*200) * p.direction
	npos.X += change
	gameWorld := GetWorld(w)
	rect := gameWorld.GetBounds()

	p.Sprite.UpdateAnimation(dt)

	if !math.FullyInside(rect, p.Sprite.RectAt(npos)) {
		p.reverseDirection()
		return
	}
	p.Pos = npos
}

func (p *Enemy) nextLane() {
	p.Pos.Y -= 40
	p.reverseCount = 0
}

func (p *Enemy) reverseDirection() {
	p.reverseCount++
	if p.direction < 0 {
		p.direction = 1
		return
	}
	p.direction = -1
}

func randDirection(f float32) int32 {
	if f > 0.5 {
		return 1
	}
	return -1
}
