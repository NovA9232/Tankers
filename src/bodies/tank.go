package bodies

import (
  "github.com/gen2brain/raylib-go/raylib"
  "tools"
)

const (
  TANK_W float32 = 30
  TANK_H float32 = 50
  TANK_ACCELL = 50
  TANK_DECEL = 0.2
  TANK_TURN_SPD float32 = rl.Pi
)

var (
	halfTankW = TANK_W/2
	halfTankH = TANK_H/2

	tankBody rl.Texture2D
	tankFrame rl.Rectangle = rl.NewRectangle(0, 0, TANK_W, TANK_H)
)

type Tank struct {
  BaseBody
  colors []rl.Color
}

func NewTank(IDNum int, newPos rl.Vector2) *Tank {
	if tankBody.ID == uint32(0) {
		println("Loading tankBody.png texture.")
		tankBody = rl.LoadTexture("src/assets/exports/tankBody.png")
	}

  t := &Tank{
    BaseBody: NewBody(NewID(IDNum, "tank"), newPos, 0, 0),
    colors: []rl.Color{rl.Lime, rl.Lime, rl.SkyBlue, rl.Lime},
  }
	return t
}

func (self *Tank) Draw() {
	rl.DrawTexturePro(tankBody, tankFrame, rl.NewRectangle(self.Pos.X, self.Pos.Y, TANK_W, TANK_H), rl.NewVector2(halfTankW, halfTankH), self.Angle*rl.Rad2deg, rl.White)
}

func (self *Tank) Update(dt float32) {
  if rl.IsKeyDown(rl.KeyA) {
    self.Angle -= TANK_TURN_SPD*dt
		self.Vel = self.getNewVelAtAngle()
  }
	if rl.IsKeyDown(rl.KeyD) {
    self.Angle += TANK_TURN_SPD*dt
		self.Vel = self.getNewVelAtAngle()
	}

  self.accelerate(dt)
  self.applyResistance(dt)

  self.Pos.X += (self.Vel.X*dt)
  self.Pos.Y -= (self.Vel.Y*dt)
}

func (self *Tank) getNewVelAtAngle() rl.Vector2 {   // Call after changing angle.
	mag := tools.GetMagnitude(&self.Vel)
	return tools.GetXYComponent(float64(mag), float64(self.Angle))
}

func (self *Tank) accelerate(dt float32) {
  self.VelMag = tools.GetMagnitude(&self.Vel)
  self.VelMag += TANK_ACCELL*dt
  self.Vel = tools.GetXYComponent(float64(self.VelMag), float64(self.Angle))
}

func (self *Tank) applyResistance(dt float32) {
  self.VelMag -= (self.VelMag*TANK_DECEL*dt)
  self.Vel = tools.GetXYComponent(float64(self.VelMag), float64(self.Angle))
}
