package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"sync"
)

type pipes struct {
	mutex sync.RWMutex

	texture *sdl.Texture
	speed   int32

	pipes []*pipe
}

func newPipes(renderer *sdl.Renderer) (*pipes, error) {
	texture, err := img.LoadTexture(renderer, "res/imgs/pipe.png")
	if err != nil {
		return nil, fmt.Errorf("could load pipe texture: %v", err)
	}
	return &pipes{
		speed:   10,
		texture: texture,
	}, nil
}

func (ps *pipes) update() {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	for _, pipe := range ps.pipes {
		pipe.update()
	}
}

func (ps *pipes) destroy() {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	ps.texture.Destroy()
}

func (ps *pipes) paint() error {
	ps.mutex.RLock()
	defer ps.mutex.RUnlock()

	for _, pipe := range ps.pipes {
		if err := pipe.paint(); err != nil {
			return err
		}
	}
	return nil
}

func (ps *pipes) restart() {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	for _, pipe := range ps.pipes {
		pipe.restart()
	}
}
