package entities

import (
  "github.com/gen2brain/raylib-go/raylib"
	b2 "github.com/neguse/go-box2d-lite/box2dlite"
	"tools"
)

const (
	SHELL_TIMEOUT float32 = 6 // Time in seconds
	SHELL_MASS float64 = 2
)

var (
	shellDrawFrame rl.Rectangle = rl.NewRectangle(0, 0, 6, 10)
	shellTex rl.Texture2D
)

type Shell struct {
	Projec
	Damage int
	Size rl.Vector2    // Where X is Width, Y is Height
}

func NewShell(IDnum int, pos, parentVel rl.Vector2, vel, angle float32, damage, W, H int) *Shell {
	if shellTex.ID == 0 {
		println("Loading shell texture.")
		shellTex = rl.LoadTexture("src/assets/shell/shell.png")
	}
	v := tools.GetXYComponent(vel, (TwoPi-angle)+rl.Pi)
	v.X += parentVel.X
	v.Y -= parentVel.Y
	s.Velocity.X, s.Velocity.Y = float64(v.X), float64(v.Y)

	bodDef := box2d.NewB2BodyDef()
	bodDef.Position.Set(float64(pos.X), float64(pos.Y))
	bodDef.Active = true
	bodDef.Velocity = 

	s.Damage = damage
	s.timeOfCreation = rl.GetTime()
	s.timeLimit = SHELL_TIMEOUT

	return s
}

func (s *Shell) Draw() {
	rl.DrawTexturePro(shellTex, shellDrawFrame, rl.NewRectangle(float32(s.Position.X), float32(s.Position.Y), shellDrawFrame.Width, shellDrawFrame.Height), rl.NewVector2(shellDrawFrame.Width/2, shellDrawFrame.Height/2), float32(s.Rotation)*rl.Rad2deg, rl.White)
}

func (s *Shell) Update(dt float32) {}

