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

type Game struct {
	x     float64
	y     float64
	speed float64
}

func (g *Game) Update() error {
	dt := 1.0 / float64(ebiten.TPS())
	g.x += g.speed * dt
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	cx := float32(g.x)
	cy := float32(g.y)
	radius := float32(80)

	vector.DrawFilledCircle(screen, cx, cy, radius, color.RGBA{R: 220, G: 220, B: 220, A: 255}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Collision Playground")

	game := &Game{
		x:     float64(screenWidth) * 0.5,
		y:     float64(screenHeight) * 0.5,
		speed: 120,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
