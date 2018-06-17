package render

import (
	"fmt"
	"runtime"

	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/sirupsen/logrus"

	"github.com/andrebq/space-invaders/ces"
	"github.com/andrebq/space-invaders/ces/render/ase"
)

type (
	// Sprite is an animated visual entity
	Sprite struct {
		ces.Entity

		surf     *sdl.Surface
		tex      *sdl.Texture
		Pos      sdl.Point
		animated *ase.Animated

		zOrder int
	}
)

// NewSprite creates a new animated sprite
// usually it should be used in combination with some other entity
// that will actually control where the sprite is rendered.
//
// the animation file should be the json file exported from Aseprite using
// the array format
func NewSprite(pngFile string, animationFile string, zOrder int) (*Sprite, error) {
	s := &Sprite{}

	var err error
	s.surf, err = img.LoadPNGRW(sdl.RWFromFile(pngFile, "rb"))
	if err != nil {
		return nil, errors.Wrap(err, "render:sprite: unable to load png file")
	}

	s.animated, err = ase.NewAnimated(animationFile)
	if err != nil {
		return nil, errors.Wrap(err, "render:sprite: unable to load animation file")
	}

	if err != nil {
		return nil, err
	}

	s.Pos = sdl.Point{}

	return s, err
}

func (s *Sprite) setupSurface(rgba *sdl.Surface) error {
	runtime.SetFinalizer(s, func(s *Sprite) {
		if s.surf != nil {
			s.surf.Free()
		}
	})
	return nil
}

// UpdateAnimation is used to "play" the animation
func (s *Sprite) UpdateAnimation(dt float64) {
	s.animated.Update(dt)
}

// Cycles returns how many times the animation cycle'd
func (s *Sprite) Cycles() int {
	return s.animated.Cycles()
}

func (s *Sprite) setupTexture(target *Renderer) {
	var err error
	s.tex, err = target.CreateTextureFromSurface(s.surf)
	if err != nil {
		// here there is nothing that we can do
		// we are in the middle of the game loop
		// abort or continue?!?!?!
		panic(fmt.Sprintf("render:sprite:setupTexture -> unable to create texture due to %v", err))
	}

	s.surf.Free()
	s.surf = nil

	runtime.SetFinalizer(s, func(s *Sprite) {
		if s.tex != nil {
			s.tex.Destroy()
		}
	})
}

// ZOrder is self-explanatory
func (s *Sprite) ZOrder() int {
	return s.zOrder
}

// RectAt returns the current frame rectangle at the
// given position
func (s *Sprite) RectAt(p sdl.Point) sdl.Rect {
	return point2Rect(p, ase2sdlRect(s.animated.Rect()))
}

// Paint implements renderable interface
func (s *Sprite) Paint(target *Renderer) {
	if s.tex == nil {
		s.setupTexture(target)
	}

	rect := ase2sdlRect(s.animated.Rect())
	destRect := point2Rect(s.Pos, rect)

	err := target.CopyBottom(s.tex, &rect, &destRect)
	if err != nil {
		// now, we can just log
		logrus.WithError(err).WithField("system", "render:sprite").Error("unable to render texture")
	}
}

func ase2sdlRect(r ase.Rect) sdl.Rect {
	return sdl.Rect{
		X: r.X,
		Y: r.Y,
		W: r.W,
		H: r.H,
	}
}

func point2Rect(p sdl.Point, rect sdl.Rect) sdl.Rect {
	rect.X = p.X
	rect.Y = p.Y
	return rect
}
