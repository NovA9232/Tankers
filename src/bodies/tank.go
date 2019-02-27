package bodies

import (
  "github.com/gen2brain/raylib-go/raylib"
  "tools"
	"math"
)

const (
  TANK_W float32 = 30
  TANK_H float32 = 50
  TANK_ACCELL = 200
  TANK_DECEL = 0.99
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
	Cannon *tankCannon
  colors []rl.Color
}

func NewTank(IDNum int, newPos rl.Vector2) *Tank {
	if tankBody.ID == uint32(0) {
		println("Loading tankBody.png texture.")
		tankBody = rl.LoadTexture("src/assets/exports/tankBody.png")
	}

  t := &Tank{
    BaseBody: NewBody(NewID(IDNum, "tank"), newPos, 0, 0),
		Cannon: nil
    colors: []rl.Color{rl.Lime, rl.Lime, rl.SkyBlue, rl.Lime},
  }
	t.newCannon()

	return t
}

func (t *Tank) Draw() {
	rl.DrawTexturePro(tankBody, tankFrame, rl.NewRectangle(t.Pos.X, t.Pos.Y, TANK_W, TANK_H), rl.NewVector2(halfTankW, halfTankH), t.Angle*rl.Rad2deg, rl.White)
}

func (t *Tank) Update(dt float32) {
  if rl.IsKeyDown(rl.KeyA) {
    t.Angle -= TANK_TURN_SPD * dt
  }
	if rl.IsKeyDown(rl.KeyD) {
    t.Angle += TANK_TURN_SPD * dt
	}
	if rl.IsKeyDown(rl.KeyW) {
		t.accelerate(dt, 1)
	}
	if rl.IsKeyDown(rl.KeyS) {
		t.accelerate(dt, -1)
	}

  t.applyResistance(dt)

	vel := tools.GetXYComponent(float64(t.VelMag), float64(t.Angle))
  t.Pos.X += (vel.X*dt)
  t.Pos.Y -= (vel.Y*dt)
}

func (t *Tank) accelerate(dt, totalPower float32) {
  t.VelMag += (TANK_ACCELL * dt * totalPower)
}

func (t *Tank) applyResistance(dt float32) {
  t.VelMag = t.VelMag * (TANK_DECEL * float32(math.Min(144/float64(dt), 1)))
}


type tankCannon struct {
	parent *Tank
	angle float64
}

func (t *Tank) newCannon *tankCannon {
	return &tankCannon {
		parent: t,
		angle: t.Angle,
	}
}
