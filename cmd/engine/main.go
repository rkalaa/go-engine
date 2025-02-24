package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/rkalaa/go-engine/internal/game"
)

const (
	screenHeight = 1200
	screenWidth  = 700
)

type Game struct {
	Engine game.Engine
}

func (g *Game) Update() error {

	fmt.Printf("Object 1 x: %v Object 1 y: %v\n", g.Engine.Objects[0].Position.XValue, g.Engine.Objects[0].Position.XValue)
	fmt.Printf("Object 2 x: %v Object 2 y %v\n", g.Engine.Objects[1].Position.XValue, g.Engine.Objects[1].Position.YValue)
	g.Engine.CheckCollisions() // make an if for handlecol later
	g.Engine.IncrementPosition()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, object := range g.Engine.Objects {
		vector.DrawFilledRect(screen, float32(object.Position.XValue), float32(object.Position.YValue), float32(object.Width), float32(object.Height), color.White, true)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {

	Engine := game.Engine{}
	Engine.SetWalls(screenWidth, screenHeight)
	Engine.StartOppositeObjects(screenWidth, screenHeight, 20, -2, 5, 5)
	game := &Game{Engine: Engine}
	ebiten.SetWindowSize(screenHeight, screenWidth)
	ebiten.SetWindowTitle("Go-Engine")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
