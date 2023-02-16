package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"runtime"
	"time"
)

type game struct {
	renderer *sdl.Renderer
}

func (g *game) run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("could not init SDL: %v", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("could not init ttf: %v", err)
	}

	window, err := g.createWindowAndRenderer()
	if err != nil {
		return err
	}
	defer window.Destroy()

	if err := drawTitle(g.renderer, "Flappy Gopher!"); err != nil {
		return fmt.Errorf("could not draw title: %v", err)
	}
	time.Sleep(1 * time.Second)

	scene, err := newScene(g.renderer)
	if err != nil {
		return fmt.Errorf("could not create scene: %v", err)
	}
	defer scene.destroy()

	events := make(chan sdl.Event)
	errc := scene.run(events)

	runtime.LockOSThread()
	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errc:
			return err
		}
	}
}

func drawTitle(renderer *sdl.Renderer, text string) error {
	font, err := ttf.OpenFont("res/fonts/Flappy.ttf", 20)
	if err != nil {
		return fmt.Errorf("could not load font: %v", err)
	}
	defer font.Close()

	redColor := sdl.Color{R: 255, G: 0, B: 0, A: 255}

	surface, err := font.RenderUTF8Solid(text, redColor)
	if err != nil {
		return fmt.Errorf("could not render title: %v", err)
	}
	defer surface.Free()

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return fmt.Errorf("could not create surface: %v", err)
	}
	defer texture.Destroy()

	rect := &sdl.Rect{
		X: 40,
		Y: 50,
		W: windowWidth - 80,
		H: windowHeight - 100,
	}
	if err := renderer.Copy(texture, nil, rect); err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}

	renderer.Present()

	return nil
}

func (g *game) createWindowAndRenderer() (*sdl.Window, error) {
	window, renderer, err := sdl.CreateWindowAndRenderer(windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return nil, fmt.Errorf("could not create window: %v", err)
	}
	g.renderer = renderer
	return window, nil
}
