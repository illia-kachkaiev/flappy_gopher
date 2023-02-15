package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"sync"
)

type character struct {
	mutex                    sync.RWMutex
	time                     int
	textures                 []*sdl.Texture
	renderer                 *sdl.Renderer
	speed                    float64
	yPosition, width, height int32
	dead                     bool
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
func (c *character) update() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.time++
	c.yPosition -= int32(c.speed)

	if c.yPosition < 0 {
		c.dead = true
	}
	c.speed += gravity
}

func (c *character) paint() error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	rect := &sdl.Rect{X: windowWidth / 10, Y: windowHeight - c.yPosition - c.height/2, W: c.width, H: c.height}

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
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for _, texture := range c.textures {
		texture.Destroy()
	}
}

func (c *character) isDead() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.dead
}

func (c *character) restart() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.speed = 0
	c.yPosition = 300
	c.dead = false

}
