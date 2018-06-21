package sfx

import (
	"runtime"

	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/render"

	"github.com/sirupsen/logrus"

	"github.com/veandco/go-sdl2/mix"
)

type (
	// Music is a music
	Music struct {
		snd *mix.Music

		state  effectState
		length int
	}
)

// NewMusic creates the given sound effect and prepares it to start playing
// when the first Paint call is made
func NewMusic(file string) (*Music, error) {
	chunk, err := mix.LoadMUS(file)
	if err != nil {
		return nil, err
	}

	effect := &Music{
		snd: chunk,
	}
	runtime.SetFinalizer(effect, func(effect *Music) {
		if effect.snd != nil {
			effect.snd.Free()
		}
	})

	return effect, nil
}

// OnAdd handles entity add
func (e *Music) OnAdd(w *ces.World) {
}

// OnRemove
func (e *Music) OnRemove(w *ces.World) {
	mix.HaltMusic()
	e.snd.Free()
}

// Update checks if the sound should be played or not
func (e *Music) Update(dt float64, w *ces.World) {
	switch e.state {
	case initial:
		e.state = shouldPlay
	}
}

// ZOrder returns the 'render' order
func (e *Music) ZOrder() int {
	return 0
}

// Paint 'renders' the sound to the 'screen'
func (e *Music) Paint(renderer *render.Renderer) {
	switch e.state {
	case shouldPlay:
		e.state = playing
		err := e.snd.Play(-1)
		if err != nil {
			logrus.WithError(err).Error("unable to play sound effect")
		}
		return
	}
}
