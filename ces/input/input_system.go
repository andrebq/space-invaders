package input

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/andrebq/space-invaders/ces"
)

type (
	// System is the input system
	System struct {
		ces.BaseSystem

		quit bool
	}

	keyAwareEntity interface {
		KeyUp(ev *sdl.KeyboardEvent, w *ces.World)
		KeyDown(ev *sdl.KeyboardEvent, w *ces.World)
	}

	keyAwareEntityFilterType struct{}
)

// Filter implements ces.EntityFilter
func (keyAwareEntityFilterType) Filter(e ces.Entity) bool {
	_, ok := e.(keyAwareEntity)
	return ok
}

var (
	theInput *System

	keyAwareEntityFilter keyAwareEntityFilterType

	lock sync.Mutex
)

// Get returns the global input system (no need for more than one)
func Get() *System {
	lock.Lock()
	defer lock.Unlock()

	if theInput == nil {
		theInput = &System{
			BaseSystem: ces.BaseSystem{},
		}
	}

	return theInput
}

// EntityFilter implements ces.System
func (s *System) EntityFilter() ces.EntityFilter {
	return keyAwareEntityFilter
}

// ShouldQuit returns true if sdl.QuitEvent was received
func (s *System) ShouldQuit() bool {
	return s.quit
}

// Input implements ces.InputSystem
func (s *System) Input(dt float64, w *ces.World) {
	for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {
		switch ev := ev.(type) {
		case *sdl.QuitEvent:
			s.quit = true
		case *sdl.KeyboardEvent:
			if ev.Type == sdl.KEYDOWN {
				s.dispatchKeyEvent(ev, w)
			} else if ev.Type == sdl.KEYUP {
				s.dispatchKeyEvent(ev, w)
			}
			logrus.WithField("ev", fmt.Sprintf("%#v", ev)).Debug()
		}
	}
}

func (s *System) dispatchKeyEvent(ev *sdl.KeyboardEvent, w *ces.World) {
	s.BaseSystem.ForEach(func(e ces.Entity) {
		if ev.Type == sdl.KEYDOWN {
			e.(keyAwareEntity).KeyDown(ev, w)
		} else if ev.Type == sdl.KEYUP {
			e.(keyAwareEntity).KeyUp(ev, w)
		}
	})
}
