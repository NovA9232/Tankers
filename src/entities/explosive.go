package entities

import (
  "github.com/gen2brain/raylib-go/raylib"
	"math"
	"fmt"

	"tools"
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
			fmt.Println(dist ,"dist")
			damage := int(math.Pow(3, (-dist*3/float64(e.MaxDamage)))*float64(e.MaxDamage))   // Inverse square law
			fmt.Println(damage, "damage")
			G.Ent.players[i].Health.Val -= damage
			if !e.hasExploded { e.hasExploded = true }
		}
	}
}

//func (e *baseExplosive) checkForProjectilesInside() bool {  // True if projectile is in circle bounds
//	
//}
