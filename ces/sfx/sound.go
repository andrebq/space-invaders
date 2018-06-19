package sfx

import (
	"io/ioutil"
	"runtime"

	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/render"

	"github.com/sirupsen/logrus"

	"github.com/veandco/go-sdl2/mix"
)

type (
	// Effect is a simple sound effect played only once
	Effect struct {
		snd *mix.Chunk

		state  effectState
		length int
	}

	effectState byte
)

const (
	initial = effectState(iota)
	shouldPlay
	playing
	played
)

// NewEffect creates the given sound effect and prepares it to start playing
// when the first Paint call is made
func NewEffect(file string) (*Effect, error) {
	// Load entire WAV data from file
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return NewEffectFromData(data)
}

// NewEffectFromData creates a sound effect using the provided buffer as the data
func NewEffectFromData(input []byte) (*Effect, error) {
	data := make([]byte, len(input))
	copy(data, input)
	// Load WAV from data (memory)
	chunk, err := mix.QuickLoadWAV(data)
	if err != nil {
		return nil, err
	}

	effect := &Effect{
		snd: chunk,
	}
	runtime.SetFinalizer(effect, func(effect *Effect) {
		if effect.snd != nil {
			effect.snd.Free()
		}
	})
	effect.length = chunk.LengthInMs()

	return effect, nil
}

// Update checks if the sound should be played or not
func (e *Effect) Update(dt float64, w *ces.World) {
	switch e.state {
	case initial:
		e.state = shouldPlay
	case playing:
		if e.length <= 0 {
			e.state = played
			w.RemoveEntity(e)
		} else {
			e.length -= int(dt * 1000)
		}
	}
}

// ZOrder returns the 'render' order
func (e *Effect) ZOrder() int {
	return 0
}

// Paint 'renders' the sound to the 'screen'
func (e *Effect) Paint(renderer *render.Renderer) {
	switch e.state {
	case shouldPlay:
		e.state = playing
		_, err := e.snd.Play(1, 0)
		if err != nil {
			logrus.WithError(err).Error("unable to play sound effect")
		}
		return
	}
}
