package sfx

import "github.com/veandco/go-sdl2/mix"

// InitAudio will setup audio and return another function to close it
// if no error is found
func InitAudio() (func(), error) {
	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		return nil, err
	}
	return func() {
		mix.CloseAudio()
	}, nil
}
