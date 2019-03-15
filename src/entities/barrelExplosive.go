package entities

import (
  "github.com/gen2brain/raylib-go/raylib"
	"tools"
)

const (
	BARREL_EXPL_MAX_DAMAGE int = 200
	BARREL_EXPL_RADIUS float32 = 100
	BARREL_EXPL_RADIUS_SQR float32 = BARREL_EXPL_RADIUS*BARREL_EXPL_RADIUS
	BARREL_EXPL_SHAPE_RADIUS float32 = 10
)

type BarrelExplosive struct {
	baseExplosive
}

func NewBarrelExplosive(pos rl.Vector2) *BarrelExplosive {
	return &BarrelExplosive {
		baseExplosive: baseExplosive {
			Pos: pos,
			MaxDamage: BARREL_EXPL_MAX_DAMAGE,
			ExplosionRadius: BARREL_EXPL_RADIUS,
			explRadiusSqr: BARREL_EXPL_RADIUS_SQR,
			hasExploded: false,
		},
	}
}

func (e *BarrelExplosive) Draw() {
	rl.DrawCircleV(e.Pos, BARREL_EXPL_RADIUS, rl.Green)
	rl.DrawCircleV(e.Pos, BARREL_EXPL_SHAPE_RADIUS, rl.Red)
}

func (e *BarrelExplosive) Update(dt float32) {
	for i := 0; i < len(G.Ent.players); i++ {
		if tools.InCircle(e.Pos, e.explRadiusSqr, G.Ent.players[i].Pos) {
			println("Exploding.")
			e.Explode()
		}
	}
}
