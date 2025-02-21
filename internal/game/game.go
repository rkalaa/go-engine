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

func CheckCollisions(AllObjects [2]objects.Object) bool {
	/* Checks if x, only then will compute y to optimize
	good: (1L)(1R)(2L)(2R)
	bad: (1L)(2L)(1R)(2R)
	if bad check y
	*/
	Object1, Object2 := AllObjects[0], AllObjects[1]
	Position1, Position2 := AllObjects[0].Position, AllObjects[1].Position

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

func IncrementPosition(AllObjects *[2]objects.Object) {
	(*AllObjects)[0].Position = (*AllObjects)[0].Position.Add((*AllObjects)[0].Velocity)
	(*AllObjects)[1].Position = (*AllObjects)[0].Position.Add((*AllObjects)[1].Velocity)

}
