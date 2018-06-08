package ces

type (
	// BaseSystem is a no-op system that watches all entities and doesn't implement any
	// specific system
	BaseSystem struct {
		watch map[Entity]bool
	}

	staticFilter bool
)

// Filter implements ces.EntityFilter
func (s staticFilter) Filter(e Entity) bool {
	return bool(s)
}

// EntityFilter implements ces.System
func (b *BaseSystem) EntityFilter() EntityFilter {
	return staticFilter(true)
}

// EntityLifecycle implements ces.System
func (b *BaseSystem) EntityLifecycle(_ EntityLifecycleEvent) {
}

// Watch adds/removes the given entity to the watch list
func (b *BaseSystem) Watch(e Entity, v bool) {
	if b.watch == nil {
		b.watch = make(map[Entity]bool)
	}
	if v {
		b.watch[e] = true
	} else {
		delete(b.watch, e)
	}
}

// ForEach performs the given function for each entity
func (b *BaseSystem) ForEach(fn func(Entity)) {
	for e := range b.watch {
		fn(e)
	}
}
