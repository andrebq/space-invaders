package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"

	"github.com/andrebq/space-invaders/ces/input"
	"github.com/sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	//logrus.SetLevel(logrus.DebugLevel)
	initSDL()

	defer quitSDL()

	win := createWindow()

	ctx, cancel := withSigCatch(context.Background())
	input := input.Get()
	world, err := setupWorld(win, input, cancel)
	if err != nil {
		logrus.WithError(err).Error("error during setup")
	}

	err = loopForever(ctx, win, world, input)
	if err != nil {
		logrus.WithError(err).Error("loopForever err'ed out")
	}
}

func createWindow() *sdl.Window {
	win, err := sdl.CreateWindow("Go Invade some Spaces!", sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED, 800, 400, sdl.WINDOW_BORDERLESS)
	if err != nil {
		logrus.WithError(err).Error("unable to create window")
		// if we don't have a window, there is no need to
		// render anything...
		panic(err)
	}
	return win
}

func quitSDL() {
	sdl.Quit()
}

func initSDL() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		logrus.WithError(err).Error("unable to start SDL")
		// no need to try anything else since
		// we couldn't start SDL
		// and this shouldn't happen actually....
		panic(err.Error())
	}
}

func withSigCatch(ctx context.Context) (context.Context, context.CancelFunc) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		s := <-sig
		logrus.WithField("signal", fmt.Sprintf("%v", s)).Print("signal caught!")
		signal.Stop(sig)

		cancel()
	}()

	return ctx, cancel
}
