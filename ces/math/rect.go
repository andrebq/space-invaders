package math

import "github.com/veandco/go-sdl2/sdl"

// FullyInside checks if rect a is fully contained into rect b
func FullyInside(bounds, test sdl.Rect) bool {
	tr, br, bl, tl := Corners(test)
	return Inside(bounds, tr) &&
		Inside(bounds, br) &&
		Inside(bounds, bl) &&
		Inside(bounds, tl)
}

// FullyOutside checks if rect a is fully outside of b
func FullyOutside(bounds, test sdl.Rect) bool {
	tr, br, bl, tl := Corners(test)
	return !(Inside(bounds, tr) ||
		Inside(bounds, br) ||
		Inside(bounds, bl) ||
		Inside(bounds, tl))
}

// Inside returns if the given rect contains the given point
func Inside(bounds sdl.Rect, p sdl.Point) bool {
	return p.X <= bounds.X+bounds.W &&
		p.X >= bounds.X &&
		p.Y <= bounds.Y+bounds.H &&
		p.Y >= bounds.Y
}

// Corners returns the rectangle 4 corners, followin the clock order (tr, br, bl, tl)
func Corners(r sdl.Rect) (tr, br, bl, tl sdl.Point) {
	tr = sdl.Point{
		Y: r.Y + r.H,
		X: r.X + r.W,
	}
	br = sdl.Point{
		Y: r.Y,
		X: r.X + r.W,
	}
	bl = sdl.Point{
		Y: r.Y,
		X: r.X,
	}
	tl = sdl.Point{
		Y: r.Y + r.H,
		X: r.X,
	}
	return
}

// ExpandFromCenter expands the rectangle by the given ammount but keeps the center
// if ammount is negative the rectangle is shrinked
func ExpandFromCenter(r sdl.Rect, ammount int32) sdl.Rect {
	r.X -= ammount
	r.Y -= ammount
	r.W += ammount
	r.H += ammount
	return r
}
