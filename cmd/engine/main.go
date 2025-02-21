package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenHeight = 1200
	screenWidth  = 700
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {

	game := &Game{}
	ebiten.SetWindowSize(screenHeight, screenWidth)
	ebiten.SetWindowTitle("Engine")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
