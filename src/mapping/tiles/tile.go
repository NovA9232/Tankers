package tiles

import (
	"github.com/gen2brain/raylib-go/raylib"
	"anim"
)

type Tile struct {
	Pos rl.Vector2
	Animated bool
	AnimData anim.AnimationData
	Resistance float32 // Deceleration multiplier when on tile
	texture *rl.Texture2D
}

func (t *Tile) Draw() {
	rl.DrawTextureEx(*t.texture, t.Pos, 0, 1, rl.White)
}
