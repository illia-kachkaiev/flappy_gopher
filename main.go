package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(2)
	}
}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("could not init SDL: %v", err)
	}
	defer sdl.Quit()

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("could not init ttf: %v", err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer(800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("could not create window: %v", err)
	}
	defer window.Destroy()

	if err := drawTitle(renderer); err != nil {
		return fmt.Errorf("could not draw title: %v", err)
	}
	time.Sleep(2 * time.Second)

	if err := drawBackground(renderer); err != nil {
		return fmt.Errorf("could not draw background: %v", err)
	}
	time.Sleep(2 * time.Second)

	return nil
}

func drawBackground(renderer *sdl.Renderer) error {
	renderer.Clear()

	texture, err := img.LoadTexture(renderer, "res/imgs/background.png")
	if err != nil {
		return fmt.Errorf("could not load background image: %v", err)
	}
	if err := renderer.Copy(texture, nil, nil); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	renderer.Present()

	return nil
}

func drawTitle(renderer *sdl.Renderer) error {
	font, err := ttf.OpenFont("res/fonts/Flappy.ttf", 20)
	if err != nil {
		return fmt.Errorf("could not load font: %v", err)
	}
	defer font.Close()

	surface, err := font.RenderUTF8Solid("Flappy Gopher", sdl.Color{R: 255, G: 0, B: 0, A: 255})
	if err != nil {
		return fmt.Errorf("could not render title: %v", err)
	}
	defer surface.Free()

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return fmt.Errorf("could not create surface: %v", err)
	}
	defer texture.Destroy()

	if err := renderer.Copy(texture, nil, nil); err != nil {
		return fmt.Errorf("could not copy texture: %v", err)
	}

	renderer.Present()

	return nil
}
