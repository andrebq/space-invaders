.PHONY=build run dep

build:
	go build -x -o space-invaders-bin .

install:
	go install -x github.com/veandco/go-sdl2/sdl
	go install -x github.com/veandco/go-sdl2/img

run: build
	./space-invaders-bin

dep:
	dep ensure -v