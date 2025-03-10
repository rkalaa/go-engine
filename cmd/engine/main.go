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
	screenWidth  int     = 1000
	screenHeight int     = 1000
	logicSizeX   float32 = float32(screenWidth) / 100
	logicSizeY   float32 = float32(screenHeight) / 100
)

type Game struct {
	Engine game.Engine
}

func (g *Game) Update() error {

	fmt.Printf("Object 1 x: %v Object 1 y: %v\n", g.Engine.Objects[0].Position.XValue, g.Engine.Objects[0].Position.XValue)
	fmt.Printf("Object 2 x: %v Object 2 y %v\n", g.Engine.Objects[1].Position.XValue, g.Engine.Objects[1].Position.YValue)
	if g.Engine.CheckCollisions() {
		if g.Engine.HandleCollision() != nil {
			println()
		}
	} // make an if for handlecol later
	g.Engine.IncrementPosition()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, object := range g.Engine.Objects {
		vector.DrawFilledRect(screen, float32(object.Position.XValue), float32(object.Position.YValue), float32(object.Width), float32(object.Height), color.White, true)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 100, 100
}

func main() {

	Engine := game.Engine{}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	Engine.SetWalls(screenWidth, screenHeight)
	Engine.StartOppositeObjects(100, 100, 3, -2, 10, 10)
	game := &Game{Engine: Engine}
	ebiten.SetTPS(10)
	ebiten.SetWindowTitle("Go-Engine")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
