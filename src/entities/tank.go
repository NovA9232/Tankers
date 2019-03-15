package entities

import (
  "github.com/gen2brain/raylib-go/raylib"
	"math"
  "tools"
	"anim"
)

const (
  TANK_W float32 = 60
  TANK_H float32 = 100
  TANK_ACCELL = 400
  TANK_DECEL = 5
  TANK_TURN_SPD float32 = rl.Pi
	TANK_CANN_TURN_SPD float32 = TwoPi
	TANK_SHELL_SPEED float32 = 1400
	TANK_FIRE_COOLDOWN float32 = 0.2

	HALF_TANK_W = TANK_W/2
	HALF_TANK_H = TANK_H/2
)

var (
	tankTex rl.Texture2D
	tankFrame rl.Rectangle = rl.NewRectangle(0, 2, TANK_W, TANK_H-2) // Part of tank was redrawn (small line)
	tankCannonTex rl.Texture2D
)

type Tank struct {
  BaseBody
	Deceleration float64
	Cannon *tankCannon
}

func NewTank(IDNum int, pos rl.Vector2) *Tank {
	if tankTex.ID == 0 {
		println("Loading tankBody.png texture.")
		tankTex = rl.LoadTexture("src/assets/tank/tankBodySand.png")
	}

	t := new(Tank)
  t.BaseBody = NewBody(NewID(IDNum, "tank"), pos, 0, 0)
	t.Deceleration = TANK_DECEL
	t.newCannon()

	return t
}

func (t *Tank) Draw() {
	rl.BeginShaderMode(*Shader)
	rl.DrawTexturePro(tankTex, tankFrame, rl.NewRectangle(t.Pos.X, t.Pos.Y, TANK_W, TANK_H), rl.NewVector2(HALF_TANK_W, HALF_TANK_H), t.Angle*rl.Rad2deg, rl.White)
	t.Cannon.draw()
	rl.EndShaderMode()
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

	vel := tools.GetXYComponent(t.VelMag, t.Angle)
  t.Pos.X += (vel.X*dt)
  t.Pos.Y -= (vel.Y*dt)

	t.Cannon.update(dt)
}

func (t *Tank) accelerate(dt, totalPower float32) {
  t.VelMag += (TANK_ACCELL * dt * totalPower)
}

func (t *Tank) applyResistance(dt float32) {
	tileDec := G.WorldMap.GetResistanceAt(t.Pos.X, t.Pos.Y)

  t.VelMag *= float32(math.Pow(t.Deceleration, float64(-dt*tileDec)))
	if math.Abs(float64(t.VelMag)) < 0.1 {
		t.VelMag = 0
	}
}


type tankCannon struct {
	parent *Tank
	angle float32
	lastShotTime float32
	firstShot bool

	recoilTimer float32
	recoilFrame int
	maxRecoilFrame int
}

func (t *Tank) newCannon() {
	if tankCannonTex.ID == uint32(0) {
		println("Loading tankCann texture.")
		tankCannonTex = rl.LoadTexture("src/assets/animations/tankCann/tankCannSandRecoil.png")
	}

	t.Cannon = &tankCannon {
		parent: t,
		angle: t.Angle,
		lastShotTime: 0,
		firstShot: true,
	}
}

func (c *tankCannon) draw() {   // Not exported since the parent should draw + update.
	rl.DrawTexturePro(tankCannonTex, rl.NewRectangle(TANK_W * float32(c.recoilFrame), 0, TANK_W, TANK_H), rl.NewRectangle(c.parent.Pos.X, c.parent.Pos.Y, TANK_W, TANK_H), rl.NewVector2(HALF_TANK_W, HALF_TANK_H), (c.angle + PiOv2)*rl.Rad2deg, rl.White)
}

func (c *tankCannon) update(dt float32) {
	c.angle = float32(math.Mod(float64(c.angle), float64(TwoPi)))
	angleToMouse := tools.GetAngle( tools.SubVec(c.parent.Pos, rl.GetMousePosition()) )
	diff := angleToMouse - c.angle
	if diff > rl.Pi {
		diff -= TwoPi
	} else if diff < -rl.Pi {
		diff += TwoPi
	}
	angleToTurn := TANK_CANN_TURN_SPD*dt

	if float32(math.Abs(float64(diff))) > angleToTurn/2 {
		if diff > 0 {
			c.angle += angleToTurn
		} else {
			c.angle -= angleToTurn
		}
	} else {
		c.angle = angleToMouse
	}

	if rl.IsMouseButtonDown(0) && (rl.GetTime() - c.lastShotTime > TANK_FIRE_COOLDOWN || c.firstShot) {
		c.Fire()
		if c.firstShot {
			c.firstShot = false
		}
	}

	if c.recoilFrame != 0 {
		c.recoilTimer += dt
		if c.recoilTimer > 0.02 {
			c.recoilTimer = 0
			c.recoilFrame++
			if c.recoilFrame > 10 {   // 10 frames in image
				c.recoilFrame = 0
			}
		}
	}
}

func (c *tankCannon) getEndOfCannPos() rl.Vector2 {
	pos := rl.NewVector2(c.parent.Pos.X, c.parent.Pos.Y - HALF_TANK_H)   // Position when
	tools.RotateVec(&pos, &c.parent.Pos, c.angle + (rl.Pi/2))
	pos.Y += HALF_TANK_H
	return pos
}

func (c *tankCannon) Fire() {
	endPos := c.getEndOfCannPos()
	G.Ent.projec = append(G.Ent.projec, Projectile( NewShell(len(G.Ent.projec), endPos, tools.GetXYComponent(c.parent.VelMag, c.parent.Angle), TANK_SHELL_SPEED, c.angle + (rl.Pi/2), 100, int(shellDrawFrame.Width),  int(shellDrawFrame.Height))))
	c.lastShotTime = rl.GetTime()

	G.Anim = append(G.Anim, anim.NewExplosion(endPos))
	c.recoilFrame = 1
}

