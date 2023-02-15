package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"sync"
)

type pipe struct {
	mutex                           sync.RWMutex
	xPosition, height, speed, width int32
	inverted                        bool
	renderer                        *sdl.Renderer
	texture                         *sdl.Texture
}

func newPipe(renderer *sdl.Renderer) (*pipe, error) {
	texture, err := img.LoadTexture(renderer, "res/imgs/pipe.png")
	if err != nil {
		return nil, fmt.Errorf("could load pipe texture: %v", err)
	}
	return &pipe{
		xPosition: 400,
		height:    300,
		width:     50,
		speed:     10,
		inverted:  true,
		renderer:  renderer,
		texture:   texture,
	}, nil
}

func (p *pipe) update() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.xPosition -= p.speed
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

func (p *pipe) restart() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.xPosition = 400

}
