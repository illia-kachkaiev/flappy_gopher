package main

import (
	"context"
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type scene struct {
	time       int
	background *sdl.Texture
	renderer   *sdl.Renderer
	character  *character
}

func newScene(renderer *sdl.Renderer) (*scene, error) {
	background, err := img.LoadTexture(renderer, "res/imgs/background.png")
	if err != nil {
		return nil, fmt.Errorf("could not load background image: %v", err)
	}
	character, err := newCharacter(renderer)
	if err != nil {
		return nil, err
	}
	return &scene{
		background: background,
		renderer:   renderer,
		character:  character,
	}, nil
}

func (s *scene) paint() error {
	s.time++

	s.renderer.Clear()
	if err := s.renderer.Copy(s.background, nil, nil); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	s.character.paint()

	s.renderer.Present()
	return nil
}

func (s *scene) destroy() {
	s.background.Destroy()
	s.character.destroy()
}

func (s *scene) run(ctx context.Context) <-chan error {
	errc := make(chan error)

	go func() {
		defer close(errc)
		for range time.Tick(100 * time.Millisecond) {
			select {
			case <-ctx.Done():
				return
			default:
				if err := s.paint(); err != nil {
					errc <- err
				}
			}
		}
	}()
	return errc
}
