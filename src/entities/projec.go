package entities

import (
	b2 "github.com/neguse/go-box2d-lite/box2dlite"
)

type Projec struct {
	b2.Body
	timeOfCreation float32
	timeLimit float32
}

func (p *Projec) getTimeOfCreation() float32 {
	return p.timeOfCreation
}

func (p *Projec) getTimeLimit() float32 {
	return p.timeLimit
}
