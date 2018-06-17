package ase

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

type (
	// Animated loads an animation definition
	// and controls the animation by selecting
	// which source frame should be used at any given time
	Animated struct {
		// Animation is the source of the animated object
		Animation *Animation

		// current holds the frame index from the animation
		current int

		// for how long the current frame was visible
		elapsed float64

		cycles int
	}
)

// NewAnimated returns a new Animated object
func NewAnimated(animationConfig string) (*Animated, error) {
	buf, err := ioutil.ReadFile(animationConfig)
	if err != nil {
		return nil, errors.Wrap(err, "render:ase:animated unable to load animation config")
	}

	var out Animated
	out.Animation = &Animation{}
	err = json.Unmarshal(buf, out.Animation)
	if err != nil {
		return nil, errors.Wrap(err, "render:ase:animated unable to decode json")
	}
	return &out, nil
}

// Update is used to inform the animation that some time has elapsed since the last
// frame was rendered
func (a *Animated) Update(dt float64) {
	a.elapsed += dt * 1000
	cf := a.Animation.Frames[a.current]
	if a.elapsed >= cf.Duration {
		// we are past the last frame
		a.nextFrame()
		// instead of simply zero'ing
		// include the extra time
		a.elapsed = a.elapsed - cf.Duration
	}
}

// Cycles returns how many times the animation cycled
func (a *Animated) Cycles() int {
	return a.cycles
}

// Rect returns the current rect from the input
func (a *Animated) Rect() Rect {
	return a.Animation.Frames[a.current].Frame
}

func (a *Animated) nextFrame() {
	a.current++
	if a.current >= len(a.Animation.Frames) {
		a.current = 0
		a.cycles++
	}
}
