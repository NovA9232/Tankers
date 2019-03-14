package entities

import (
  "github.com/gen2brain/raylib-go/raylib"
	"tools"
)

const (
	SHELL_TIMEOUT float32 = 6 // Time in seconds
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
	rl.DrawTexturePro(shellTex, shellDrawFrame, rl.NewRectangle(s.Pos.X, s.Pos.Y, shellDrawFrame.Width, shellDrawFrame.Height), rl.NewVector2(shellDrawFrame.Width/2, shellDrawFrame.Height/2), s.Angle*rl.Rad2deg, rl.White)
}

func (s *Shell) Update(dt float32) {
	vel := tools.GetXYComponent(s.VelMag, s.Angle)
	s.Pos.X += (vel.X*dt)
	s.Pos.Y -= (vel.Y*dt)
}

