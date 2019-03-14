package entities

import (
  "github.com/gen2brain/raylib-go/raylib"
)

// Base class for bodies.
type BaseBody struct {
  Id *ID
  Pos rl.Vector2        // Position to middle of body
  VelMag float32
  Angle float32
}

func NewBody(newId *ID, newPos rl.Vector2, newVelMag, newAngle float32) BaseBody {
  return BaseBody{
    Id: newId,
    Pos: newPos,
    VelMag: newVelMag,
    Angle: newAngle,
  }
}
