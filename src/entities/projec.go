package entities

import (
	"github.com/ByteArena/box2d"
)

type Projec struct {
	box2d.B2Body
	timeOfCreation float32
	timeLimit float32
}

func (p *Projec) getTimeOfCreation() float32 {
	return p.timeOfCreation
}

func (p *Projec) getTimeLimit() float32 {
	return p.timeLimit
}
