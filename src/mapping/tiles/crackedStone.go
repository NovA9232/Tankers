package tiles

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	crackedStoneTex rl.Texture2D
)

func NewCrackedStoneTile(pos rl.Vector2) *Tile {
	if crackedStoneTex.ID == 0 {
		println("Loading cracked stone tile texture.")
		crackedStoneTex = rl.LoadTexture("src/assets/map/crackedStoneTile.png")
	}
	return &Tile{
		Pos: pos,
		Resistance: 1.1,   // Multiplier of current deceleration number
		texture: &crackedStoneTex,
	}
}
