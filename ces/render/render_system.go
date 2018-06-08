package render

import (
	"github.com/andrebq/space-invaders/ces"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	renderSystem struct {
		base     ces.BaseSystem
		win      *sdl.Window
		renderer *sdl.Renderer
	}

	renderComponent interface {
		Texture() *sdl.Texture
		Rect() sdl.Rect
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
	renderer, err := win.GetRenderer()
	if err != nil {
		errors.Wrapf(err, "render: unable to get window renderer")
	}
	return &renderSystem{
		base:     b,
		win:      win,
		renderer: renderer,
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
		b.base.Watch(e, true)
	case e.Deleted:
		b.base.Watch(e, false)
	}
}

// Render prints stuff on the screen
func (b *renderSystem) Render(dt float64) {
	logrus.WithField("system", "render").Debug()
	b.base.ForEach(func(e ces.Entity) {
		logrus.WithField("entity", e).Debug()
	})
}
