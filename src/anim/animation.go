package anim

import (
  "github.com/gen2brain/raylib-go/raylib"
)

type Animation struct {
	AnimationData
	pos rl.Vector2
	w float32
	h float32
	halfW float32
	halfH float32
	rotation float32
	texture *rl.Texture2D
}

type AnimationData struct {
	DrawTimer float32
	CurrFrame int
	MaxFrame int
	FrameTime float32
}

func (a *Animation) Draw() {
	rl.DrawTexturePro(*a.texture, rl.NewRectangle(a.w * float32(a.CurrFrame), 0, a.w, a.h), rl.NewRectangle(a.pos.X, a.pos.Y, a.w, a.h), rl.NewVector2(a.halfW, a.halfH), a.rotation, rl.White)
}

func (a *AnimationData) Update(dt float32) {
	a.DrawTimer += dt
	if a.DrawTimer > a.FrameTime {
		a.DrawTimer = 0
		a.CurrFrame++
	}
}
