package entities

import (
  "github.com/gen2brain/raylib-go/raylib"
)

type Shell struct {
	BaseBody
	Damage int
	Size rl.Vector2    // Where X is Width, Y is Height
}

func (s *Shell) Draw() {
	rl.DrawRectangle(int32(s.Pos.X), int32(s.Pos.Y), int32(s.Size.X), int32(s.Size.Y), rl.Red)
}
