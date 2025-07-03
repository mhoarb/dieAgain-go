package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

var (
	screenWidth  int
	screenHeight int
)

const (
	groundY   = 400
	gravity   = 0.1
	jumpPower = -10
)

type Game struct {
	playerX   float64
	playerY   float64
	velocityY float64
	onGround  bool
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.playerX += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.playerX -= 2
	}

	if ebiten.IsKeyPressed(ebiten.KeySpace) && g.onGround {
		g.velocityY = jumpPower
		g.onGround = false
	}

	g.velocityY += gravity
	g.playerY += g.velocityY

	if g.playerY > groundY {
		g.playerY = groundY
		g.velocityY = 0
		g.onGround = true
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 135, G: 206, B: 235, A: 255})

	ebitenutil.DrawRect(screen, 0, groundY+20, float64(screenWidth), 60, color.RGBA{R: 100, G: 200, B: 100, A: 255})

	playerSize := 20.0

	ebitenutil.DrawRect(screen, g.playerX, g.playerY, playerSize, playerSize, color.RGBA{R: 255, A: 255})

	ebitenutil.DebugPrint(screen, "Left/Right to move, Space to jump")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func init() {
	w, h := ebiten.ScreenSizeInFullscreen()
	screenWidth = w
	screenHeight = h
}
func main() {
	game := &Game{
		playerX: 100,
		playerY: groundY,
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Die Again - Gravity & Jump")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
