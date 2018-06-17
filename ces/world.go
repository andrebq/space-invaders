package ces

type (
	// World holds the system and entitys
	World struct {
		entities map[Entity]bool
		systems  map[System]bool

		entitiesIndex map[interface{}][]Entity
	}
)

// NewWorld constructs a new world
func NewWorld(systems ...System) *World {
	w := &World{
		entities:      make(map[Entity]bool),
		entitiesIndex: make(map[interface{}][]Entity),
		systems:       make(map[System]bool),
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

func (w *World) addToIndex(e Entity) {
	ie, ok := e.(indexableEntity)
	if ok {
		w.entitiesIndex[ie.Key()] = append(w.entitiesIndex[ie.Key()], e)
	}
}

func (w *World) removeFromIndex(e Entity) {
	ie, ok := e.(indexableEntity)
	if ok {
		slice := w.entitiesIndex[ie.Key()]
		if len(slice) == 0 {
			return
		}
		for i, v := range slice {
			if v == e {
				slice[i] = nil
				switch i {
				case len(slice) - 1:
					slice = slice[:len(slice)-1]
				default:
					slice = append(slice[:i], slice[i+1:]...)
				}
			}
		}
		w.entitiesIndex[ie.Key()] = slice
	}
}

// AddEntity is called to add a new entity to the world
// add entity will handle lifecycle events
func (w *World) AddEntity(entities ...Entity) {
	for _, e := range entities {
		w.entities[e] = true
		w.addToIndex(e)

		if lcae, ok := e.(lifecycleAwareEntity); ok {
			lcae.OnAdd(w)
		}
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
			w.entities[e] = false
			w.removeFromIndex(e)

			if lcae, ok := e.(lifecycleAwareEntity); ok {
				lcae.OnRemove(w)
			}
		}
	}
}

// FindFirstEntity returns the first entity indexed by the given key
func (w *World) FindFirstEntity(key interface{}) (Entity, bool) {
	e, ok := w.entitiesIndex[key]
	if ok {
		return e[0], true
	}
	return nil, false
}

// FindAllEntities returns the entities indexed by the given key
func (w *World) FindAllEntities(key interface{}) ([]Entity, bool) {
	e, ok := w.entitiesIndex[key]
	if !ok {
		return nil, ok
	}
	return e, len(e) > 0
}
