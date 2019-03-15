package entities

import (
  "github.com/gen2brain/raylib-go/raylib"
	"tools"
	"math"
)

type Explosive interface {
	Entity
	Explode()
	Exploded() bool
}

type baseExplosive struct {
	Pos rl.Vector2
	MaxDamage int
	ExplosionRadius float32
	explRadiusSqr float32
	hasExploded bool
}

func (e *baseExplosive) Exploded() bool {
	return e.hasExploded
}

func (e *baseExplosive) Explode() {
	for i := 0; i < len(G.Ent.players); i++ {
		if tools.InCircle(e.Pos, e.explRadiusSqr, G.Ent.players[i].Pos) {
			dist := math.Sqrt( math.Pow(float64(e.Pos.X - G.Ent.players[i].Pos.X), 2) + math.Pow(float64(e.Pos.Y - G.Ent.players[i].Pos.Y), 2) )
			G.Ent.players[i].Health.Val -= int(float64(e.MaxDamage)/(dist*dist))   // Inverse square law
			if !e.hasExploded { e.hasExploded = true }
		}
	}
}

//func (e *baseExplosive) checkForProjectilesInside() bool {  // True if projectile is in circle bounds
//	
//}
