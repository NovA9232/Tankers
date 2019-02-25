package bodies

import (
  "github.com/gen2brain/raylib-go/raylib"
)

var (
	SCREEN_W float32
	SCREEN_H float32
)

// Base class for bodies.
type BaseBody struct {
  Id *ID
  Pos rl.Vector2        // Position to middle of tank
  VelMag float32
  Angle float32
}

type Body interface { // Interface for body
  Draw()
  Update()
}

func NewBody(newId *ID, newPos rl.Vector2, newVelMag, newAngle float32) BaseBody {
  return BaseBody{
    Id: newId,
    Pos: newPos,
    VelMag: newVelMag,
    Angle: newAngle,
  }
}
