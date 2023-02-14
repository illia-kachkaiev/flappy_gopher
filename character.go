package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type character struct {
	time     int
	textures []*sdl.Texture
	renderer *sdl.Renderer
}

func newCharacter(renderer *sdl.Renderer) (*character, error) {
	var frames []*sdl.Texture
	for i := 1; i <= 4; i++ {
		path := fmt.Sprintf("res/imgs/bird_frame_%d.png", i)
		frame, err := img.LoadTexture(renderer, path)
		frames = append(frames, frame)
		if err != nil {
			return nil, fmt.Errorf("could not load character frame image: %v", err)
		}
	}
	return &character{textures: frames, renderer: renderer}, nil
}
