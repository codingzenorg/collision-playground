package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	cx := float32(screenWidth) * 0.5
	cy := float32(screenHeight) * 0.5
	radius := float32(80)

	vector.DrawFilledCircle(screen, cx, cy, radius, color.RGBA{R: 220, G: 220, B: 220, A: 255}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Collision Playground")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
