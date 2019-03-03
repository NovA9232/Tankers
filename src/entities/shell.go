package entities

import (
  "github.com/gen2brain/raylib-go/raylib"
	"tools"
)

const (
	SHELL_TIMEOUT float32 = 10 // Time in seconds
	SHELL_WIDTH int = 5
	SHELL_HEIGHT int = 3
)

type Shell struct {
	Projec
	Damage int
	Size rl.Vector2    // Where X is Width, Y is Height
}

func NewShell(IDnum int, pos, parentVel rl.Vector2, vel, angle float32, damage, W, H int) *Shell {
	s := new(Shell)
  s.BaseBody = NewBody(NewID(IDnum, "shell"), pos, vel, angle)
	s.Damage = damage
	s.Size = rl.NewVector2(float32(W), float32(H))
	s.timeOfCreation = rl.GetTime()
	s.timeLimit = SHELL_TIMEOUT


	v := tools.GetXYComponent(s.VelMag, s.Angle)
	v.X += parentVel.X
	v.Y += parentVel.Y
	s.VelMag = tools.GetMagnitude(&v)
	return s
}

func (s *Shell) Draw() {
	rl.DrawRectangle(int32(s.Pos.X), int32(s.Pos.Y), int32(s.Size.X), int32(s.Size.Y), rl.Red)
}

func (s *Shell) Update(dt float32) {
	vel := tools.GetXYComponent(s.VelMag, s.Angle)
	s.Pos.X += (vel.X*dt)
	s.Pos.Y -= (vel.Y*dt)
}

