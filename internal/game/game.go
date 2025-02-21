package game

import (
	"github.com/rkalaa/go-engine/internal/objects"
)

func StartOppositeObjects(posX, posY float64, speedX, speedY float64, Width, Height int) [2]objects.Object {
	//first object assigned
	Position := objects.Vector{XValue: posX, YValue: posY}
	Velocity := objects.Vector{XValue: speedX, YValue: posY}
	Object1 := objects.Object{Position: Position, Velocity: Velocity, Width: Width, Height: Height}
	// second object given negative x values
	InvertPosition := objects.Vector{XValue: -posX, YValue: posY}
	InvertVelocity := objects.Vector{XValue: -speedX, YValue: posY}
	Object2 := objects.Object{Position: InvertPosition, Velocity: InvertVelocity, Width: Width, Height: Height}

	AllObjects := [2]objects.Object{Object1, Object2}
	return AllObjects
}

func CheckCollisions(AllObjects [2]objects.Object) {
	/* Checks if x, only then will compute y to optimize
	good: (1L)(1R)(2L)(2R)
	bad: (1L)(2L)(1R)(2R)
	if bad check y
	*/
	Object1, Object2 := AllObjects[0], AllObjects[1]
	Position1, Position2 := AllObjects[0].Position, AllObjects[1].Position
	Object1LX, Object1RX := Position1.XValue-float64(Object1.Width), Position1.XValue+float64(Object1.Width)
	Object2LX, Object2RX := Position2.XValue-float64(Object2.Width), Position2.XValue+float64(Object2.Width)

}

func IncrementPosition()
