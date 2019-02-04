package bodies

import (
  "github.com/gen2brain/raylib-go/raylib"
  "tools"
)

const (
  TANK_W float32 = 30
  TANK_H float32 = 50
  TANK_ACCELL = 20
  TANK_DECEL = 10
  TANK_TURN_SPD float32 = rl.Pi
)

type Tank struct {
  BaseBody
  bodyDraw rl.Rectangle  // For drawing the tank body.
  colors []rl.Color
}

func NewTank(IDNum int, newPos rl.Vector2) *Tank {
  return &Tank{
    BaseBody: NewBody(NewID(IDNum, "tank"), newPos, 0, 0),
    bodyDraw: rl.NewRectangle(newPos.X, newPos.Y, float32(TANK_W), float32(TANK_H)),
    colors: []rl.Color{rl.Lime, rl.Lime, rl.SkyBlue, rl.Lime},
  }
}

func (self *Tank) Draw() {
  println(self.Pos.X, self.Pos.Y, self.Angle)
  rl.DrawRectanglePro(self.bodyDraw, rl.NewVector2(-self.Pos.X, -self.Pos.Y), self.Angle, self.colors)
}

func (self *Tank) Update(dt float32) {
  if rl.IsKeyDown(rl.KeyA) {
    self.Angle -= TANK_TURN_SPD*dt
  }

  self.applyResistance(dt)
  self.accelerate(dt)
  self.Pos.X += (self.Vel.X*dt)
  self.Pos.Y += (self.Vel.Y*dt)
}

func (self *Tank) accelerate(dt float32) {
  self.VelMag = tools.GetMagnitude(&self.Vel)
  self.VelMag += TANK_ACCELL*dt
  self.Vel = tools.GetXYComponent(float64(self.VelMag), float64(self.Angle))
}

func (self *Tank) applyResistance(dt float32) {
  self.VelMag -= (TANK_DECEL*dt)
  self.Vel = tools.GetXYComponent(float64(self.VelMag), float64(self.Angle))
}
