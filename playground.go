package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Playground struct {
	entities []Object
	selected Object
}

func NewPlayground() *Playground {
	return &Playground{entities: []Object{}}
}

func (p *Playground) AddObject(e Object) {
	p.entities = append(p.entities, e)
}

func (p Playground) GetEntitiesAtMousePosition() []Object {
	var entities []Object
	for _, e := range p.entities {
		if e.IsMousedOver() {
			entities = append(entities, e)
		}
	}
	return entities
}

func (p Playground) GetTopEntityAtMousePosition() (int, Object) {
	for i, e := range p.entities {
		if e.IsMousedOver() {
			return i, e
		}
	}
	return -1, nil
}

func (p *Playground) Update() {

	p.DoMouse()
	for _, e := range p.entities {
		e.Update()
	}

}

func (p *Playground) MoveObjectToTop(i int) {
	p.entities = moveObjectToTop(p.entities, i)
}

func (p *Playground) DoMouse() {
	i, mousedOverObject := p.GetTopEntityAtMousePosition()
	if mousedOverObject != nil {
		if rl.IsMouseButtonPressed(0) {
			p.selected = mousedOverObject
			p.MoveObjectToTop(i)
		}
	}

	if p.selected != nil {
		pos := p.selected.GetPos()
		mouseDelta := rl.GetMouseDelta()
		p.selected.SetPos(pos.X+mouseDelta.X, pos.Y+mouseDelta.Y)
	}

	if rl.IsMouseButtonReleased(0) {
		p.selected = nil
	}
}

func (p *Playground) Draw() {
	for _, e := range p.entities {
		e.Draw()

		_, mousedOver := p.GetTopEntityAtMousePosition()
		if mousedOver != nil {
			mousedOver.DrawMouseBox()
		}
	}
}
