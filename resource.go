package main

import (
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
