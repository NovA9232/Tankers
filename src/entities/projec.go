package entities

type Projec struct {
	BaseBody
	timeOfCreation float32
	timeLimit float32
}

func (p *Projec) getTimeOfCreation() float32 {
	return p.timeOfCreation
}

func (p *Projec) getTimeLimit() float32 {
	return p.timeLimit
}
