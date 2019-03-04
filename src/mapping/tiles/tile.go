package tiles

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Tile struct {
	Pos rl.Vector2
	Resistance float32 // Deceleration multiplier when on tile
	texture *rl.Texture2D
}

func (t *Tile) Draw() {
	rl.DrawTextureEx(*t.texture, t.Pos, 0, 1, rl.White)
}
