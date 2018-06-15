package render

import (
	"fmt"
	"runtime"

	"github.com/pkg/errors"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"

	"github.com/sirupsen/logrus"

	"github.com/andrebq/space-invaders/ces"
)

type (
	// Sprite is an animated visual entity
	Sprite struct {
		ces.Entity

		surf    *sdl.Surface
		tex     *sdl.Texture
		Pos     sdl.Rect
		srcRect sdl.Rect

		zOrder int
	}
)

// NewSprite creates a new animated (not yet) sprite
// usually it should be used in combination with some other entity
// that will actually control where the sprite is rendered
func NewSprite(pngFile string, zOrder int) (*Sprite, error) {
	s := &Sprite{}

	var err error
	s.surf, err = img.LoadPNGRW(sdl.RWFromFile(pngFile, "rb"))
	if err != nil {
		return nil, errors.Wrap(err, "render:sprite: unable to load png file")
	}
	s.srcRect = s.surf.ClipRect

	if err != nil {
		return nil, err
	}

	s.Pos = s.srcRect

	return s, err
}

func (s *Sprite) setupSurface(rgba *sdl.Surface) error {
	runtime.SetFinalizer(s, func(s *Sprite) {
		if s.surf != nil {
			s.surf.Free()
		}
	})
	// copy the original surface for later usage
	s.srcRect = s.surf.ClipRect
	return nil
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

// Paint implements renderable interface
func (s *Sprite) Paint(target *Renderer) {
	if s.tex == nil {
		s.setupTexture(target)
	}

	err := target.CopyBottom(s.tex, &s.srcRect, &s.Pos)
	if err != nil {
		// now, we can just log
		logrus.WithError(err).WithField("system", "render:sprite").Error("unable to render texture")
	}
}
