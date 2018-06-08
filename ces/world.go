package ces

type (
	// World holds the system and entitys
	World struct {
		entities map[Entity]bool
		systems  map[System]bool
	}
)

// NewWorld constructs a new world
func NewWorld(systems ...System) *World {
	w := &World{
		entities: make(map[Entity]bool),
		systems:  make(map[System]bool),
	}
	w.addSystem(systems...)
	return w
}

// AddSystem add a new system to the world
func (w *World) addSystem(systems ...System) {
	for _, s := range systems {
		w.systems[s] = true
	}
}

// AddEntity is called to add a new entity to the world
// add entity will handle lifecycle events
func (w *World) AddEntity(entities ...Entity) {
	for _, e := range entities {
		w.entities[e] = true
	}

	for k := range w.systems {
		for _, e := range entities {
			ev := EntityLifecycleEvent{
				Entity:  e,
				Created: true,
			}
			if k.EntityFilter().Filter(e) {
				k.EntityLifecycle(ev)
			}
		}
	}
}

// RemoveEntity removes the given entity from the system
// entities are removed after the event is called
func (w *World) RemoveEntity(entities ...Entity) {
	for _, e := range entities {
		ev := EntityLifecycleEvent{
			Entity:  e,
			Deleted: true,
		}
		if w.entities[e] {
			for s := range w.systems {
				if s.EntityFilter().Filter(e) {
					s.EntityLifecycle(ev)
				}
			}
		}
	}
}
