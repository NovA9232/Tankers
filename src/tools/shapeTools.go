package tools

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

func InCircle(circPos rl.Vector2, radiusSqr float32, pos rl.Vector2) bool {    // (x - cx)^2 + (y - cy)^2 < r^2
	return bool(math.Pow(float64(pos.X - circPos.X), 2) + math.Pow(float64(pos.Y - circPos.Y), 2) < float64(radiusSqr))
}
