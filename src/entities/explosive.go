package entities

import (
	"github.com/gen2brain/raylib-go/raylib"
)

type Explosive struct {
	Pos rl.Vector2
	Damage int
	RadiusOfExplosion float32
	
}
