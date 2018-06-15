package main

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/input"
	"github.com/andrebq/space-invaders/ces/render"
	"github.com/andrebq/space-invaders/game"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/veandco/go-sdl2/sdl"
)

func setupWorld(win *sdl.Window, input *input.System) (*ces.World, error) {

	renderSys, err := render.New(win)
	if err != nil {
		return nil, err
	}
	w := ces.NewWorld(input, renderSys, game.NewArcade())
	colorful.FastHappyColor()

	var bounds sdl.Rect
	bounds.W, bounds.H = win.GetSize()

	w.AddEntity(game.NewWorld(bounds))
	w.AddEntity(render.NewBackground(colorful.FastHappyColor()))

	player, err := game.NewPlayer(findResource("player.png"))
	if err != nil {
		return nil, err
	}
	w.AddEntity(player)

	return w, nil
}
