package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math/rand"
	"sync"
)

const (
	pipeHeight = 300
	pipeWidth  = 50
)

type pipe struct {
	mutex                           sync.RWMutex
	xPosition, height, speed, width int32
	inverted                        bool
	renderer                        *sdl.Renderer
	texture                         *sdl.Texture
}

func newPipe(renderer *sdl.Renderer, texture *sdl.Texture, xPosition int32) *pipe {
	return &pipe{
		xPosition: xPosition,
		height:    100 + int32(rand.Intn(300)),
		width:     pipeWidth,
		inverted:  rand.Float32() > 0.5,
		renderer:  renderer,
		texture:   texture,
	}
}

func (p *pipe) destroy() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.texture.Destroy()
}

func (p *pipe) createRectangle() *sdl.Rect {
	y := windowHeight - p.height
	if p.inverted {
		y = 0
	}
	return &sdl.Rect{
		X: p.xPosition,
		Y: y,
		W: p.width,
		H: p.height,
	}
}

func (p *pipe) paint() error {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	flip := sdl.FLIP_NONE
	if p.inverted {
		flip = sdl.FLIP_VERTICAL
	}
	if err := p.renderer.CopyEx(p.texture, nil, p.createRectangle(), 0, nil, flip); err != nil {
		return fmt.Errorf("could not copy pipe: %v", err)
	}
	return nil
}
