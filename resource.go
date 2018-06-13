package main

import (
	"image"
	"image/png"
	"os"
	"path"
	"path/filepath"
)

// will panic if resource isn't found
func findResource(name string) string {
	absPath, err := filepath.Abs(
		filepath.FromSlash(path.Join(".", "assets", name)))
	if err != nil {
		panic(err)
	}
	return absPath
}

// findPNG will look for the given file and process it as png
// it will panic if not find
func findPNG(name string) image.Image {
	abs := findResource(name)
	file, err := os.Open(abs)
	if err != nil {
		panic("assets: " + err.Error())
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic("assets:png: " + err.Error())
	}
	return img
}
