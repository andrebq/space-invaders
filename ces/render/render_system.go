package render

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	renderSystem struct {
		base     ces.BaseSystem
		win      *sdl.Window
		renderer *Renderer

		size sdl.Rect

		layers layers
	}

	layers map[int]*renderList

	renderComponent interface {
		ces.Entity
		Paint(*Renderer)
		ZOrder() int
	}

	// Renderer adds some helper function on top of sdl default renderer api
	Renderer struct {
		*sdl.Renderer
	}

	isRenderable struct{}
)

// Filter returns true if the given entity is renderable
func (isRenderable) Filter(e ces.Entity) bool {
	_, ok := e.(renderComponent)
	return ok
}

// New returns a new render system
func New(win *sdl.Window) (ces.RenderSystem, error) {
	b := ces.BaseSystem{}

	renderer, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		errors.Wrapf(err, "render: unable to get window renderer")
	}
	return &renderSystem{
		base:     b,
		win:      win,
		renderer: &Renderer{renderer},
		layers:   make(layers),
	}, nil
}

// EntityFilter implements ces.System
func (b *renderSystem) EntityFilter() ces.EntityFilter {
	return isRenderable{}
}

// EntityLifecycle implements ces.System
func (b *renderSystem) EntityLifecycle(e ces.EntityLifecycleEvent) {
	switch {
	case e.Created:
		b.add(e.Entity.(renderComponent))
	case e.Deleted:
		b.remove(e.Entity.(renderComponent))
	}
}

func (b *renderSystem) add(e renderComponent) {
	e = e.(renderComponent)

	list := b.layers[e.ZOrder()]
	if list == nil {
		list = new(renderList)
		b.layers[e.ZOrder()] = list
	}
	list.add(e)
}

func (b *renderSystem) remove(e renderComponent) {
	e = e.(renderComponent)

	list := b.layers[e.ZOrder()]
	if list == nil {
		list.remove(e)
	}
}

// Render prints stuff on the screen
func (b *renderSystem) Render(dt float64) {

	b.renderer.Clear()

	for i := 0; i < 100; i++ {
		l := b.layers[i]
		if l != nil {
			for _, renderable := range *l {
				renderable.Paint(b.renderer)
			}
		}
	}

	b.renderer.Present()
}
