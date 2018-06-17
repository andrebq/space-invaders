package render

import (
	"github.com/veandco/go-sdl2/sdl"
)

// CopyBottom works like copy but it will match the bottom-left corner of pos
// to the bottom-left of the screen,
//
// Eg.: a 32x32 rect at position (1, 1) rendered into a screen of (800, 800) will be
// rendered at (1, 800 /* screen height */ - 32 /* rect height */ - 1 /* rect Y position */)
//
// The original rect isn't changed in this operation
func (r *Renderer) CopyBottom(tex *sdl.Texture, src, dest *sdl.Rect) error {
	destCopy := *dest
	vp := r.GetViewport()

	destCopy.Y = vp.H - destCopy.H - destCopy.Y
	return r.Copy(tex, src, &destCopy)
}
