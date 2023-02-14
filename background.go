package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type background struct {
	texture  *sdl.Texture
	renderer *sdl.Renderer
}

func newBackground(renderer *sdl.Renderer) (*background, error) {
	texture, err := img.LoadTexture(renderer, "res/imgs/background.png")
	if err != nil {
		return nil, fmt.Errorf("could not load background image: %v", err)
	}
	return &background{texture: texture, renderer: renderer}, nil
}

func (b *background) paint() error {
	if err := b.renderer.Copy(b.texture, nil, nil); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}
	return nil
}

func (b *background) destroy() {
	b.texture.Destroy()
}
