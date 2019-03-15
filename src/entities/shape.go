package entities

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

type Shape interface {
	GetArea() float32
	IsInside(float32, float32) bool
}

type Rectangle struct {
	rl.Rectangle
}

func (r *Rectangle) GetArea() float32 {
	return r.X * r.Y
}

func (r *Rectangle) IsInside(x, y float32) bool {  // x and y are relative to origin (top left) of shape
	return bool(x > 0 && x < r.Width && y > 0 && y < r.Height)

type Circle struct {
	Radius float32
}

func (c *Circle) GetArea() float32 {
	return rl.Pi * float32(math.Pow(float64(c.Radius), 2))
}

func (c *Circle) IsInside(x, y float32) bool {  // x^2 + y^2 = r^2

}
