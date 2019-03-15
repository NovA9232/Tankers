package ui

import (
  "github.com/gen2brain/raylib-go/raylib"
)

const (
	HEALTH_BAR_W float32 = 60
	HEALTH_BAR_H float32 = 10
)

type HealthBar struct {
	Pos rl.Vector2
	Val int
	maxHealth int
	normalised float32  // Between 0 and 1
}

func NewHealthBar(value, maxValue int) *HealthBar {
	return &HealthBar {
		Pos: rl.NewVector2(0, 0),
		Val: value,
		maxHealth: maxValue,
		normalised: 1,
	}
}

func (h *HealthBar) Update() {
	h.normalised = float32(h.Val)/float32(h.maxHealth)
}

func (h *HealthBar) Draw() {
	rl.DrawRectangleLinesEx(rl.NewRectangle(h.Pos.X, h.Pos.Y, HEALTH_BAR_W, HEALTH_BAR_H), 2, rl.Black)
	rl.DrawRectangleRec(rl.NewRectangle(float32(int(h.Pos.X) + 2), float32(int(h.Pos.Y) + 2), (HEALTH_BAR_W * h.normalised) - 4, HEALTH_BAR_H - 4), rl.Red)
}
