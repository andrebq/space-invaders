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
		ces.System

		quit bool
	}
)

var (
	theInput *System

	lock sync.Mutex
)

// Get returns the global input system (no need for more than one)
func Get() *System {
	lock.Lock()
	defer lock.Unlock()

	if theInput == nil {
		theInput = &System{
			System: &ces.BaseSystem{},
		}
	}

	return theInput
}

// ShouldQuit returns true if sdl.QuitEvent was received
func (s *System) ShouldQuit() bool {
	return s.quit
}

// Input implements ces.InputSystem
func (s *System) Input(dt float64) {
	for ev := sdl.PollEvent(); ev != nil; ev = sdl.PollEvent() {
		switch ev := ev.(type) {
		case *sdl.QuitEvent:
			s.quit = true
		default:
			logrus.WithField("ev", fmt.Sprintf("%#v", ev)).Debug()
		}
	}
}
