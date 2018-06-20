package main

import (
	"context"
	"time"

	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/input"
	"github.com/veandco/go-sdl2/sdl"
)

func loopForever(ctx context.Context, win *sdl.Window, world *ces.World, input *input.System) error {
	ticker := time.NewTicker(time.Millisecond * 15)
	defer ticker.Stop()
	oldTime := time.Now()
	for {
		select {
		case <-ticker.C:
			oldTime = world.Iterate(oldTime)
			if input.ShouldQuit() {
				return ctx.Err()
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
