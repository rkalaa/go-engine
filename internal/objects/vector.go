package objects

import (
	"math"
)

type Vector struct {
	XValue float64
	YValue float64
}

func (v *Vector) Distance(Vector2 Vector) float64 {
	horizonal := math.Pow(v.XValue-Vector2.XValue, 2)
	vertical := math.Pow(v.YValue-Vector2.YValue, 2)
	distance := math.Sqrt(horizonal + vertical)
	return distance
}
func (v *Vector) Add(Vector2 Vector) Vector {
	finalX := v.XValue + Vector2.XValue
	finalY := v.YValue + Vector2.YValue
	resultingVector := Vector{finalX, finalY}
	return resultingVector
}

func (v *Vector) Subtract(Vector2 Vector) Vector {
	finalX := v.XValue - Vector2.XValue
	finalY := v.YValue - Vector2.YValue
	resultingVector := Vector{finalX, finalY}
	return resultingVector
}

func (v *Vector) Scale(Scalar float64) Vector {
	finalX := v.XValue * Scalar
	finalY := v.YValue * Scalar
	resultingVector := Vector{finalX, finalY}
	return resultingVector
}

func (v *Vector) Magnitude() float64 {
	squaredX := v.XValue * v.XValue
	squaredY := v.YValue * v.YValue
	size := math.Sqrt(squaredX + squaredY)
	return size
}

func (v *Vector) DotProduct(Vector2 Vector) float64 {
	directionX := v.XValue * Vector2.XValue
	directionY := v.YValue * Vector2.YValue
	return directionX + directionY
}
