package tiles

import (
	"github.com/gen2brain/raylib-go/raylib"
)

var (
	concreteTex rl.Texture2D
)

func NewConcreteTile(pos rl.Vector2) *Tile {
	if concreteTex.ID == 0 {
		println("Loading cracked stone tile texture.")
		concreteTex = rl.LoadTexture("src/assets/map/concreteTile.png")
	}
	return &Tile{
		Pos: pos,
		Resistance: 1,   // Multiplier of current deceleration number
		texture: &concreteTex,
	}
}
