package main

import (
	"context"
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type scene struct {
	time            int
	background      *sdl.Texture
	renderer        *sdl.Renderer
	characterFrames []*sdl.Texture
}

func newScene(renderer *sdl.Renderer) (*scene, error) {
	background, err := img.LoadTexture(renderer, "res/imgs/background.png")
	if err != nil {
		return nil, fmt.Errorf("could not load background image: %v", err)
	}
	var frames []*sdl.Texture
	for i := 1; i <= 4; i++ {
		path := fmt.Sprintf("res/imgs/bird_frame_%d.png", i)
		frame, err := img.LoadTexture(renderer, path)
		frames = append(frames, frame)
		if err != nil {
			return nil, fmt.Errorf("could not load character frame image: %v", err)
		}
	}
	return &scene{background: background, renderer: renderer, characterFrames: frames}, nil
}

func (s *scene) paint() error {
	s.time++

	s.renderer.Clear()
	if err := s.renderer.Copy(s.background, nil, nil); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	frameIdx := s.time % len(s.characterFrames)
	rect := &sdl.Rect{X: 100, Y: 130, W: 50, H: 43}
	if err := s.renderer.Copy(s.characterFrames[frameIdx], nil, rect); err != nil {
		return fmt.Errorf("could not copy character: %v", err)
	}

	s.renderer.Present()
	return nil
}

func (s *scene) destroy() {
	s.background.Destroy()
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
