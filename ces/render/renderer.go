package render

import (
	"github.com/veandco/go-sdl2/sdl"
)

// ToBottom takes the given input rectangle and transforms it using
// the current viewport to match the bottom-left corner of the rectangle
// to the bottom-left corner of the viewport.
//
// Usually (0,0) maps to the top-left corner of the output, and most of the time
// we develop considering (0,0) to be the bottom-right.
//
// Eg.: a 32x32 rect at position (1, 1) rendered into a screen of (800, 800) will be
// rendered at (1, 800 /* screen height */ - 32 /* rect height */ - 1 /* rect Y position */)
func (r *Renderer) ToBottom(rect sdl.Rect) *sdl.Rect {
	vp := r.GetViewport()

	rect.Y = vp.H - rect.H - rect.Y
	return &rect
}
