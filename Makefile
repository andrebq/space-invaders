.PHONY=build

build:
	go build -o space-invaders-bin .

run: build
	./space-invaders-bin