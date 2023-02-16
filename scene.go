package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"time"
)

type scene struct {
	time       int
	background *background
	renderer   *sdl.Renderer
	character  *character
	pipes      *pipes
}

func (s *scene) restart() {
	s.character.restart()
	s.pipes.restart()
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
	pipes, err := newPipes(renderer)
	if err != nil {
		return nil, err
	}

	return &scene{
		background: background,
		renderer:   renderer,
		character:  character,
		pipes:      pipes,
	}, nil
}

func (s *scene) update() {
	s.pipes.detectCollision(s.character)
	s.character.update()
	s.pipes.update()
}

func (s *scene) paint() error {
	s.renderer.Clear()

	if err := s.background.paint(); err != nil {
		return err
	}

	if err := s.character.paint(); err != nil {
		return err
	}
	if err := s.pipes.paint(); err != nil {
		return err
	}

	s.renderer.Present()
	return nil
}

func (s *scene) destroy() {
	s.background.destroy()
	s.character.destroy()
	s.pipes.destroy()
}

func (s *scene) advance(errc chan error) {
	s.update()
	if s.character.isDead() {
		drawTitle(s.renderer, "Game Over")
		time.Sleep(time.Second)
		s.restart()
	}
	if err := s.paint(); err != nil {
		errc <- err
	}
}

func (s *scene) run(events <-chan sdl.Event) <-chan error {
	errc := make(chan error)

	go func() {
		defer close(errc)
		tick := time.Tick(100 * time.Millisecond)
		for {
			select {
			case event := <-events:
				if isDone := s.handleEvent(event); isDone {
					drawTitle(s.renderer, "Good Buy")
					time.Sleep(time.Second)
					return
				}
			case <-tick:
				s.advance(errc)
			}
		}
	}()
	return errc
}

func (s *scene) handleEvent(event sdl.Event) bool {
	switch event.(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.MouseButtonEvent:
		s.character.jump()
	case *sdl.MouseMotionEvent, *sdl.WindowEvent, *sdl.TouchFingerEvent:
	default:
		log.Printf("unknown even %T", event)
	}
	return false
}
