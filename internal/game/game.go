package game

import (
	"github.com/rkalaa/go-engine/internal/objects"
)

type Engine struct {
	Objects [2]objects.Object
	Walls   [4]objects.Object
}

func (e *Engine) SetWalls(screenWidth, screenHeight int) {
	floatscreenWidth, floatscreenHeight := float64(screenWidth), float64(screenHeight)
	iteratex := [4]float64{floatscreenWidth / 2, 0, floatscreenWidth, floatscreenWidth / 2}
	iteratey := [4]float64{0, floatscreenHeight / 2, floatscreenHeight / 2, floatscreenHeight}
	for i := 0; i < len(iteratex) && i < len(iteratey); i++ {
		WallPosition := objects.Vector{XValue: iteratex[i], YValue: iteratey[i]}
		WallVelocity := objects.Vector{XValue: 0, YValue: 0}
		Wall := objects.Object{Position: WallPosition, Velocity: WallVelocity, Width: int(iteratex[i]), Height: int(iteratey[i])}
		e.Walls[i] = Wall
	}

}

func (e *Engine) StartOppositeObjects(posX, posY float64, speedX, speedY float64, Width, Height int) {
	//first object assigned
	Position := objects.Vector{XValue: posX, YValue: posY}
	Velocity := objects.Vector{XValue: speedX, YValue: posY}
	Object1 := objects.Object{Position: Position, Velocity: Velocity, Width: Width, Height: Height}
	// second object given negative x values
	InvertPosition := objects.Vector{XValue: -posX, YValue: posY}
	InvertVelocity := objects.Vector{XValue: -speedX, YValue: posY}
	Object2 := objects.Object{Position: InvertPosition, Velocity: InvertVelocity, Width: Width, Height: Height}

	AllObjects := [2]objects.Object{Object1, Object2}
	e.Objects = AllObjects
}

func (e *Engine) CheckCollisions() bool {
	/* Checks if x, only then will compute y to optimize AABB methodology
	good: (1L)(1R)(2L)(2R)
	bad: (1L)(2L)(1R)(2R)
	if bad check y
	*/
	Object1, Object2 := e.Objects[0], e.Objects[1]
	Position1, Position2 := e.Objects[0].Position, e.Objects[1].Position

	Object1LX := Position1.XValue - float64(Object1.Width)
	Object1LeftVector := objects.Vector{XValue: Object1LX, YValue: 0}

	Object2LX := Position2.XValue - float64(Object2.Width)
	Object2LeftVector := objects.Vector{XValue: Object2LX, YValue: 0}

	if float64(Object1.Width) > Object1LeftVector.Distance(Object2LeftVector) {
		Object1UY := Position1.YValue + float64(Object1.Height)
		Object1UpVector := objects.Vector{XValue: 0, YValue: Object1UY}

		Object2UY := Position2.YValue + float64(Object2.Height)
		Object2UpVector := objects.Vector{XValue: 0, YValue: Object2UY}
		if float64(Object1.Height) > Object1UpVector.Distance(Object2UpVector) {
			return true
		}
	}
	return false
}

func (e *Engine) IncrementPosition() {
	(e.Objects)[0].Position = (e.Objects)[0].Position.Add((e.Objects)[0].Velocity)
	(e.Objects)[1].Position = (e.Objects)[0].Position.Add((e.Objects)[1].Velocity)

}
