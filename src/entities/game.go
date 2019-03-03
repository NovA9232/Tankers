package entities

import (
  "github.com/gen2brain/raylib-go/raylib"
	"anim"
)

const (
	PiOv2 float32 = rl.Pi/2
	TwoPi float32 = rl.Pi * 2
)

var (
	SCREEN_W float32
	SCREEN_H float32
)

type Thing interface {  // Can be drawn and updated
  Draw()
  Update(float32)
}

type Game struct {
	Ent *Entities
	Anim []anim.Animation
}

func (g *Game) Update(dt float32) {
	g.Ent.UpdateAllEntites(rl.GetFrameTime())
	for i := 0; i < len(g.Anim); i++ {
		g.Anim[i].Update(dt)
	}
	g.Anim = anim.CheckAnimations(g.Anim)
}

func (g *Game) Draw() {
	g.Ent.DrawAllEntites()
	for i := 0; i < len(g.Anim); i++ {
		g.Anim[i].Draw()
	}
}
