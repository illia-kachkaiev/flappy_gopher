package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"sync"
	"time"
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
	//var newPipes []*pipe
	//newPipe, _ := newPipe(renderer, texture, 400, false)
	//newPipes = append(newPipes, newPipe)
	ps := &pipes{
		texture: texture,
		speed:   25,
	}
	go addPipe(ps, renderer, texture)
	return ps, nil
}

func addPipe(ps *pipes, renderer *sdl.Renderer, texture *sdl.Texture) {
	for {
		ps.mutex.Lock()
		ps.pipes = append(ps.pipes, newPipe(renderer, texture, windowWidth-50))
		ps.mutex.Unlock()
		time.Sleep(time.Second)
	}
}

func (ps *pipes) update() {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	for _, pipe := range ps.pipes {
		pipe.mutex.Lock()
		pipe.xPosition -= ps.speed
		pipe.mutex.Unlock()
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

	ps.pipes = nil
}

func (ps *pipes) detectCollision(character *character) {
	ps.mutex.RLock()
	defer ps.mutex.RUnlock()

	character.mutex.Lock()
	defer character.mutex.Unlock()

	charRect := character.createRectangle()
	for _, pipe := range ps.pipes {
		if _, isCollide := charRect.Intersect(pipe.createRectangle()); isCollide {
			character.dead = true
		}
	}
}
