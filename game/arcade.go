package game

import (
	"context"

	"github.com/andrebq/space-invaders/ces"
)

type (
	// basic arcade style of dynamic system
	arcade struct {
		ces.BaseSystem

		stop context.CancelFunc
	}

	onlyDynamicEntities struct{}

	dynamicEntity interface {
		Update(dt float64, w *ces.World)
	}
)

var (
	theFilter onlyDynamicEntities
)

func (onlyDynamicEntities) Filter(e ces.Entity) bool {
	_, ok := e.(dynamicEntity)
	return ok
}

// NewArcade returns a simple arcade dynamic system
func NewArcade(stop context.CancelFunc) ces.DynamicSystem {
	return &arcade{stop: stop}
}

// EntityFilter returns the expected entity filter
func (a *arcade) EntityFilter() ces.EntityFilter {
	return theFilter
}

// Update implements ces.DynamicSystem
func (a *arcade) Update(dt float64, w *ces.World) {
	// usually for a simple arcade game
	// there is no need to step down the simulation
	// so we just call the entities with the given dt

	a.BaseSystem.ForEach(func(e ces.Entity) {
		e.(dynamicEntity).Update(dt, w)
	})

	_, playerAlive := w.FindAllEntities(playerKey)
	_, endAnimationAlive := w.FindAllEntities(endAnimationKey)

	if !playerAlive && !endAnimationAlive {
		a.stop()
	}
}
