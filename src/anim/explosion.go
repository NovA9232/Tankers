package anim

import (
  "github.com/gen2brain/raylib-go/raylib"
)

var (
	ExplTex rl.Texture2D

	ExplAnimationData AnimationData = AnimationData {
		drawTimer: 0,
		CurrFrame: 0,
		MaxFrame: 5,
		frameTime: 0.05,
	}
)

func NewExplosion(p rl.Vector2) *Animation {
	if ExplTex.ID == 0 {
		println("Loading Expl texture.")
		ExplTex = rl.LoadTexture("src/assets/animations/explosion/explosion.png")
	}

	return &Animation {
		AnimationData: ExplAnimationData,
		pos: p,
		w: 20,
		h: 20,
		halfW: 10,
		halfH: 10,
		rotation: 0,
		texture: &ExplTex,
	}
}

