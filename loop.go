package main

import (
	"context"
	"fmt"
	"time"

	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/input"
	"github.com/sirupsen/logrus"
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

func processEvents(win *sdl.Window) bool {
	for ev := sdl.WaitEvent(); ev != nil; ev = sdl.PollEvent() {
		switch ev := ev.(type) {
		case *sdl.QuitEvent:
			return true
		case *sdl.WindowEvent:
			if ev.Type == sdl.WINDOWEVENT_CLOSE &&
				sameWindow(win, ev.WindowID) {
				// since we only have one window
				// if it is closed
				// we can quit
				return true
			}
		default:
			logrus.WithField("event", fmt.Sprintf("%#v", ev)).Print()
		}
	}
	return false
}

func sameWindow(win *sdl.Window, evWindowID uint32) bool {
	winID, err := win.GetID()
	if err != nil {
		logrus.WithError(err).Error("unable to obtain ID from window")
		return false
	}
	return winID == evWindowID
}
