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
)

type Circle struct {
	x      float64
	y      float64
	vx     float64
	vy     float64
	radius float64
	color  color.RGBA
}

type Game struct {
	circles []Circle
}

func (g *Game) Update() error {
	dt := 1.0 / float64(ebiten.TPS())
	for i := range g.circles {
		c := &g.circles[i]
		c.x += c.vx * dt
		c.y += c.vy * dt

		if c.x-c.radius <= 0 {
			c.x = c.radius
			c.vx = math.Abs(c.vx)
		} else if c.x+c.radius >= float64(screenWidth) {
			c.x = float64(screenWidth) - c.radius
			c.vx = -math.Abs(c.vx)
		}

		if c.y-c.radius <= 0 {
			c.y = c.radius
			c.vy = math.Abs(c.vy)
		} else if c.y+c.radius >= float64(screenHeight) {
			c.y = float64(screenHeight) - c.radius
			c.vy = -math.Abs(c.vy)
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, c := range g.circles {
		vector.DrawFilledCircle(
			screen,
			float32(c.x),
			float32(c.y),
			float32(c.radius),
			c.color,
			true,
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Collision Playground")

	game := &Game{
		circles: []Circle{
			{x: 120, y: 120, vx: 90, vy: 60, radius: 28, color: color.RGBA{R: 230, G: 96, B: 96, A: 255}},
			{x: 280, y: 140, vx: -110, vy: 40, radius: 36, color: color.RGBA{R: 96, G: 200, B: 220, A: 255}},
			{x: 420, y: 200, vx: 70, vy: -90, radius: 22, color: color.RGBA{R: 120, G: 220, B: 120, A: 255}},
			{x: 640, y: 160, vx: -130, vy: 55, radius: 30, color: color.RGBA{R: 220, G: 180, B: 90, A: 255}},
			{x: 180, y: 320, vx: 80, vy: -70, radius: 40, color: color.RGBA{R: 190, G: 110, B: 210, A: 255}},
			{x: 360, y: 360, vx: 140, vy: 30, radius: 24, color: color.RGBA{R: 90, G: 160, B: 240, A: 255}},
			{x: 520, y: 340, vx: -100, vy: -60, radius: 34, color: color.RGBA{R: 230, G: 140, B: 170, A: 255}},
			{x: 700, y: 420, vx: -85, vy: 95, radius: 26, color: color.RGBA{R: 140, G: 230, B: 160, A: 255}},
			{x: 260, y: 480, vx: 60, vy: -110, radius: 32, color: color.RGBA{R: 210, G: 210, B: 120, A: 255}},
			{x: 560, y: 500, vx: 115, vy: 75, radius: 38, color: color.RGBA{R: 120, G: 140, B: 230, A: 255}},
		},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
