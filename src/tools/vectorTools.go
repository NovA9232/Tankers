package tools

import (
  "math"
  "github.com/gen2brain/raylib-go/raylib"
)

func GetMagnitude(vec *rl.Vector2) float32 {
  return float32(math.Sqrt(math.Pow(float64(vec.X), 2) + math.Pow(float64(vec.Y), 2)))
}

func GetXYComponent(magnitude, angle float64) rl.Vector2 {
  return rl.NewVector2(float32(magnitude*math.Sin(angle)), float32(magnitude*math.Cos(angle)))
}

func GetAngle(vec rl.Vector2) float32 {
  return float32(math.Atan2(float64(vec.Y), float64(vec.X)))
}

func SubVec(v1, v2 rl.Vector2) rl.Vector2 {    // AB = OB - OA
	return rl.NewVector2(v2.X - v1.X, v2.Y - v1.Y)
}

func RotateVec(v1, v2 *rl.Vector2, B float32) {  // Edits v1 directly
	var (
		b float64 = float64(B)
		temp rl.Vector2 = rl.NewVector2(0, 0)
	)
	v := SubVec(*v2, *v1)
	temp.X = (v.X * float32(math.Cos(b))) - (v.Y * float32(math.Sin(b)))
	temp.Y = (v.X * float32(math.Sin(b))) + (v.Y * float32(math.Cos(b)))
	v1.X += temp.X
	v1.Y += temp.Y
}
