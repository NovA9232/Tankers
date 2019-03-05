package tiles

import (
	"github.com/gen2brain/raylib-go/raylib"
	"anim"
)

var (
	waterTex rl.Texture2D
	waterAnimData anim.AnimationData = anim.AnimationData {
		DrawTimer: 0,
		CurrFrame: 0,
		MaxFrame: 0,
		FrameTime: 0.05,
	}
)

type waterTile struct {
	AnimatedTile
}

func (w *waterTile) Draw() {
	rl.DrawTexturePro(*w.texture, rl.NewRectangle(float32(w.animData.CurrFrame), 0, w.W, w.H), rl.NewRectangle(w.Pos.X, w.Pos.Y, w.W, w.H), rl.NewVector2(0, 0), 0, rl.White)
}

func NewWaterTile(pos rl.Vector2, w, h float32) Tile {
	if waterTex.ID == 0 {
		println("Loading water texture.")
		waterTex = rl.LoadTexture("src/assets/map/tiles/water.png")
	}

	return Tile ( &waterTile {
		AnimatedTile: AnimatedTile {
			BaseTile: BaseTile {
				Pos: pos,
				Resistance: 2,
				texture: &waterTex,
				W: w,
				H: h,
			},
			halfW: w/2,
			halfH: h/2,
			animData: waterAnimData,
		},
	})
}
