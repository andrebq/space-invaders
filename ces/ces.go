package ces

type (
	// Entity represents any object in game world
	// entities work as a bag of components manipulated by a System
	Entity interface {
	}

	indexableEntity interface {
		Entity

		// Key returns the key used to index this entity to the system,
		// it shouldn't change after time
		Key() interface{}
	}

	// EntityLifecycleEvent holds information about events associated with an entity
	EntityLifecycleEvent struct {
		// Entity associated with the event
		Entity Entity

		// Created indicates that the given entity was just created
		Created bool

		// Deleted means the entity was removed was removed, this is the last event
		Deleted bool

		// Updated means the entity was updated by some other system and that system
		// wanted to notify others.
		//
		// Not every update needs to be notified
		Updated bool
	}

	// EntityFilter is used to filter if events from a given entity should be associated with the
	// given System
	EntityFilter interface {
		// Filter indicates if the entity is of interest or not
		Filter(Entity) bool
	}

	// System represents some behavior of the game
	System interface {
		// EntityFilter must return the filter used to decide if any event should be
		// notified.
		//
		// This method MUST return fast
		EntityFilter() EntityFilter

		// EntityLifecycle is called by the world to notify a system about an entity activity
		// all lifecylcle events are sent to this method except Created ones, which use the specific
		// handler
		EntityLifecycle(EntityLifecycleEvent)
	}

	// DynamicSystem is a system with an dynamic behavior
	DynamicSystem interface {
		System
		// Update is used to update entity components
		Update(dt float64, w *World)
	}

	// InputSystem is a system that process user input
	InputSystem interface {
		System

		// Input is called to notify the system that some input was obtained
		Input(dt float64)
	}

	// RenderSystem is a system to render game content
	RenderSystem interface {
		System

		// Render is called to allow the system to print something on the screen
		Render(float64)
	}

	// HUDSystem is a system to render the game UI
	HUDSystem interface {
		System

		// Render is called to allow the system to print something on the screen
		RenderHUD(float64)
	}
)
