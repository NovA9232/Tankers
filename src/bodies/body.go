package bodies

import (
  "github.com/gen2brain/raylib-go/raylib"
  "tools"
)

// Base class for bodies.
type BaseBody struct {
  Id *ID
  Pos rl.Vector2        // Position to middle of tank
  Vel rl.Vector2
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
    Vel: tools.GetXYComponent(float64(newVelMag), float64(newAngle)),
    VelMag: newVelMag,
    Angle: newAngle,
  }
}
