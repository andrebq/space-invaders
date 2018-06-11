package render

import (
	"image/color"

	"github.com/veandco/go-sdl2/sdl"
)

type (
	// Background is a renderable which can be used
	// and usually should be the last element to be rendered
	Background struct {
		color sdl.Color
	}
)

// NewBackground returns an empty background with the given color
func NewBackground(c color.Color) *Background {
	r, g, b, a := c.RGBA()
	return &Background{
		color: sdl.Color{
			A: uint8(a),
			G: uint8(g),
			B: uint8(b),
			R: uint8(r),
		},
	}
}

// ZOrder is always zero
func (b *Background) ZOrder() int { return 0 }

// Paint will update target with the current background
func (b *Background) Paint(target *sdl.Renderer) {
	target.SetDrawColor(
		b.color.R,
		b.color.G,
		b.color.B,
		b.color.A,
	)
	rect := target.GetClipRect()
	target.FillRect(&rect)
}
