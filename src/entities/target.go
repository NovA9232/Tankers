package entities

import (
	"github.com/gen2brain/raylib-go/raylib"
	"github.com/ByteArena/box2d"
)

type Target struct {
	box2d.B2Body
}

func NewTarget(pos rl.Vector2, w, h int) *Target {
	bodDef := box2d.NewB2BodyDef()
	bodDef.Position.Set(float64(pos.X), float64(pos.Y))
	bodDef.Active = true

	shape := box2d.NewB2EdgeShape()
	shape.Set(box2d.MakeB2Vec2(float32()))

	t := &Target {
		box2d.B2Body: box2d.NewB2Body(bodDef, G.PhysicsWorld),
	}
}

func (t *Target) Draw() {
	rl.DrawRectangleRec(rl.NewRectangle(float32(t.Position.X-(t.Width.X/2)), float32(t.Position.Y-(t.Width.Y/2)), float32(t.Width.X), float32(t.Width.Y)), rl.Red)
}

func (t *Target) Update(dt float32) {
	t.Velocity.X, t.Velocity.Y = 0, 0
}
