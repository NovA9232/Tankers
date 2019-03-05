package tiles

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	crackedStoneTex rl.Texture2D
)

func NewCrackedStoneTile(pos rl.Vector2, w, h float32) Tile {
	if crackedStoneTex.ID == 0 {
		println("Loading cracked stone tile texture.")
		crackedStoneTex = rl.LoadTexture("src/assets/map/tiles/crackedStone.png")
	}
	return Tile (&BaseTile {
		Pos: pos,
		Resistance: 1.1,   // Multiplier of current deceleration number
		texture: &crackedStoneTex,
		W: w,
		H: h,
	})
}
