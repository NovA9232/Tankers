package anim

import (
  "github.com/gen2brain/raylib-go/raylib"
)

const (
	ExplFrameTime float32 = 0.05
)

var (
	ExplTex rl.Texture2D
)

func NewExplosion(p rl.Vector2) *Animation {
	if ExplTex.ID == uint32(0) {
		println("Loading Expl texture.")
		ExplTex = rl.LoadTexture("src/assets/animations/explosion/explosion.png")
	}

	return &Animation {
		pos: p,
		drawTimer: 0,
		CurrFrame: 0,
		MaxFrame: 5,
		frameTime: ExplFrameTime,
		w: 20,
		h: 20,
		halfW: 10,
		halfH: 10,
		rotation: 0,
		texture: &ExplTex,
	}
}

