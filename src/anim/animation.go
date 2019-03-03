package anim

import (
  "github.com/gen2brain/raylib-go/raylib"
)

type Animation struct {
	pos rl.Vector2
	drawTimer float32
	CurrFrame int
	MaxFrame int
	frameTime float32
	w float32
	h float32
	halfW float32
	halfH float32
	rotation float32
	texture *rl.Texture2D
}

func (a *Animation) Draw() {
	rl.DrawTexturePro(ExplTex, rl.NewRectangle(a.w * float32(a.CurrFrame), 0, a.w, a.h), rl.NewRectangle(a.pos.X, a.pos.Y, a.w, a.h), rl.NewVector2(a.halfW, a.halfH), a.rotation, rl.White)
}

func (a *Animation) Update(dt float32) {
	a.drawTimer += dt
	if a.drawTimer > a.frameTime {
		a.drawTimer = 0
		a.CurrFrame++
	}
}
