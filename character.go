package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type character struct {
	time             int
	textures         []*sdl.Texture
	renderer         *sdl.Renderer
	speed, yPosition float64
	width, height    int32
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
	return &character{textures: frames, renderer: renderer, yPosition: windowHeight / 2, width: 50, height: 43}, nil
}

func (c *character) paint() error {
	c.time++
	c.yPosition -= c.speed
	characterHalfHeight := c.height / 2

	if c.yPosition < 0 {
		c.speed = -c.speed
	}
	c.speed += gravity

	rect := &sdl.Rect{X: windowWidth / 5, Y: windowHeight - int32(c.yPosition) - characterHalfHeight, W: c.width, H: c.height}

	frameIdx := c.time % len(c.textures)
	if err := c.renderer.Copy(c.textures[frameIdx], nil, rect); err != nil {
		return fmt.Errorf("could not copy character: %v", err)
	}
	return nil
}

func (c *character) jump() {
	c.speed = -25
}

func (c *character) destroy() {
	for _, texture := range c.textures {
		texture.Destroy()
	}
}
