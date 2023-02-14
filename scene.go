package main

import (
	"context"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type scene struct {
	time       int
	background *background
	renderer   *sdl.Renderer
	character  *character
}

func newScene(renderer *sdl.Renderer) (*scene, error) {
	background, err := newBackground(renderer)
	if err != nil {
		return nil, err
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
	s.renderer.Clear()

	if err := s.background.paint(); err != nil {
		return err
	}

	if err := s.character.paint(); err != nil {
		return err
	}

	s.renderer.Present()
	return nil
}

func (s *scene) destroy() {
	s.background.destroy()
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
