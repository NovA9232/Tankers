package bodies

import (
  "github.com/gen2brain/raylib-go/raylib"
  "tools"
	"math"
)

const (
  TANK_W float32 = 60
  TANK_H float32 = 100
  TANK_ACCELL = 200
  TANK_DECEL = 0.99
  TANK_TURN_SPD float32 = rl.Pi
	TANK_CANN_TURN_SPD float32 = TwoPi

	HALF_TANK_W = TANK_W/2
	HALF_TANK_H = TANK_H/2
)

var (
	tankTex rl.Texture2D
	tankFrame rl.Rectangle = rl.NewRectangle(0, 2, TANK_W, TANK_H) // Part of tank was redrawn (small line)
	tankCannFrame rl.Rectangle = rl.NewRectangle(0, 0, TANK_W, TANK_H)

	tankCannonTex rl.Texture2D
)

type Tank struct {
  BaseBody
	Cannon *tankCannon
  colors []rl.Color
}

func NewTank(IDNum int, newPos rl.Vector2) *Tank {
	if tankTex.ID == uint32(0) {
		println("Loading tankBody.png texture.")
		tankTex = rl.LoadTexture("src/assets/exports/tankBodySand.png")
	}

	t := new(Tank)
  t.BaseBody = NewBody(NewID(IDNum, "tank"), newPos, 0, 0)
  t.colors = []rl.Color{rl.Lime, rl.Lime, rl.SkyBlue, rl.Lime}
	t.newCannon()

	return t
}

func (t *Tank) Draw() {
	rl.DrawTexturePro(tankTex, tankFrame, rl.NewRectangle(t.Pos.X, t.Pos.Y, TANK_W, TANK_H), rl.NewVector2(HALF_TANK_W, HALF_TANK_H), t.Angle*rl.Rad2deg, rl.White)
	t.Cannon.draw()
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

	t.Cannon.update(dt)
}

func (t *Tank) accelerate(dt, totalPower float32) {
  t.VelMag += (TANK_ACCELL * dt * totalPower)
}

func (t *Tank) applyResistance(dt float32) {
  t.VelMag = t.VelMag * (TANK_DECEL * float32(math.Min(144/float64(dt), 1)))
}


type tankCannon struct {
	parent *Tank
	angle float32
}

func (t *Tank) newCannon() {
	if tankCannonTex.ID == uint32(0) {
		println("Loading tankCann.png texture.")
		tankCannonTex = rl.LoadTexture("src/assets/exports/tankCannSand.png")
	}

	t.Cannon = &tankCannon {
		parent: t,
		angle: t.Angle,
	}
}

func (c *tankCannon) draw() {   // Not exported since the parent should draw + update.
	rl.DrawTexturePro(tankCannonTex, tankCannFrame, rl.NewRectangle(c.parent.Pos.X, c.parent.Pos.Y, TANK_W, TANK_H), rl.NewVector2(HALF_TANK_W, HALF_TANK_H), (c.angle + PiOv2)*rl.Rad2deg, rl.White)
}

func (c *tankCannon) update(dt float32) {
	mPos := rl.GetMousePosition()
	tmp := tools.SubVec(&c.parent.Pos, &mPos)
//	rl.DrawLineEx(mPos, c.parent.Pos, 2, rl.Red)
	angleToMouse := tools.GetAngle(&tmp)
	diff := float32(math.Mod(float64(c.angle - angleToMouse), float64(TwoPi)))
	angleToTurn := TANK_CANN_TURN_SPD*dt

	if float32(math.Abs(float64(diff))) > angleToTurn/2 {
		if (diff < 0 || diff > rl.Pi) {
			c.angle += angleToTurn
		} else {
			c.angle -= angleToTurn
		}
	}
}
