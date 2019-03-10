package entities

import (
  "github.com/gen2brain/raylib-go/raylib"
  "tools"
)

// Base class for bodies.
type BaseBody struct {
  Id *ID
  Pos rl.Vector2        // Position to middle of body
	W float32
	H float32
	Vel rl.Vector2    // Use velmag only if you want directional travel
  VelMag float32
  Angle float32
}

func (b *BaseBody) Update(dt float32) {
  b.Pos.X += (b.Vel.X*dt)
  b.Pos.Y -= (b.Vel.Y*dt)
}

func NewBody(newId *ID, newPos rl.Vector2, w, h float32, newVelMag, newAngle float32) BaseBody {
  return BaseBody{
    Id: newId,
    Pos: newPos,
		W: w,
		H: h,
		Vel: tools.GetXYComponent(newVelMag, newAngle),
    VelMag: newVelMag,
    Angle: newAngle,
  }
}
