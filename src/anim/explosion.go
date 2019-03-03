package anim

import (
  "github.com/gen2brain/raylib-go/raylib"
)

var (
	ExplTex rl.Texture2D
	ExplFrameTime float32 = 0.1
)

type Explosion struct {
	pos rl.Vector2
	drawTimer float32
	currFrame int
	maxFrame int
	frameTime float32
	w int
	h int
	texture *rl.Texture2D
}

func NewExplosion(p rl.Vector2) *Explosion {
	if ExplTex.ID == uint32(0) {
		println("Loading Expl texture.")
		ExplTex = rl.LoadTexture("src/assets/animations/explosion/explosion.png")
	}

	return &Explosion {
		pos: p,
		drawTimer: 0,
		currFrame: 0,
		maxFrame: 5,
		frameTime: ExplFrameTime,
		w: 20,
		h: 20,
		texture: &ExplTex,
	}
}

func (e *Explosion) Draw() {
	rl.DrawTexturePro(ExplTex, rl.NewRectangle(float32(20 * e.currFrame), 0, 20, 20), rl.NewRectangle(e.pos.X, e.pos.Y, 20, 20), rl.NewVector2(10, 10), 0, rl.White)
}

func (e *Explosion) Update(dt float32) {
	e.drawTimer += dt
	if e.drawTimer > e.frameTime {
		e.drawTimer = 0
		e.currFrame++
	}
}

func (e *Explosion) CheckAnimationFinished() bool {
	return bool(e.currFrame >= e.maxFrame)
}
