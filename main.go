package main

import (
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 800
	screenHeight = 600
	circleRadius = 80
)

type Game struct {
	x  float64
	y  float64
	vx float64
	vy float64
}

func (g *Game) Update() error {
	dt := 1.0 / float64(ebiten.TPS())
	g.x += g.vx * dt
	g.y += g.vy * dt

	if g.x-float64(circleRadius) <= 0 {
		g.x = float64(circleRadius)
		g.vx = math.Abs(g.vx)
	} else if g.x+float64(circleRadius) >= float64(screenWidth) {
		g.x = float64(screenWidth) - float64(circleRadius)
		g.vx = -math.Abs(g.vx)
	}

	if g.y-float64(circleRadius) <= 0 {
		g.y = float64(circleRadius)
		g.vy = math.Abs(g.vy)
	} else if g.y+float64(circleRadius) >= float64(screenHeight) {
		g.y = float64(screenHeight) - float64(circleRadius)
		g.vy = -math.Abs(g.vy)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	cx := float32(g.x)
	cy := float32(g.y)
	radius := float32(circleRadius)

	vector.DrawFilledCircle(screen, cx, cy, radius, color.RGBA{R: 220, G: 220, B: 220, A: 255}, true)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Collision Playground")

	game := &Game{
		x:  float64(screenWidth) * 0.5,
		y:  float64(screenHeight) * 0.5,
		vx: 120,
		vy: 0,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
