package game

import (
	"github.com/andrebq/space-invaders/ces"
)

type (
	// basic arcade style of dynamic system
	arcade struct {
		ces.BaseSystem
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
func NewArcade() ces.DynamicSystem {
	return &arcade{}
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
}
