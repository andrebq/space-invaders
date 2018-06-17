package input

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	onEsc struct {
		fn func()
	}
)

func (o *onEsc) KeyDown(ev *sdl.KeyboardEvent, w *ces.World) {
}

func (o *onEsc) KeyUp(ev *sdl.KeyboardEvent, w *ces.World) {
	if ev.Keysym.Sym == sdl.K_ESCAPE {
		println("leaing")
		o.fn()
	}
}

// OnEsc returns an entity which will respond when the ESC key
// is pressed-down and then up
func OnEsc(fn func()) ces.Entity {
	return &onEsc{fn}
}
