package main

import (
	"context"

	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/input"
	"github.com/andrebq/space-invaders/ces/render"
	"github.com/andrebq/space-invaders/game"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/veandco/go-sdl2/sdl"
)

func setupWorld(win *sdl.Window, inputSys *input.System, cancel context.CancelFunc) (*ces.World, error) {

	renderSys, err := render.New(win)
	if err != nil {
		return nil, err
	}
	w := ces.NewWorld(inputSys, renderSys, game.NewArcade(cancel))
	w.AddEntity(input.OnEsc(cancel))
	colorful.FastHappyColor()

	var bounds sdl.Rect
	bounds.W, bounds.H = win.GetSize()

	w.AddEntity(game.NewWorld(bounds))
	err = game.GetWorld(w).Stage1(w)
	if err != nil {
		return nil, err
	}
	w.AddEntity(render.NewBackground(colorful.FastHappyColor()))

	return w, nil
}
