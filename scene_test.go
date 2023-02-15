package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"testing"
)

func TestNewSceneCreated(t *testing.T) {
	window, renderer, _ := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	defer window.Destroy()

	scene, _ := newScene(renderer)
	if scene.background == nil {
		t.Fatalf("There's no background created for scene.")
	}
	if scene.character == nil {
		t.Fatalf("There's no character created for scene.")
	}
	if scene.renderer != renderer {
		t.Fatalf("Rendere is no assign to scene struct.")
	}
}
