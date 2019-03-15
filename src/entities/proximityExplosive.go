package entities

import (
  "github.com/gen2brain/raylib-go/raylib"
	"tools"
)

const (
	PROX_EXPL_MAX_DAMAGE int = 100
	PROX_EXPL_RADIUS float32 = 120
	PROX_EXPL_RADIUS_SQR float32 = PROX_EXPL_RADIUS*PROX_EXPL_RADIUS
	PROX_EXPL_SHAPE_RADIUS float32 = 10
	PROX_TIME_TO_EX = 0.5
)

type ProximityExplosive struct {
	baseExplosive
	timer float32
	fuseLit bool
}

func NewProximityExplosive(pos rl.Vector2) *ProximityExplosive {
	return &ProximityExplosive {
		baseExplosive: baseExplosive {
			Pos: pos,
			MaxDamage: PROX_EXPL_MAX_DAMAGE,
			ExplosionRadius: PROX_EXPL_RADIUS,
			explRadiusSqr: PROX_EXPL_RADIUS_SQR,
			hasExploded: false,
		},
	}
}

func (e *ProximityExplosive) Draw() {
	rl.DrawCircleV(e.Pos, PROX_EXPL_RADIUS, rl.NewColor(255, 0, 0, 64))
	var clr rl.Color = rl.Green
	if e.fuseLit {
		clr = rl.Red
	}
	rl.DrawCircleV(e.Pos, PROX_EXPL_SHAPE_RADIUS, clr)
}

func (e *ProximityExplosive) Update(dt float32) {
	if e.fuseLit {
		e.timer += dt
		if e.timer >= PROX_TIME_TO_EX {
			e.Explode()
		}
	} else {
		for i := 0; i < len(G.Ent.players); i++ {
			if tools.InCircle(e.Pos, e.explRadiusSqr, G.Ent.players[i].Pos) {
				println("Exploding.")
				e.fuseLit = true
			}
		}
	}
}
